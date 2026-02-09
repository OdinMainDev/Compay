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
	Id               types.LogTransactionId
	PaymentId        types.PaymentId
	MerchantWalletId types.MerchantWalletId
	TxId             types.PaymentTxId
	IsComplete       bool
	IsUnSpent        bool
}

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	if helper.PaymentId != 0 {
		chain = chain.Where("payment_id = ?", helper.PaymentId)
	}
	if helper.MerchantWalletId != 0 {
		chain = chain.Where("merchant_wallet_id = ?", helper.MerchantWalletId)
	}
	if helper.TxId != "" {
		chain = chain.Where("tx_id = ?", helper.TxId)
	}
	if helper.IsComplete {
		chain = chain.Where("is_complete = ?", true)
	}
	if helper.IsUnSpent {
		chain = chain.Where("is_spent = ?", false)
	}
	return chain.Order("id ASC")
}

func UpdateStatus(logId types.LogTransactionId) {
	efw.DBTransaction[Model]().
		Where("id = ?", logId).
		Update(context.Background(), "is_complete", true)

}
