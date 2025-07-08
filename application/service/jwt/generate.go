package jwt

import (
	"compay/application/entities/user"
	"compay/common/config"
	"compay/common/constants"
	"compay/core/utils"
	"time"

	jwtLib "github.com/golang-jwt/jwt/v5"
)

func GenerateAuthToken(userModel *user.Model) (string, error) {
	tokenTTL := constants.UserTokenTtl
	userIdEncrypted, err := userDataToEncryptedString(userModel, userModel.Id)
	if err != nil {
		return "", err
	}

	secretTokenEncrypted, err := utils.EncryptString(userModel.SecretToken)
	if err != nil {
		return "", err
	}

	token := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, jwtLib.MapClaims{
		"id":  userIdEncrypted,                                              // user id
		"stk": secretTokenEncrypted,                                         // secret token
		"iat": time.Now().Unix(),                                            // created at
		"exp": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(), // expired at
	})
	return token.SignedString([]byte(config.AppSecret))
}
