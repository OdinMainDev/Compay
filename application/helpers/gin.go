package helpers

import (
	"compay/core/eslog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinErrorResponse(context *gin.Context, err any, httpStatus ...int) {
	respStatus := http.StatusUnprocessableEntity
	if len(httpStatus) == 1 {
		respStatus = httpStatus[0]
	}

	if v, ok := err.(error); ok {
		context.JSON(respStatus, gin.H{"status": respStatus, "success": false, "error": v.Error()})
		return
	}

	if v, ok := err.(string); ok {
		context.JSON(respStatus, gin.H{"status": respStatus, "success": false, "error": v})
		return
	}

	context.JSON(respStatus, gin.H{"status": respStatus, "success": false, "errors": err})
}

func GinSuccessResponse(context *gin.Context, data gin.H) {
	resp := gin.H{"status": 200, "success": true}
	for key, value := range data {
		if key == "status" || key == "success" {
			eslog.PanicLog("FATAL ERROR ON GIN SUCCESS RESPONSE WITH KEY: ", key, " VALUE: ", value)
		}
		resp[key] = value
	}

	context.JSON(http.StatusOK, resp)
}
