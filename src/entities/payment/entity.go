package payment

import (
	"compay/common/types"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[types.PaymentId] `gorm:"embedded"`
	TrxId                              string                   `json:"trx_id" gorm:"column:trx_id"`
	MerchantId                         types.MerchantId         `json:"merchant_id" gorm:"column:merchant_id"`
	Currency                           entity.CurrencyCodeAlpha `json:"currency" gorm:"column:currency"`
	Status                             types.PaymentStatus      `json:"status" gorm:"column:status"`
	Address                            types.PaymentAddress     `json:"address" gorm:"column:address"`
	Amount                             entity.Amount            `json:"amount" gorm:"column:amount"`
	AmountUsd                          entity.Amount            `json:"amount_usd" gorm:"column:amount_usd"`
	ReceivedAmount                     entity.Amount            `json:"received_amount" gorm:"column:received_amount"`
	ReceivedAmountUsd                  entity.Amount            `json:"received_amount_usd" gorm:"column:received_amount_usd"`
	RequiredConfirmations              types.Confirmation       `json:"required_confirmations" gorm:"column:required_confirmations"`
	ExpiresAt                          int64                    `json:"expires_at" gorm:"column:expires_at"`
}

func (*Entity) TableName() string { return "payments" }
