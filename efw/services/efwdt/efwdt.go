package efwdt

import (
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChainFilter[T any] func(gorm.ChainInterface[T]) gorm.ChainInterface[T]

type FilterType string
type Filter struct {
	Name   string     `json:"name"`
	Type   FilterType `json:"type"`
	Values any        `json:"values"`
}

func Datable[TModel, TRequest any](
	context *gin.Context,
	columnsOrder []string,
	additionalFilters []ChainFilter[TModel],
) {
	var input Request[TRequest]
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	if err := input.ValidateOrderColumn(columnsOrder); err != nil {
		efwginhelpers.ResponseError(context, err)
		return
	}

	data, recordsCount, recordsFiltered, recordsTotal, err := getQueryResponse(input, additionalFilters)
	if err != nil {
		efwginhelpers.ResponseError(context, err)
		return
	}

	efwginhelpers.ResponseSuccess(context, response(data, recordsTotal, recordsFiltered, recordsCount))
}

func Filters(context *gin.Context, filters []Filter) {
	efwginhelpers.ResponseSuccess(context, gin.H{"data": filters})
}
