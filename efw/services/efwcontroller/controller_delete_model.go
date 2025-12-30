package efwcontroller

import (
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteModel[TModel any, TQueryHelper any](
	context *gin.Context,
	getModel func(queryHelper TQueryHelper, additionalCacheKeys ...int64) *TModel,
	deleteModel func(model *TModel, additionalCacheKeys ...int64),
	constructQueryHelper func(int64) TQueryHelper,
	additionalCacheKeys ...int64,
) (int64, bool) {
	var input RequestWithId
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return 0, false
	}

	model := getModel(constructQueryHelper(input.Id), additionalCacheKeys...)
	if model == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return 0, false
	}

	deleteModel(model, additionalCacheKeys...)
	efwginhelpers.ResponseSuccess(context, nil)
	return input.Id, true
}
