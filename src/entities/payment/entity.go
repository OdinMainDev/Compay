package payment

import (
	"compay/common/types"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[types.PaymentId] `gorm:"embedded"`
	TrxId                              string                  `json:"trx_id" gorm:"column:trx_id"`
	MerchantId                         types.MerchantId        `json:"merchant_id" gorm:"column:merchant_id"`
	Currency                           types.CurrencyCodeAlpha `json:"currency" gorm:"column:currency"`
	Status                             types.PaymentStatus     `json:"status" gorm:"column:status"`
	Address                            string                  `json:"address" gorm:"column:address"`
	Amount                             types.Amount            `json:"amount" gorm:"column:amount"`
	AmountUsd                          types.Amount            `json:"amount_usd" gorm:"column:amount_usd"`
	RecievedAmount                     types.Amount            `json:"received_amount" gorm:"column:received_amount"`
	RecievedAmountUsd                  types.Amount            `json:"received_amount_usd" gorm:"column:received_amount_usd"`
	PrivateKey                         entity.EncryptedString  `json:"private_key" gorm:"column:private_key"`
	PrivateKeyWif                      entity.EncryptedString  `json:"private_key_wif" gorm:"column:private_key_wif"`
	PublicKey                          entity.EncryptedString  `json:"public_key" gorm:"column:public_key"`
	PublicKeyFull                      entity.EncryptedString  `json:"public_key_full" gorm:"column:public_key_full"`
	IsActive                           bool                    `json:"is_active" gorm:"column:is_active"`
	ExpiresAt                          int64                   `json:"expires_at" gorm:"column:expires_at"`
}

func (*Entity) TableName() string { return "payments" }
