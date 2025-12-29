package efwdt

import (
	"context"
	"efw"
	"efw/core/efwlog"
	"sync"
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
