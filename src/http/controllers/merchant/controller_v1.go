package merchant

import (
	"efw/components/entity"
	"efw/utils"
)

type signData struct {
	token    string
	currency entity.CurrencyCodeAlpha
	sign     string
}

func signVerify(secretKey string, data signData) bool {
	sign := "token:" + data.token + ";currency:" + string(data.currency) + ";secret:" + secretKey
	return data.sign == utils.GetHashString(sign)
}
