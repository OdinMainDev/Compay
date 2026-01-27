package efwcontroller

import (
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetModel[TModel any, TQueryHelper any](
	context *gin.Context,
	getModel func(queryHelper TQueryHelper) *TModel,
	modelToDto func(model *TModel) any,
	constructQueryHelper func(int64) TQueryHelper,
) (*TModel, int64, bool) {
	var input RequestWithId
	if !efwgin.ShouldBindUri(context, &input) {
		return nil, 0, false
	}

	model := getModel(constructQueryHelper(input.Id))
	if model == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return nil, 0, false
	}

	if modelToDto == nil {
		efwginhelpers.ResponseSuccess(context, gin.H{"data": model})
		return model, 0, false
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": modelToDto(model)})
	return model, input.Id, true
}
