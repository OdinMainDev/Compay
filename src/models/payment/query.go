package payment

import (
	"compay/common/types"
	paymentEntity "compay/src/entities/payment"
	"context"
	"efw"
	"efw/components/model"

	"gorm.io/gorm"
)

type QueryHelper struct {
	model.QueryHelper
	Id                  types.PaymentId
	TrxId               string
	DontCheckMerchantId bool
	MerchantId          types.MerchantId
	Statuses            []types.PaymentStatus
	Currency            types.CurrencyCodeAlpha
}

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {

	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	if !helper.DontCheckMerchantId {
		chain = chain.Where("merchant_id = ?", helper.MerchantId)
	}
	if helper.TrxId != "" {
		chain = chain.Where("trx_id = ?", helper.TrxId)
	}
	if helper.Currency != "" {
		chain = chain.Where("currency = ?", helper.Currency)
	}
	if len(helper.Statuses) != 0 {
		chain = chain.Where("status IN (?)", helper.Statuses)
	}
	return chain.Order("id DESC")
}

func IncreaseAmount(paymentId types.PaymentId, amount types.Amount) {
	efw.DBTransaction[paymentEntity.Entity]().
		Where("id = ?", paymentId).
		Update(context.Background(), "received_amount", gorm.Expr("received_amount + ?", amount))
}

func UpdateStatus(paymentId types.PaymentId, status types.PaymentStatus) {
	efw.DBTransaction[paymentEntity.Entity]().
		Where("id = ?", paymentId).
		Update(context.Background(), "status", status)

}
