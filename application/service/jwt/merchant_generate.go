package jwt

import (
	"compay/application/entities/merchantuser"
	"compay/common/config"
	"compay/common/cperrors"
	"compay/core/utils"
	"fmt"
	"time"

	jwtLib "github.com/golang-jwt/jwt/v5"
)

func userDataToEncryptedString(userModel *merchantuser.Model, data any) (string, error) {
	return utils.EncryptString(fmt.Sprint(data, "-", userModel.SecretToken))
}

func GenerateMerchantJWT(userModel *merchantuser.Model, tokenTTL int) (string, error) {
	userIdEncrypted, err := userDataToEncryptedString(userModel, userModel.Id)
	if err != nil {
		return "", cperrors.ErrorJwtBadMerchantEncryption
	}

	secretTokenEncrypted, err := utils.EncryptString(string(userModel.SecretToken))
	if err != nil {
		return "", cperrors.ErrorJwtBadMerchantEncryption2
	}

	token := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, jwtLib.MapClaims{
		"id":  userIdEncrypted,                                              // user id
		"stk": secretTokenEncrypted,                                         // secret token
		"im":  *userModel.GetMerchantId() > 0,                               // is merchant
		"iat": time.Now().Unix(),                                            // created at
		"exp": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(), // expired at
	})
	resp, err := token.SignedString([]byte(config.AppSecret))
	if err != nil {
		return "", cperrors.ErrorJwtBadMerchantTokenSign
	}
	return resp, nil
}
