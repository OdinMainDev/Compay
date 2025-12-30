package efwcontroller

import (
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetModel[TModel any, TQueryHelper any](
	context *gin.Context,
	getModel func(queryHelper TQueryHelper, additionalCacheKeys ...int64) *TModel,
	modelToDto func(model *TModel) any,
	constructQueryHelper func(int64) TQueryHelper,
	additionalCacheKeys ...int64,
) (int64, bool) {
	var input RequestWithId
	if !efwgin.ShouldBindUri(context, &input) {
		return 0, false
	}

	model := getModel(constructQueryHelper(input.Id), additionalCacheKeys...)
	if model == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return 0, false
	}

	if modelToDto == nil {
		efwginhelpers.ResponseSuccess(context, gin.H{"data": model})
		return 0, false
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": modelToDto(model)})
	return input.Id, true
}
