package jwt

import (
	"compay/application/entities/user"
	"compay/common/config"
	"compay/common/cperrors"
	"compay/core/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	jwtLib "github.com/golang-jwt/jwt/v5"
)

func userDataToEncryptedString(userModel *user.Model, data any) (string, error) {
	return utils.EncryptString(fmt.Sprint(data, "-", userModel.SecretToken))
}

func getJwtToken(context *gin.Context) (*jwtLib.Token, error) {
	authHeaderToken := getAuthHeaderTokenFromRequest(context)
	if authHeaderToken == "" {
		return nil, cperrors.ErrorAuthJwtTokenIsNull
	}
	return getJwtFromAuthHeader(authHeaderToken)
}

func getAuthHeaderTokenFromRequest(context *gin.Context) string {
	token := context.Request.Header.Get("A-Token")
	return token
}

func getJwtFromAuthHeader(authHeaderToken string) (*jwtLib.Token, error) {
	token, err := jwtLib.Parse(authHeaderToken, func(token *jwtLib.Token) (interface{}, error) {
		if token == nil {
			return nil, cperrors.ErrorAuthJwtTokenInvalid
		}

		if _, ok := token.Method.(*jwtLib.SigningMethodHMAC); !ok {
			return nil, cperrors.ErrorAuthJwtTokenInvalidHashMethod
		}

		return []byte(config.AppSecret), nil
	})

	if err != nil {
		return nil, cperrors.ErrorAuthJwtTokenInvalid
	}

	return token, nil
}
