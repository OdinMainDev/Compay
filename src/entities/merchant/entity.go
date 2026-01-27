package merchant

import (
	"compay/common/types"
	"database/sql/driver"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[types.MerchantId] `gorm:"embedded"`
	Name                                string                 `json:"name" gorm:"column:name"`
	Token                               entity.EncryptedString `json:"-" gorm:"column:token"`
	CreatorId                           entity.UserId          `json:"creator_id" gorm:"column:creator_id"`
	OwnerId                             entity.UserId          `json:"owner_id" gorm:"column:owner_id"`
	IsEnabled                           bool                   `json:"is_enabled" gorm:"column:is_enabled"`
	Settings                            Settings               `json:"settings" gorm:"column:settings"`
}

type Settings struct {
	Conditions     Conditions      `json:"conditions"`
	MinimumAmounts []MinimumAmount `json:"minimum_amounts"`
	FeeAmounts     []FeeAmount     `json:"fee_amounts"`
	PublicSettings PublicSettings  `json:"public_settings"`
}

type Conditions struct {
	BlockedCountries []string `json:"blocked_countries"`
	AllowedCountries []string `json:"allowed_countries"`
}

type MinimumAmount struct {
	CurrencyName string       `json:"currency_name"`
	Amount       types.Amount `json:"amount"`
}

type FeeAmount struct {
	CurrencyName string                 `json:"currency_name"`
	Amount       types.FeePercentAmount `json:"amount"`
}

type PublicSettings struct {
	NotificationTelegramChatId int64 `json:"notification_telegram_chat_id"`
}

func (*Entity) TableName() string                  { return "merchants" }
func (data *Settings) Scan(value any) error        { return entity.ScanStruct(data, value) }
func (data Settings) Value() (driver.Value, error) { return entity.ValueStruct(data) }
