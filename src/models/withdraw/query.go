package withdraw

import (
	"compay/common/types"
	"efw/components/model"

	"gorm.io/gorm"
)

type QueryHelper struct {
	model.QueryHelper
	Id               types.WithdrawId
	MerchantWalletId types.MerchantWalletId
	IsComplete       bool
	// TrxId               string
	// DontCheckMerchantId bool
	// MerchantId          types.MerchantId
	// MerchantIds         []types.MerchantId
	// Statuses            []types.PaymentStatus
	// Currency            entity.CurrencyCodeAlpha
}

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {

	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	if helper.MerchantWalletId != 0 {
		chain = chain.Where("merchant_wallet_id = ?", helper.MerchantWalletId)
	}
	if helper.IsComplete {
		chain = chain.Where("is_complete = ?", true)
	}
	return chain.Order("id DESC")
}
