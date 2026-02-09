package merchantwallet

import (
	"compay/common/types"
	"context"
	"efw"
	"efw/components/entity"
	"efw/components/model"
	"efw/core/efwlog"

	"gorm.io/gorm"
)

type QueryHelper struct {
	model.QueryHelper
	Id                  types.MerchantWalletId
	MerchantId          types.MerchantId
	DontCheckMerchantId bool
	MerchantIds         []types.MerchantId
	PaymentId           types.PaymentId
	Currency            entity.CurrencyCodeAlpha
	WithPayment         bool
}

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {

	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	if !helper.DontCheckMerchantId {
		chain = chain.Where("merchant_id = ?", helper.MerchantId)
	}
	if len(helper.MerchantIds) != 0 {
		chain = chain.Where("merchant_id IN (?)", helper.MerchantIds)
	}
	if helper.PaymentId != 0 {
		chain = chain.Where("payment_id = ?", helper.PaymentId)
	}
	if helper.Currency != "" {
		chain = chain.Where("currency = ?", helper.Currency)
	}
	if helper.WithPayment {
		chain = chain.Preload("Payment", func(db gorm.PreloadBuilder) error { return nil })
	}
	return chain.Order("id DESC")
}

func IncreaseSentAmount(merchantWalletId types.MerchantWalletId, amount entity.Amount) {
	_, err := efw.DBTransaction[any]().Table("merchant_wallets").
		Where("id = ?", merchantWalletId).
		Updates(context.Background(), map[string]any{
			"sent_amount": gorm.Expr("sent_amount + ?", amount),
			"amount":      gorm.Expr("amount - ?", amount),
		})
	if err != nil {
		efwlog.LogByLoggerName("gorm_error", 0, efwlog.FmtGormError(0, 0, err.Error(), "IncreaseSentAmount"))
	}
}
