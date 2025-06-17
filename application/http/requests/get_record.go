package requests

import (
	"compay/application/helpers"
	"compay/application/service/esvalidator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Identifiable interface {
	GetId() int64
}

type RequestGetById struct {
	Id int64 `uri:"id" json:"id" binding:"required,min=0"`
}

func GetRecord[T Identifiable, T2 any](
	context *gin.Context,
	getModel func(id int64, queryHelper T2) T,
	queryHelper T2,
) {
	var input RequestGetById

	if err := context.ShouldBindUri(&input); err != nil {
		helpers.GinErrorResponse(context, esvalidator.ValidatorResponse(err))
		return
	}

	model := getModel(input.Id, queryHelper)

	if model.GetId() == 0 {
		helpers.GinErrorResponse(context, nil, http.StatusNotFound)
		return
	}

	helpers.GinSuccessResponse(context, gin.H{"data": model})
}
