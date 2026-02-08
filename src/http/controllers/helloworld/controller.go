package helloworld

import (
	"compay/src/models/logauth"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
)

func Hello(context *gin.Context) {
	efwginhelpers.ResponseSuccess(context, gin.H{"data": "Hello World"})
}

func Hello2(context *gin.Context) {
	logAuthModel := logauth.GetModel(logauth.QueryHelper{Id: 1})
	if logAuthModel != nil {
		efwginhelpers.ResponseSuccess(context, gin.H{"data": logAuthModel})
		return
	}

	logAuthModel = &logauth.Model{}
	logAuthModel.UserId = 1
	logAuthModel.MerchantId = 2
	logAuthModel.Action = "testaction"
	logAuthModel.Url = "testurl"
	logAuthModel.FingerPrint = "testfp"
	logAuthModel.Ip = "testip"
	logAuthModel.UserAgent = "testua"
	logauth.Save(logAuthModel, logauth.QueryHelper{})

	efwginhelpers.ResponseSuccess(context, gin.H{"data": logAuthModel})
}
