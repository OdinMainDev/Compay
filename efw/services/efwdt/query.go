package efwdt

import (
	"context"
	"efw"
	"efw/core/efwlog"
	"efw/utils"
	"reflect"
	"sync"

	"gorm.io/gorm"
)

func getQueryResponse[TModel, TRequest any](
	model any,
	req Request[TRequest],
	availFilterColumns []string,
	additionalFilters []ChainFilter[TModel],
) (any, int64, int64, int64, error) {
	dbTx := efw.DBTransaction[TModel]()

	baseTx := applyAdditionalFilters(dbTx, additionalFilters)
	filteredTx := applyRequestFilters(baseTx, req)
	dataTx, err := prepareBaseQuery(filteredTx, req, availFilterColumns)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	var wg sync.WaitGroup
	var result any
	var recordsFiltered int64
	var recordsTotal int64
	var queryErr, recordsFilteredErr, recordsTotalErr error

	wg.Add(3)

	go func() {
		defer wg.Done()
		result, queryErr = queryData(dataTx, model)
	}()

	go func() {
		defer wg.Done()
		recordsFiltered, recordsFilteredErr = filteredTx.Count(context.Background(), "*")
	}()

	go func() {
		defer wg.Done()
		recordsTotal, recordsTotalErr = baseTx.Count(context.Background(), "*")
	}()

	wg.Wait()

	if queryErr != nil {
		efwlog.LogError(queryErr)
		return nil, 0, 0, 0, queryErr
	}

	if recordsFilteredErr != nil {
		efwlog.LogError(recordsFilteredErr)
		return nil, 0, 0, 0, recordsFilteredErr
	}

	if recordsTotalErr != nil {
		efwlog.LogError(recordsTotalErr)
		return nil, 0, 0, 0, recordsTotalErr
	}

	recordsCount := int64(recordsCount(result))
	return result, recordsTotal, recordsFiltered, recordsCount, nil
}

type ChainFilter[T any] func(gorm.ChainInterface[T]) gorm.ChainInterface[T]

type dbFilter[T any] interface {
	ApplyDBFilters(gorm.ChainInterface[T]) gorm.ChainInterface[T]
}

func applyAdditionalFilters[T any](chain gorm.ChainInterface[T], filters []ChainFilter[T]) gorm.ChainInterface[T] {
	for _, filter := range filters {
		if filter == nil {
			continue
		}
		chain = filter(chain)
	}
	return chain
}

func applyRequestFilters[TModel, TRequest any](chain gorm.ChainInterface[TModel], request Request[TRequest]) gorm.ChainInterface[TModel] {
	filter, ok := any(request.Conditions).(dbFilter[TModel])
	if !ok || isNilInterface(filter) {
		return chain
	}
	return filter.ApplyDBFilters(chain)
}

func queryData[T any](chain gorm.ChainInterface[T], model any) (any, error) {
	if model == nil {
		return chain.Find(context.Background())
	}

	dest, ok := normalizeDestination(model)
	if !ok {
		return chain.Find(context.Background())
	}

	if err := chain.Scan(context.Background(), dest); err != nil {
		return nil, err
	}
	return derefDestination(dest), nil
}

func normalizeDestination(model any) (any, bool) {
	value := reflect.ValueOf(model)
	if !value.IsValid() {
		return nil, false
	}

	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			ptr := reflect.New(value.Type().Elem())
			return ptr.Interface(), true
		}
		return model, true
	}

	ptr := reflect.New(value.Type())
	ptr.Elem().Set(value)
	return ptr.Interface(), true
}

func derefDestination(dest any) any {
	value := reflect.ValueOf(dest)
	for value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
		if value.IsNil() {
			return nil
		}
		value = value.Elem()
	}
	if !value.IsValid() {
		return nil
	}
	return value.Interface()
}

func isNilInterface(value any) bool {
	if value == nil {
		return true
	}
	kind := reflect.ValueOf(value).Kind()
	switch kind {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func:
		return reflect.ValueOf(value).IsNil()
	default:
		return false
	}
}

func recordsCount(col any) int {
	if col == nil {
		return 0
	}
	value := reflect.ValueOf(col)
	for value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
		if value.IsNil() {
			return 0
		}
		value = value.Elem()
	}
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		return value.Len()
	}
	return 0
}

func prepareBaseQuery[TModel, TRequest any](dbTrx gorm.ChainInterface[TModel], request Request[TRequest], availFilterColumns []string) (gorm.ChainInterface[TModel], error) {
	if !utils.InArray(request.OrderColumn, availFilterColumns) {
		request.OrderColumn = "id"
	}

	if request.OrderDirection != "asc" && request.OrderDirection != "desc" {
		request.OrderDirection = "asc"
	}

	return dbTrx.Order(request.OrderColumn + " " + request.OrderDirection).Offset(int(request.Start)).Limit(request.Length), nil
}
