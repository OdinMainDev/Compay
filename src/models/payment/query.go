package payment

import (
	"compay/common/types"
	"efw/components/model"

	"gorm.io/gorm"
)

type QueryHelper struct {
	model.QueryHelper
	Id         types.PaymentId
	TrxId      string
	Currency   types.CurrencyCodeAlpha
	MerchantId types.MerchantId
}

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	chain = chain.Where("merchant_id = ?", helper.MerchantId)

	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	if helper.TrxId != "" {
		chain = chain.Where("trx_id = ?", helper.TrxId)
	}
	if helper.Currency != "" {
		chain = chain.Where("currency = ?", helper.Currency)
	}
	return chain.Order("id DESC")
}
