package efwjwt

import (
	"efw/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func getJwtToken(context *gin.Context) (*jwt.Token, error) {
	authHeaderToken := getAuthHeaderTokenFromRequest(context)
	if authHeaderToken == "" {
		return nil, ErrorJwtNilAuthorizationHeader
	}
	return parseAuthHeaderToken(authHeaderToken)
}

func getAuthHeaderTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}

func parseAuthHeaderToken(authHeaderToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(authHeaderToken, func(token *jwt.Token) (interface{}, error) {
		if token == nil {
			return nil, ErrorJwtInvalidToken
		}

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrorJwtUnexpectedSigningMethod
		}

		return []byte(config.AppSecret), nil
	})

	if err != nil {
		return nil, ErrorJwtUnexpectedParsingError(err)
	}

	return token, nil
}
