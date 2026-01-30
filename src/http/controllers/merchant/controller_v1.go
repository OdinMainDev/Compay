package merchant

import (
	"compay/common/types"
	"compay/src/models/merchant"
	"efw/utils"
)

type signData struct {
	token    string
	currency types.CurrencyCodeAlpha
	sign     string
}

func signVerify(merchantModel *merchant.Model, data signData) bool {
	sign := "token:" + data.token + ";currency:" + string(data.currency) + ";secret:" + string(merchantModel.SecretKey)
	return data.sign == utils.GetHashString(sign)
}
