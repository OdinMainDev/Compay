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
	NotificationSettings NotificationSettings `json:"notification_settings"`
}
type NotificationSettings struct {
	TelegramChatId             int64 `json:"telegram_chat_id"`
	DisableCreateInvoice       bool  `json:"disable_create_invoice"`
	DisableDetectedTransaction bool  `json:"disable_detected_transaction"`
	DisableFundsRecieved       bool  `json:"disable_funds_recieved"`
	DisableCompletedInvoice    bool  `json:"disable_completed_invoice"`
	DisableExpireInvoice       bool  `json:"disable_expire_invoice"`
	DisableInsufficientFunds   bool  `json:"disable_insufficient_funds"`
}

func (*Entity) TableName() string                  { return "merchants" }
func (data *Settings) Scan(value any) error        { return entity.ScanStruct(data, value) }
func (data Settings) Value() (driver.Value, error) { return entity.ValueStruct(data) }
