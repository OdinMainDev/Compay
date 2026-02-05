package efwcontroller

import (
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/core/efwlog"
	"efw/services/efwdt"

	"github.com/gin-gonic/gin"
)

func Datable[TModel, TRequest any](
	context *gin.Context,
	columnsOrder []string,
	additionalFilters []efwdt.ChainFilter[*TModel],
	specialFilters efwdt.SpecialFilterMap,
	dtoFunc func(model *TModel) any,
) {
	var input efwdt.Request[TRequest]
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	if err := input.ValidateOrderColumn(columnsOrder); err != nil {
		efwginhelpers.ResponseError(context, err)
		return
	}

	data, recordsCount, recordsFiltered, recordsTotal, err := efwdt.GetQueryResponse(input, additionalFilters, specialFilters)
	if err != nil {
		efwlog.LogError(err)
		efwginhelpers.ResponseError(context, err)
		return
	}

	efwginhelpers.ResponseSuccess(context, efwdt.Response(data, recordsTotal, recordsFiltered, recordsCount, dtoFunc))
}

func DatatableFilters(context *gin.Context, filters []efwdt.Filter) {
	efwginhelpers.ResponseSuccess(context, gin.H{"data": filters})
}
