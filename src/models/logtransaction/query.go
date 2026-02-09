package logtransaction

import (
	"compay/common/types"
	"context"
	"efw"
	"efw/components/model"

	"gorm.io/gorm"
)

type QueryHelper struct {
	model.QueryHelper
	Id        types.LogPaymentId
	PaymentId types.PaymentId
	TxId      types.PaymentTxId
}

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	if helper.PaymentId != 0 {
		chain = chain.Where("payment_id = ?", helper.PaymentId)
	}
	if helper.TxId != "" {
		chain = chain.Where("tx_id = ?", helper.TxId)
	}
	return chain.Order("id ASC")
}

func UpdateStatus(logId types.LogPaymentId) {
	efw.DBTransaction[Model]().
		Where("id = ?", logId).
		Update(context.Background(), "is_complete", true)

}
