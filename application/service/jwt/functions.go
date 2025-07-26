package jwt

import (
	"compay/common/config"
	"compay/common/cperrors"
	"strings"

	"github.com/gin-gonic/gin"
	jwtLib "github.com/golang-jwt/jwt/v5"
)

func getJwtToken(context *gin.Context) (*jwtLib.Token, error) {
	authHeaderToken := GetAuthHeaderTokenFromRequest(context)
	if authHeaderToken == "" {
		return nil, cperrors.ErrorJwtBadHeaders
	}
	return getJwtFromAuthHeader(authHeaderToken)
}

func GetAuthHeaderTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}

func getJwtFromAuthHeader(authHeaderToken string) (*jwtLib.Token, error) {
	token, err := jwtLib.Parse(authHeaderToken, func(token *jwtLib.Token) (interface{}, error) {
		if token == nil {
			return nil, cperrors.ErrorJwtBadToken
		}

		if _, ok := token.Method.(*jwtLib.SigningMethodHMAC); !ok {
			return nil, cperrors.ErrorJwtBadToken2
		}

		return []byte(config.AppSecret), nil
	})

	if err != nil {
		return nil, cperrors.ErrorJwtBadToken3
	}

	return token, nil
}
