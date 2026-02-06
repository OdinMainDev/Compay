package monitoring

import (
	"compay/common/types"
	"compay/src/models/merchantcurrency"
	"efw/components/entity"
)

func (that *monitoring) getMerchantCurrencyModel(merchantCurrencyModel *merchantcurrency.Model, merchantId types.MerchantId, codeAlpha entity.CurrencyCodeAlpha) *merchantcurrency.Model {
	if merchantCurrencyModel != nil {
		return merchantCurrencyModel
	}
	merchantCurrencyModel = merchantcurrency.GetModel(merchantcurrency.QueryHelper{
		MerchantId: merchantId,
		Currency:   codeAlpha,
	})
	return merchantCurrencyModel
}
