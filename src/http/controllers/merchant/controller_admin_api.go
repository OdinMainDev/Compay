package merchant

import (
	"compay/common/constants"
	"compay/src/models/user"
	"compay/src/services/cconverter"
	"compay/src/services/processing"
	"efw/core/efwcache"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	startMerchantMu        = sync.Mutex{}
	startMerchantIsStarted = false
)

func AdminApiFeeTypes(context *gin.Context) {
	efwginhelpers.ResponseSuccess(context, gin.H{"data": constants.FeeTypes})
}

func AdminApiClearCache(context *gin.Context) {
	efwcache.ClearCache()
	efwginhelpers.ResponseSuccess(context, nil)
}

func AdminApiInitMerchant(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	startMerchantMu.Lock()
	defer startMerchantMu.Unlock()
	if startMerchantIsStarted {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	cconverter.Start()
	processing.StartProcessing()
	startMerchantIsStarted = true
	efwginhelpers.ResponseSuccess(context, nil)
}
