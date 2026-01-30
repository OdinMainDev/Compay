package merchantcurrency

import (
	"compay/common/types"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[types.MerchantCurrencyId] `gorm:"embedded"`
	MerchantId                                  types.MerchantId        `json:"merchant_id" gorm:"column:merchant_id"`
	Currency                                    types.CurrencyCodeAlpha `json:"currency" gorm:"column:currency"`
	MinAmount                                   types.AmountInt         `json:"min_amount" gorm:"column:min_amount"`
	MaxAmount                                   types.AmountInt         `json:"max_amount" gorm:"column:max_amount"`
	MinConfirms                                 types.Confirmation      `json:"min_confirms" gorm:"column:min_confirms"`
	FeeType                                     types.FeeType           `json:"fee_type" gorm:"column:min_confirms"`
	FeeFixed                                    types.Confirmation      `json:"fee_fixed" gorm:"column:fee_fixed"`
	FeePercent                                  types.Confirmation      `json:"fee_percent" gorm:"column:fee_percent"`
}

func (*Entity) TableName() string {
	return "merchant_currencies"
}
