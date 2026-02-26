package merchant

import (
	"compay/common/constants"
	"compay/src/models/user"
	"efw/core/efwcache"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminApiFeeTypes(context *gin.Context) {
	efwginhelpers.ResponseSuccess(context, gin.H{"data": constants.FeeTypes})
}

func AdminApiClearCache(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	efwcache.ClearCache()
	efwginhelpers.ResponseSuccess(context, nil)
}
