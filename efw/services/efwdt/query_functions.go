package efwdt

import (
	"context"
	"efw/utils"
	"reflect"

	"gorm.io/gorm"
)

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
