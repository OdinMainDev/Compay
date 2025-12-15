package efwjwt

import (
	"efw/utils"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(context *gin.Context) error {
	token, err := getJwtToken(context)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims == nil {
		return ErrorJwtUnexpectedClaimsType
	}

	if claims["iat"] == nil || claims["exp"] == nil || claims["stk"] == nil {
		return ErrorJwtInvalidJwtClaims
	}

	secretToken, err := utils.DecryptString(fmt.Sprint(claims["stk"]))
	if err != nil {
		return err
	}

	jwtClaimsData := make(JwtClaimsData)
	for key, val := range claims {
		if utils.InArray(key, []string{"iat", "exp", "stk"}) {
			continue
		}

		keyNew := key
		isEncrypted := false
		if strings.HasSuffix(keyNew, EncryptedKeySuffix) {
			keyNew = strings.TrimSuffix(keyNew, EncryptedKeySuffix)
			isEncrypted = true
		}

		if !isEncrypted {
			jwtClaimsData[keyNew] = fmt.Sprint(val)
			continue
		}

		valDecrypted, err := utils.DecryptString(fmt.Sprint(val))
		if err != nil {
			return err
		}

		jwtClaimsData[keyNew] = strings.ReplaceAll(valDecrypted, "-"+secretToken, "")
	}

	context.Set(GinJwtClaimsData, jwtClaimsData)
	return nil
}
