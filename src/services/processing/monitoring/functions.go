package monitoring

import (
	"compay/common/types"
	"compay/src/models/merchantcurrency"
)

func (that *monitoring) getMerchantCurrencyModel(merchantCurrencyModel *merchantcurrency.Model, merchantId types.MerchantId, codeAlpha types.CurrencyCodeAlpha) *merchantcurrency.Model {
	if merchantCurrencyModel != nil {
		return merchantCurrencyModel
	}
	merchantCurrencyModel = merchantcurrency.GetModel(merchantcurrency.QueryHelper{
		MerchantId: merchantId,
		Currency:   codeAlpha,
	})
	return merchantCurrencyModel
}
