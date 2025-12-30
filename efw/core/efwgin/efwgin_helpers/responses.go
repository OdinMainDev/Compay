package efwginhelpers

import (
	"efw/core/efwlog"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NotHiddenError error

func ResponseError(context *gin.Context, err any, httpStatus ...int) {
	respStatus := http.StatusUnprocessableEntity
	if len(httpStatus) == 1 {
		respStatus = httpStatus[0]
	}

	if respStatus == http.StatusNotFound {
		context.JSON(respStatus, gin.H{"status": respStatus, "success": false})
		return
	}

	if v, ok := err.(error); ok {
		hideErrorAndResponse(context, v.Error(), respStatus)
		return
	}

	if v, ok := err.(string); ok {
		hideErrorAndResponse(context, v, respStatus)
		return
	}

	if v, ok := err.(NotHiddenError); ok {
		context.JSON(respStatus, gin.H{"status": respStatus, "success": false, "error": v.Error()})
		return
	}

	context.JSON(respStatus, gin.H{"status": respStatus, "success": false, "errors": err})
}

func ResponseSuccess(context *gin.Context, data gin.H) {
	resp := gin.H{"status": 200, "success": true}
	for key, value := range data {
		if key == "status" || key == "success" {
			continue
		}

		resp[key] = value
	}

	context.JSON(http.StatusOK, resp)
}

func hideErrorAndResponse(context *gin.Context, err string, respStatus int) {
	errId := uuid.NewString()
	efwlog.LogGinError(fmt.Sprintf("Id: %s Error: %s", errId, err))
	context.JSON(respStatus, gin.H{"status": respStatus, "success": false, "error": errId})
}
