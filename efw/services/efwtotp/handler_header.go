package efwtotp

import (
	"efw/utils"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateHeaderToken(username string) (string, error) {
	now := time.Now().Unix()
	rawToken := fmt.Sprintf("%s-%d", username, now)

	encrypted, err := utils.EncryptString(rawToken)
	if err != nil {
		return "", err
	}

	return encrypted, nil
}

func ValidateRequest(context *gin.Context, username string) error {
	tokenEncrypted := context.Request.Header.Get(TotpHeaderName)
	if tokenEncrypted == "" {
		return ErrorNilHeader
	}

	token, err := utils.DecryptString(tokenEncrypted)
	if err != nil {
		return err
	}

	lastCheckAtStr := strings.ReplaceAll(token, username+"-", "")
	lastCheckAt, err := strconv.ParseInt(lastCheckAtStr, 10, 64)
	if err != nil {
		return ErrorTokenParseError(err)
	}

	if time.Now().Unix()-lastCheckAt > TotpSessionTime {
		return ErrorTokenExpired
	}

	return nil
}
