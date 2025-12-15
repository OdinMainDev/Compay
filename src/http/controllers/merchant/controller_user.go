package merchant

import (
	"compay/common/constants"
	"compay/src/models/user"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
)

func UserGetMe(context *gin.Context) {
	userModel := context.MustGet(constants.GinUser).(*user.Model)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": user.DTO(userModel)})
}
