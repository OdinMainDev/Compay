package merchant

import (
	"compay/application/entities/merchantuser"
	"compay/application/helpers"
	"compay/application/service/esvalidator"
	"compay/application/service/jwt"
	"compay/common/cperrors"
	"compay/core/utils"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type requestLogin struct {
	Token string `json:"token" binding:"required"`
}

func AuthLogin(context *gin.Context) {
	var input requestLogin
	if err := context.ShouldBindJSON(&input); err != nil {
		helpers.GinErrorResponse(context, esvalidator.ValidatorResponse(err))
		return
	}

	userId, authTime, err := parseAuthToken(input.Token)
	if err != nil {
		helpers.GinErrorResponse(context, err)
		return
	}

	if time.Now().Unix()-authTime >= 45 {
		helpers.GinErrorResponse(context, cperrors.ErrorAuthMerchantExpiredAuthToken)
		return
	}

	userModel := merchantuser.GetModel(merchantuser.QueryUserHelper{
		Id: userId,
	})

	if userModel == nil || userModel.GetId() == 0 || !userModel.IsApproved {
		helpers.GinErrorResponse(context, cperrors.ErrorAuthMerchantInvalidUser)
		return
	}

	// TODO: вернуть 1800
	token, err := jwt.GenerateMerchantJWT(userModel, 3600000000)
	if err != nil {
		helpers.GinErrorResponse(context, err)
		return
	}

	helpers.GinSuccessResponse(context, gin.H{"token": token})
}

var authTokenRegex = regexp.MustCompile(`^do_auth_user_as_admin_(\d+)_(\d+)$`)

func parseAuthToken(str string) (int64, int64, error) {
	decrypt, err := utils.DecryptString(str)
	if err != nil {
		return 0, 0, cperrors.ErrorAuthMerchantBadAuthToken
	}

	matches := authTokenRegex.FindStringSubmatch(decrypt)
	if matches == nil {
		return 0, 0, cperrors.ErrorAuthMerchantInvalidAuthToken
	}

	id1, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return 0, 0, cperrors.ErrorAuthMerchantInvalidAuthToken2
	}
	id2, err := strconv.ParseInt(matches[2], 10, 64)
	if err != nil {
		return 0, 0, cperrors.ErrorAuthMerchantInvalidAuthToken3
	}
	return id1, id2, nil
}
