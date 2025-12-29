package efwcontroller

import (
	"efw/core/efwgin"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Deletable interface {
	GetIdInt64() int64
}

func DeleteRecord[T Deletable, T2 any](
	context *gin.Context,
	getModel func(queryHelper T2, additionalCacheKeys ...int64) T,
	queryHelper T2,
	additionalCacheKeys ...int64,
) (bool, int64) {
	var input RequestGetById
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return false, 0
	}

	model := getModel(queryHelper, additionalCacheKeys...)
	fmt.Println(model)
	// if model {
	// 	helpers.GinErrorResponse(context, nil, http.StatusNotFound)
	// 	return false, 0
	// }

	// model.Delete(model)
	// helpers.GinSuccessResponse(context, nil)
	return true, model.GetId()
}
