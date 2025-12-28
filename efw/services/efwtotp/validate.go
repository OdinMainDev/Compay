package efwtotp

import (
	"efw/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Validate(context *gin.Context, username string) error {
	tokenEncrypted := context.Request.Header.Get(TotpHeaderName)
	if tokenEncrypted == "" {
		return ErrorTotpNilHeader
	}

	token, err := utils.DecryptString(tokenEncrypted)
	if err != nil {
		return err
	}

	lastCheckAtStr := strings.ReplaceAll(token, username+"-", "")
	lastCheckAt, err := strconv.ParseInt(lastCheckAtStr, 10, 64)
	if err != nil {
		return ErrorTotpTokenParseError(err)
	}

	if time.Now().Unix()-lastCheckAt > TotpSessionTime {
		return ErrorTotpTokenExpired
	}

	return nil
}
