package efwdt

import (
	"context"
	"efw"
	"efw/core/efwlog"
	"sync"
)

func getQueryResponse[TModel, TRequest any](
	input Request[TRequest],
	additionalFilters []ChainFilter[TModel],
) (any, int64, int64, int64, error) {
	dbTx := efw.DBTransaction[TModel]()

	baseTx := applyAdditionalFilters(dbTx, additionalFilters)
	filteredTx := applyRequestFilters(baseTx, input)

	var wg sync.WaitGroup
	var result any
	var recordsFiltered int64
	var recordsTotal int64
	var queryErr, recordsFilteredErr, recordsTotalErr error

	wg.Add(3)

	go func() {
		defer wg.Done()
		result, queryErr = buildDataQuery(filteredTx, input).Find(context.Background())
	}()

	go func() {
		defer wg.Done()
		recordsFiltered, recordsFilteredErr = filteredTx.Count(context.Background(), "id")
	}()

	go func() {
		defer wg.Done()
		recordsTotal, recordsTotalErr = baseTx.Count(context.Background(), "id")
	}()

	wg.Wait()

	// тут нужно будет переделать сохранение ошибок, чтобы выглядело лучше
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

	recordsCount := int64(calcRecordsCount(result))
	return result, recordsTotal, recordsFiltered, recordsCount, nil
}
