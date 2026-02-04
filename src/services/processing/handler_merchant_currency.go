package processing

import (
	"compay/src/models/merchantcurrency"
)

func (that *Processing) getMerchantCurrency() *merchantcurrency.Model {
	if that.merchantCurrencyModel != nil {
		return that.merchantCurrencyModel
	}

	merchantCurrencyModel := merchantcurrency.GetModel(merchantcurrency.QueryHelper{
		MerchantId: that.getMerchant().GetId(),
		Currency:   that.getCurrency().CodeAlpha,
	})

	if merchantCurrencyModel != nil {
		that.merchantCurrencyModel = merchantCurrencyModel
		return that.merchantCurrencyModel
	}

	merchantCurrencyModel = merchantcurrency.DefaultCurrency(that.getCurrency().CodeAlpha, that.getMerchant().GetId(), that.getMerchant().CreatorId)
	merchantcurrency.Save(merchantCurrencyModel, merchantcurrency.QueryHelper{})

	that.merchantCurrencyModel = merchantCurrencyModel
	return that.merchantCurrencyModel
}
