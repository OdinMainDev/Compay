package notification

import (
	"compay/common/types"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[types.NotificationId] `gorm:"embedded"`
	MerchantId                              types.MerchantId         `json:"-" gorm:"column:merchant_id"`
	IsRead                                  bool                     `json:"is_read" gorm:"column:is_read"`
	Action                                  types.NotificationAction `json:"action" gorm:"column:action"`
	Text                                    types.NotificationText   `json:"text" gorm:"column:text"`
}

func (*Entity) TableName() string { return "notifications" }
