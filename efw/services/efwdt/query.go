package efwdt

import (
	"context"
	"efw"
	"fmt"
	"sync"
)

var ErrorQueryResponse = func(queryErr, recordsFilteredErr, recordsTotalErr error) error {
	var err1, err2, err3 string
	if queryErr != nil {
		err1 = fmt.Sprintf("%s %s", "QueryError:", queryErr.Error())
	}
	if recordsFilteredErr != nil {
		err2 = fmt.Sprintf("%s %s", "RecordsFilteredError:", recordsFilteredErr.Error())
	}
	if recordsTotalErr != nil {
		err3 = fmt.Sprintf("%s %s", "RecordsTotalErr:", recordsTotalErr.Error())
	}

	return fmt.Errorf("%s %s %s %s", "[EFW: DATATABLE] QueryResponseError:", err1, err2, err3)
}

func GetQueryResponse[TModel, TRequest any](
	input Request[TRequest],
	additionalFilters []ChainFilter[*TModel],
	specialFilters SpecialFilterMap,
) ([]*TModel, int64, int64, int64, error) {
	dbTx := efw.DBTransaction[*TModel]()

	baseTx := applyAdditionalFilters(dbTx, additionalFilters)
	filteredTx := applyRequestFilters(baseTx, input, specialFilters)

	var model TModel
	var tableName = any(&model).(HasTableName).TableName()
	var wg sync.WaitGroup
	var result []*TModel
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
		recordsFiltered, recordsFilteredErr = filteredTx.Count(context.Background(), tableName+".id")
	}()

	go func() {
		defer wg.Done()
		recordsTotal, recordsTotalErr = baseTx.Count(context.Background(), tableName+".id")
	}()

	wg.Wait()

	if queryErr != nil || recordsFilteredErr != nil || recordsTotalErr != nil {
		return nil, 0, 0, 0, ErrorQueryResponse(queryErr, recordsFilteredErr, recordsTotalErr)
	}

	recordsCount := int64(calcRecordsCount(result))
	return result, recordsTotal, recordsFiltered, recordsCount, nil
}
