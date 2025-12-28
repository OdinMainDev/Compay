package efwdt

import (
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Handle[TModel, TRequest any](
	context *gin.Context,
	columnsOrder []string,
	getCount func() int64,
	additionalFilters []gorm.ChainInterface[TModel],
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
		efwginhelpers.ResponseSuccess(context, responseEmpty[TModel](countRecords))
		return
	}

	var dataSlice []*TModel
	data, recordsCount, recordsFiltered, recordsTotal := getQueryResponse(dataSlice, req, columnsOrder, additionalFilters)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": req.Response(recordsTotal, recordsFiltered, recordsCount, data)})
}
