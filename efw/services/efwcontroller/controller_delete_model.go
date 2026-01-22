package efwcontroller

import (
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteModel[TModel any, TQueryHelper any](
	context *gin.Context,
	getModel func(queryHelper TQueryHelper) *TModel,
	deleteModel func(model *TModel, queryHelper TQueryHelper),
	constructQueryHelper func(int64) TQueryHelper,
	additionalCacheKeys ...int64,
) (int64, bool) {
	var input RequestWithId
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return 0, false
	}

	queryHelper := constructQueryHelper(input.Id)
	model := getModel(queryHelper)
	if model == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return 0, false
	}

	deleteModel(model, queryHelper)
	efwginhelpers.ResponseSuccess(context, nil)
	return input.Id, true
}
