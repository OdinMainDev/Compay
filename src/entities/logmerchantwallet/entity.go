package logmerchantwallet

import (
	"compay/common/types"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[types.LogMerchantWalletId] `gorm:"embedded"`
	MerchantId                                   types.MerchantId        `json:"merchant_id" gorm:"column:merchant_id"`
	PaymentId                                    types.PaymentId         `json:"payment_id" gorm:"column:payment_id"`
	WithdrawId                                   types.PaymentId         `json:"withdraw_id" gorm:"column:withdraw_id"`
	PaymentTrxId                                 string                  `json:"payment_trx_id" gorm:"column:payment_trx_id"`
	WithdrawTrxId                                string                  `json:"withdraw_trx_id" gorm:"column:withdraw_trx_id"`
	Currency                                     types.CurrencyCodeAlpha `json:"currency" gorm:"column:currency"`
	Amount                                       types.Amount            `json:"amount" gorm:"column:amount"`
	AmountUsd                                    types.Amount            `json:"amount_usd" gorm:"column:amount_usd"`
}

func (*Entity) TableName() string { return "log_merchant_wallets" }
