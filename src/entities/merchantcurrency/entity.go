package merchantcurrency

import (
	"compay/common/types"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[types.MerchantCurrencyId] `gorm:"embedded"`
	MerchantId                                  types.MerchantId         `json:"merchant_id" gorm:"column:merchant_id"`
	Currency                                    entity.CurrencyCodeAlpha `json:"currency" gorm:"column:currency"`
	CreatorId                                   entity.UserId            `json:"creator_id" gorm:"column:creator_id"`
	MinAmount                                   entity.Amount            `json:"min_amount" gorm:"column:min_amount"`
	MaxAmount                                   entity.Amount            `json:"max_amount" gorm:"column:max_amount"`
	MinAmountWithdraw                           entity.Amount            `json:"min_amount_withdraw" gorm:"column:min_amount_withdraw"`
	MaxAmountWithdraw                           entity.Amount            `json:"max_amount_withdraw" gorm:"column:max_amount_withdraw"`
	Confirms                                    types.Confirmation       `json:"confirms" gorm:"column:confirms"`
	InvoiceTime                                 int64                    `json:"invoice_time" gorm:"column:invoice_time"`
	EstimatedFeeBlockCount                      int                      `json:"estimated_fee_block_count" gorm:"column:estimated_fee_block_count"`
	FeeType                                     types.FeeType            `json:"fee_type" gorm:"column:fee_type"`
	FeeFixed                                    types.FeeAmount          `json:"fee_fixed" gorm:"column:fee_fixed"`
	FeePercent                                  types.FeePercentAmount   `json:"fee_percent" gorm:"column:fee_percent"`
}

func (*Entity) TableName() string { return "merchant_currencies" }
