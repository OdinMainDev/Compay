package efwjwt

import (
	"efw/config"
	"efw/utils"
	"time"

	jwtLib "github.com/golang-jwt/jwt/v5"
)

type JwtData struct {
	Key              string
	Value            any
	IsEncryptedValue bool
}

func GenerateJWT(userData []JwtData, secretToken string, jwtTTL int) (string, error) {
	secretTokenEncrypted, err := utils.EncryptString(secretToken)
	if err != nil {
		return "", err
	}

	jwtData := jwtLib.MapClaims{
		"stk": secretTokenEncrypted,                                       // secret token
		"iat": time.Now().Unix(),                                          // created at
		"exp": time.Now().Add(time.Second * time.Duration(jwtTTL)).Unix(), // expired at
	}

	for _, val := range userData {
		if !val.IsEncryptedValue {
			jwtData[val.Key] = val.Value
			continue
		}
		encryptedVal, err := encryptUserData(val.Value, secretToken)
		if err != nil {
			return "", err
		}
		jwtData[val.Key+EncryptedKeySuffix] = encryptedVal
	}

	token := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, jwtData)
	return token.SignedString(config.AppSecret)
}
