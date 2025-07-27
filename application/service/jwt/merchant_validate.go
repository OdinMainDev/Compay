package jwt

import (
	"compay/application/entities/basemodel"
	"compay/application/entities/merchantuser"
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/core/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	jwtLib "github.com/golang-jwt/jwt/v5"
)

func ValidateMerchantJWT(context *gin.Context) error {
	token, err := getJwtToken(context)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwtLib.MapClaims)
	if !ok || !token.Valid {
		return cperrors.ErrorJwtBadMerchantToken
	}

	if claims["id"] == nil {
		return cperrors.ErrorJwtBadMerchantTokenData
	}
	userIdEncrypted := claims["id"].(string)
	secretTokenEncrypted := claims["stk"].(string)

	secretToken, err := utils.DecryptString(secretTokenEncrypted)
	if err != nil {
		return cperrors.ErrorJwtBadMerchantTokenSecretTokenm
	}

	userIdStr, err := utils.DecryptString(userIdEncrypted)
	if err != nil {
		return cperrors.ErrorJwtBadMerchantTokenUserId
	}
	if !strings.HasSuffix(userIdStr, "-"+secretToken) {
		return cperrors.ErrorJwtBadMerchantTokenUserId2
	}
	userIdStr = strings.ReplaceAll(userIdStr, "-"+secretToken, "")

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return cperrors.ErrorJwtBadMerchantTokenUserId3
	}

	userModel := merchantuser.GetModel(merchantuser.QueryHelper{
		Id: userId,
	})
	if userModel == nil || userModel.GetId() == 0 || userModel.SecretToken != basemodel.EncryptedString(secretToken) {
		return cperrors.ErrorJwtBadMerchantTokenUserId4
	}

	context.Set(constants.GinSecretToken, secretToken)
	context.Set(constants.GinUserId, userId)
	context.Set(constants.GinUser, userModel)
	return nil
}
