package logpayment

import (
	"compay/common/types"
	"efw/components/model"

	"gorm.io/gorm"
)

type QueryHelper struct {
	model.QueryHelper
	Id        types.LogPaymentId
	PaymentId types.PaymentId
}

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	if helper.PaymentId != 0 {
		chain = chain.Where("payment_id = ?", helper.PaymentId)
	}
	return chain.Order("id ASC")
}
