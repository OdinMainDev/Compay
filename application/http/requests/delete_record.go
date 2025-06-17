package requests

import (
	"compay/application/helpers"
	"compay/application/service/esvalidator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Deletable interface {
	GetId() int64
	Delete(any)
}

func DeleteRecord[T Deletable, T2 any](
	context *gin.Context,
	getModelById func(id int64, queryHelper T2) T,
	queryHelper T2,
) (bool, int64) {
	var input RequestGetById

	if err := context.ShouldBindUri(&input); err != nil {
		helpers.GinErrorResponse(context, esvalidator.ValidatorResponse(err))
		return false, 0
	}

	model := getModelById(input.Id, queryHelper)
	if model.GetId() == 0 {
		helpers.GinErrorResponse(context, nil, http.StatusNotFound)
		return false, 0
	}

	model.Delete(model)
	helpers.GinSuccessResponse(context, nil)
	return true, model.GetId()
}
