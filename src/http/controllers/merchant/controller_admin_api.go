package merchant

import (
	"compay/common/constants"
	"efw/core/efwcache"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
)

func AdminApiFeeTypes(context *gin.Context) {
	efwginhelpers.ResponseSuccess(context, gin.H{"data": constants.FeeTypes})
}

func AdminApiClearCache(context *gin.Context) {
	efwcache.ClearCache()
	efwginhelpers.ResponseSuccess(context, nil)
}
