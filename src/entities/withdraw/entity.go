package withdraw

import (
	"compay/common/types"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[types.WithdrawId] `gorm:"embedded"`
	TrxId                               string                   `json:"trx_id" gorm:"column:trx_id"`
	MerchantId                          types.MerchantId         `json:"merchant_id" gorm:"column:merchant_id"`
	MerchantWalletId                    types.MerchantWalletId   `json:"merchant_wallet_id" gorm:"column:merchant_wallet_id"`
	Currency                            entity.CurrencyCodeAlpha `json:"currency" gorm:"column:currency"`
	IsComplete                          bool                     `json:"is_complete" gorm:"column:is_complete"`
	Address                             types.WithdrawAddress    `json:"address" gorm:"column:address"`
	Amount                              entity.Amount            `json:"amount" gorm:"column:amount"`
	AmountUsd                           entity.Amount            `json:"amount_usd" gorm:"column:amount_usd"`
	Error                               string                   `json:"error" gorm:"column:error"`
}

func (*Entity) TableName() string { return "withdraws" }
