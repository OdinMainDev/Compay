package efwcontroller

import (
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetModels[TModel any, TQueryHelper any](
	context *gin.Context,
	getModels func(queryHelper TQueryHelper) []*TModel,
	modelToDto func(model *TModel) gin.H,
	constructQueryHelper func(int64) TQueryHelper,
) ([]*TModel, int64, bool) {
	var input RequestWithId
	if !efwgin.ShouldBindUri(context, &input) {
		return nil, 0, false
	}

	models := getModels(constructQueryHelper(input.Id))
	if models == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return nil, 0, false
	}

	if modelToDto == nil {
		efwginhelpers.ResponseSuccess(context, gin.H{"data": models})
		return models, 0, false
	}

	resp := []gin.H{}
	for _, model := range models {
		resp = append(resp, modelToDto(model))
	}
	efwginhelpers.ResponseSuccess(context, gin.H{"data": resp})
	return models, input.Id, true
}
