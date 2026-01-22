package casts

import (
	"efw/components/entity"

	"github.com/gin-gonic/gin"
)

type requestFingerPrint struct {
	FingerPrint string `json:"fp"`
}

func GetClientFingerPrint(context *gin.Context) entity.FingerPrint {
	var input requestFingerPrint
	if err := context.ShouldBindBodyWithJSON(&input); err != nil {
		return "fingerprint_notset"
	}

	return entity.FingerPrint(input.FingerPrint)
}
