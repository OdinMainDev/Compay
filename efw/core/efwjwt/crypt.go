package efwjwt

import (
	"efw/utils"
	"fmt"
)

func encryptUserData(val any, secretToken string) (string, error) {
	return utils.EncryptString(fmt.Sprint(val, "-", secretToken))
}
