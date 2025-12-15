package efwtotp

import (
	"efw/utils"
	"fmt"
	"time"
)

func GenerateToken(username string) (string, error) {
	now := time.Now().Unix()
	rawToken := fmt.Sprintf("%s-%d", username, now)

	encrypted, err := utils.EncryptString(rawToken)
	if err != nil {
		return "", err
	}

	return encrypted, nil
}
