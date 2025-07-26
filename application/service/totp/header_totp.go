package totp

import (
	"compay/application/entities/merchantuser"
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/core/utils"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateHeaderToken(userModel *merchantuser.UserModel) (string, error) {
	now := time.Now().Unix()
	rawToken := fmt.Sprintf("%d-%d", userModel.GetId(), now)

	encrypted, err := utils.EncryptString(rawToken)
	if err != nil {
		return "", err
	}

	return encrypted, nil
}

func ValidateRequest(context *gin.Context) error {
	userModelData, ok := context.Get(constants.GinUser)
	if !ok {
		return cperrors.ErrorAuthTotpBadRequest
	}

	userModel := userModelData.(*merchantuser.UserModel)
	if !userModel.IsTwofaEnabled {
		return nil
	}

	tokenEncrypted := context.Request.Header.Get("C-Totp")
	if tokenEncrypted == "" {
		return cperrors.ErrorAuthTotpTokenRequired
	}

	token, err := utils.DecryptString(tokenEncrypted)
	if err != nil {
		return cperrors.ErrorAuthTotpDecryptDataError
	}

	lastCheckAtStr := strings.ReplaceAll(token, fmt.Sprint(userModel.GetId(), "-"), "")
	lastCheckAt, err := strconv.ParseInt(lastCheckAtStr, 10, 64)
	if err != nil {
		return cperrors.ErrorAuthTotpInvalidToken
	}

	if time.Now().Unix()-lastCheckAt > 1200 {
		return cperrors.ErrorAuthTotpTokenExpired
	}

	return nil
}
