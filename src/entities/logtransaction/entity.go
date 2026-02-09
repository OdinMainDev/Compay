package logtransaction

import (
	"compay/common/types"
	"database/sql/driver"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[types.LogPaymentId] `gorm:"embedded"`
	PaymentId                             types.PaymentId          `json:"payment_id" gorm:"column:payment_id"`
	PaymentTrxId                          string                   `json:"payment_trx_id" gorm:"column:payment_trx_id"`
	WithdrawId                            types.WithdrawId         `json:"withdraw_id" gorm:"column:withdraw_id"`
	WithdrawTrxId                         string                   `json:"withdraw_trx_id" gorm:"column:withdraw_trx_id"`
	MerchantWalletId                      types.MerchantWalletId   `json:"merchant_wallet_id" gorm:"column:merchant_wallet_id"`
	MerchantId                            types.MerchantId         `json:"merchant_id" gorm:"column:merchant_id"`
	Currency                              entity.CurrencyCodeAlpha `json:"currency" gorm:"column:currency"`
	IsComplete                            bool                     `json:"is_complete" gorm:"column:is_complete"`
	StartBlock                            int64                    `json:"start_block" gorm:"column:start_block"`
	TxId                                  types.PaymentTxId        `json:"tx_id" gorm:"column:tx_id"`
	TxHash                                types.PaymentTxHash      `json:"tx_hash" gorm:"column:tx_hash"`
	Address                               types.PaymentAddress     `json:"address" gorm:"column:address"`
	Data                                  Data                     `json:"data" gorm:"column:data"`
}

type Data struct {
	TxData         types.TxData         `json:"tx_data"`
	TxWithdrawData types.TxWithdrawData `json:"withdraw_data"`
	Conditions     Conditions           `json:"conditions"`
}

type Conditions struct {
	RequiredConfirmations types.Confirmation `json:"required_confirmations"`
}

func (*Entity) TableName() string              { return "log_transactions" }
func (data *Data) Scan(value any) error        { return entity.ScanStruct(data, value) }
func (data Data) Value() (driver.Value, error) { return entity.ValueStruct(data) }
