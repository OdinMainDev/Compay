package casts

import (
	"efw/components/entity"

	"github.com/gin-gonic/gin"
)

func GetRequestUri(context *gin.Context) entity.Url {
	return entity.Url(context.Request.RequestURI)
}
