package merchant

import (
	"compay/common/constants"
	"compay/src/models/user"
	"compay/src/services/processing"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	stateMerchantMu        = sync.Mutex{}
	stateMerchantIsStarted = false
)

func AdminApiStartMerchant(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	stateMerchantMu.Lock()
	defer stateMerchantMu.Unlock()
	if stateMerchantIsStarted {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	processing.StartProcessing()
	stateMerchantIsStarted = true
	efwginhelpers.ResponseSuccess(context, nil)
}

func AdminApiStopMerchant(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	stateMerchantMu.Lock()
	defer stateMerchantMu.Unlock()
	if !stateMerchantIsStarted {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	processing.StopProcessing()
	stateMerchantIsStarted = false
	efwginhelpers.ResponseSuccess(context, nil)
}

func isMerchantStarted() bool {
	stateMerchantMu.Lock()
	defer stateMerchantMu.Unlock()
	return stateMerchantIsStarted
}
