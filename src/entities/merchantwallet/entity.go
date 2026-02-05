package merchantwallet

import (
	"compay/common/types"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[types.MerchantWalletId] `gorm:"embedded"`
	MerchantId                                types.MerchantId        `json:"merchant_id" gorm:"column:merchant_id"`
	PaymentId                                 types.PaymentId         `json:"payment_id" gorm:"column:payment_id"`
	Currency                                  types.CurrencyCodeAlpha `json:"currency" gorm:"column:currency"`
	Address                                   types.PaymentAddress    `json:"address" gorm:"column:address"`
	Amount                                    types.Amount            `json:"amount" gorm:"column:amount"`
	ReceivedAmount                            types.Amount            `json:"received_amount" gorm:"column:received_amount"`
	SentAmount                                types.Amount            `json:"sent_amount" gorm:"column:sent_amount"`
	PrivateKey                                entity.EncryptedString  `json:"private_key" gorm:"column:private_key"`
	PrivateKeyWif                             entity.EncryptedString  `json:"private_key_wif" gorm:"column:private_key_wif"`
	PublicKey                                 entity.EncryptedString  `json:"public_key" gorm:"column:public_key"`
	PublicKeyFull                             entity.EncryptedString  `json:"public_key_full" gorm:"column:public_key_full"`
}

func (*Entity) TableName() string { return "merchant_wallets" }
