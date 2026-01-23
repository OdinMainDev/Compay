package casts

import (
	"github.com/gin-gonic/gin"
)

func GetRequestBodyString(context *gin.Context) string {
	return string(GetRequestBody(context))
}

func GetRequestBody(context *gin.Context) []byte {
	body := []byte{}
	if cb, ok := context.Get(gin.BodyBytesKey); ok {
		if cbb, ok := cb.([]byte); ok {
			body = cbb
		}
		return body
	}

	return []byte(string("body_null"))
}
