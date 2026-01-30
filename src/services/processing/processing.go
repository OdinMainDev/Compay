package processing

import (
	"compay/src/models/currency"
	"compay/src/models/merchant"
	"compay/src/models/merchantcurrency"
)

type Processing struct {
	merchantModel         *merchant.Model
	merchantCurrencyModel *merchantcurrency.Model
	currencyModel         *currency.Model
}

func NewService(merchantModel *merchant.Model, currencyModel *currency.Model) *Processing {
	return &Processing{
		merchantModel: merchantModel,
		currencyModel: currencyModel,
	}
}

func (that *Processing) getMerchant() *merchant.Model { return that.merchantModel }
func (that *Processing) getCurrency() *currency.Model { return that.currencyModel }
