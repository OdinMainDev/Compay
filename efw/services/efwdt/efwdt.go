package efwdt

import (
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
)

func Handle[TModel, TRequest any](
	context *gin.Context,
	columnsOrder []string,
	getCount func() int64,
	additionalFilters []ChainFilter[TModel],
) {
	var input Request[TRequest]
	if !efwgin.ShouldBindWithForm(context, &input) {
		return
	}

	if err := input.ValidateOrderColumn(columnsOrder); err != nil {
		efwginhelpers.ResponseError(context, err)
		return
	}

	countRecords := getCount()
	if input.Start > countRecords {
		efwginhelpers.ResponseSuccess(context, responseEmpty(countRecords))
		return
	}

	var dataSlice []*TModel
	data, recordsCount, recordsFiltered, recordsTotal, err := getQueryResponse(dataSlice, input, columnsOrder, additionalFilters)
	if err != nil {
		efwginhelpers.ResponseError(context, err)
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": response(data, recordsTotal, recordsFiltered, recordsCount)})
}
