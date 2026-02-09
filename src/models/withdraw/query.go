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
	IsInComplete     bool

	// TrxId               string
	DontCheckMerchantId bool
	MerchantId          types.MerchantId
	MerchantIds         []types.MerchantId
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
	if !helper.DontCheckMerchantId {
		chain = chain.Where("merchant_id = ?", helper.MerchantId)
	}
	if len(helper.MerchantIds) != 0 {
		chain = chain.Where("merchant_id IN (?)", helper.MerchantIds)
	}
	if helper.IsComplete {
		chain = chain.Where("is_complete = ?", true)
	}
	if helper.IsInComplete {
		chain = chain.Where("is_complete = ?", false)
	}
	return chain.Order("id DESC")
}
