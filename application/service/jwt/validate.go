package jwt

import (
	"compay/application/entities/user"
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/core/utils"
	"strings"

	"github.com/gin-gonic/gin"
	jwtLib "github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(context *gin.Context) error {
	token, err := getJwtToken(context)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwtLib.MapClaims)
	if !ok || !token.Valid {
		return cperrors.ErrorAuthJwtTokenClaimsNotValid
	}

	if claims["id"] == nil || claims["stk"] == nil {
		return cperrors.ErrorAuthJwtTokenIdOrStkIsNull
	}

	userIdEncrypted := claims["id"].(string)
	secretTokenEncrypted := claims["stk"].(string)

	secretToken, err := utils.DecryptString(secretTokenEncrypted)
	if err != nil {
		return cperrors.ErrorAuthJwtTokenFailedDecryptSecretToken
	}

	userIdStr, err := utils.DecryptString(userIdEncrypted)
	if err != nil {
		return cperrors.ErrorAuthJwtTokenFailedDecryptSecretToken
	}
	if !strings.HasSuffix(userIdStr, "-"+secretToken) {
		return cperrors.ErrorAuthJwtTokenFailedDecryptSecretToken
	}

	userModel := context.MustGet(constants.GinUser).(*user.Model)
	if userModel.SecretToken != secretToken {
		return cperrors.ErrorAuthJwtTokenMismatchSecretTokens
	}

	return nil
}
