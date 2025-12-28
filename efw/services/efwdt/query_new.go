package efwdt

import (
	"efw"
	"efw/utils"
	"sync"

	"gorm.io/gorm"
)

// func getQueryResponse

func getQueryResponse[TModel, TRequest any](
	model any,
	req Request[TRequest],
	availFilterColumns []string,
	additionalFilters []gorm.ChainInterface[TModel],
) (any, int64, int64, int64, error) {
	dbTrx := efw.DBTransaction[TModel]()

	baseTrx, err := prepareBaseQuery[TModel, TRequest](dbTrx, req, availFilterColumns)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	filteredTx := applyFilters(baseTrx, db, model, req)
	recordsFilteredTx := applyFilters(baseTrx, db, model, req)

	for _, value := range additionalFilters {
		req.FilterName = append(req.FilterName, key)
		req.FilterValue = append(req.FilterValue, value)
	}

	var wg sync.WaitGroup
	var result interface{}
	var recordsFiltered int64
	var recordsTotal int64
	var queryErr, recordsFilteredErr, recordsTotalErr error

	wg.Add(3)

	go func() {
		defer wg.Done()
		result, queryErr = applySpecialQueries(filteredTx, db, model, req)
	}()

	go func() {
		defer wg.Done()
		recordsFilteredTx, recordsFilteredErr = applySpecialCountQueries(recordsFilteredTx, db, model, req)
		if recordsFilteredErr != nil {
			return
		}
		recordsFilteredErr = recordsFilteredTx.Model(model).Count(&recordsFiltered).Error
	}()

	go func() {
		defer wg.Done()
		totalTx := db.Model(model)
		totalTx, recordsTotalErr = applySpecialCountQueries(totalTx, db, model, req)
		if recordsTotalErr != nil {
			return
		}
		recordsTotalErr = totalTx.Count(&recordsTotal).Error
	}()

	wg.Wait()

	if queryErr != nil {
		eslog.LogError(queryErr)
		return nil, 0, 0, 0
	}

	if recordsFilteredErr != nil {
		eslog.LogError(recordsFilteredErr)
		return nil, 0, 0, 0
	}

	if recordsTotalErr != nil {
		eslog.LogError(recordsTotalErr)
		return nil, 0, 0, 0
	}

	recordsCount := int64(recordsCount(result))
	return result, recordsTotal, recordsFiltered, recordsCount
}

func prepareBaseQuery[TModel, TRequest any](dbTrx gorm.Interface[TModel], request Request[TRequest], availFilterColumns []string) (gorm.ChainInterface[TModel], error) {
	if !utils.InArray(request.OrderColumn, availFilterColumns) {
		request.OrderColumn = "id"
	}

	if request.OrderDirection != "asc" && request.OrderDirection != "desc" {
		request.OrderDirection = "asc"
	}

	return dbTrx.Order(request.OrderColumn + " " + request.OrderDirection).Offset(int(request.Start)).Limit(request.Length), nil
}
