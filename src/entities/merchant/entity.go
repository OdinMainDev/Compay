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
	PublicSettings PublicSettings `json:"public_settings"`
}
type PublicSettings struct {
	NotificationTelegramChatId int64 `json:"notification_telegram_chat_id"`
}

func (*Entity) TableName() string                  { return "merchants" }
func (data *Settings) Scan(value any) error        { return entity.ScanStruct(data, value) }
func (data Settings) Value() (driver.Value, error) { return entity.ValueStruct(data) }
