package merchant

import (
	"compay/common/types"
	"database/sql/driver"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[types.MerchantId] `gorm:"embedded"`
	Name                                string                 `json:"name" gorm:"column:name"`
	Token                               string                 `json:"-" gorm:"column:token"`
	SecretKey                           entity.EncryptedString `json:"-" gorm:"column:secret_key"`
	CreatorId                           entity.UserId          `json:"creator_id" gorm:"column:creator_id"`
	OwnerId                             entity.UserId          `json:"owner_id" gorm:"column:owner_id"`
	IsEnabled                           bool                   `json:"is_enabled" gorm:"column:is_enabled"`
	Settings                            Settings               `json:"settings" gorm:"column:settings"`
}

type Settings struct {
	Conditions           Conditions             `json:"conditions"`
	MinimumAmounts       []MinimumAmount        `json:"minimum_amounts"`
	MinimumConfirmations []MinimumConfirmations `json:"minimum_confirmations"`
	FeeAmounts           []FeeAmount            `json:"fee_amounts"`
	FeePercentAmounts    []FeePercentAmount     `json:"fee_percent_amounts"`
	PublicSettings       PublicSettings         `json:"public_settings"`
}

type Conditions struct {
	BlockedCountries []types.Country `json:"blocked_countries"`
	AllowedCountries []types.Country `json:"allowed_countries"`
}

type MinimumAmount struct {
	Currency types.CurrencyCodeAlpha `json:"currency" binding:"required,currencyExist"`
	Amount   types.Amount            `json:"amount" binding:"required,gt=0"`
}

type MinimumConfirmations struct {
	Currency types.CurrencyCodeAlpha `json:"currency" binding:"required,currencyExist"`
	Count    types.Confirmation      `json:"count" binding:"required,gt=0"`
}

type FeeAmount struct {
	Currency types.CurrencyCodeAlpha `json:"currency" binding:"required,currencyExist"`
	FeeType  types.FeeType           `json:"fee_type" binding:"required,feeTypeExist"`
	Amount   types.FeeAmount         `json:"amount"`
}

type FeePercentAmount struct {
	Currency types.CurrencyCodeAlpha `json:"currency" binding:"required,currencyExist"`
	FeeType  types.FeeType           `json:"fee_type" binding:"required,feeTypeExist"`
	Amount   types.FeePercentAmount  `json:"amount"`
}

type PublicSettings struct {
	NotificationTelegramChatId int64 `json:"notification_telegram_chat_id"`
}

func (*Entity) TableName() string                  { return "merchants" }
func (data *Settings) Scan(value any) error        { return entity.ScanStruct(data, value) }
func (data Settings) Value() (driver.Value, error) { return entity.ValueStruct(data) }
