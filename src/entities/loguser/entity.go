package loguser

import (
	"compay/common/types"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[types.LogUserId] `gorm:"embedded"`
	UserId                             entity.UserId       `json:"user_id" gorm:"column:user_id"`
	MerchantId                         types.MerchantId    `json:"merchant_id" gorm:"column:merchant_id"`
	Action                             types.LogUserAction `json:"action" gorm:"column:action"`
	Url                                entity.Url          `json:"url" gorm:"column:url"`
	Request                            string              `json:"request" gorm:"column:request"`
	ClientData                         entity.ClientData   `json:"client_data" gorm:"column:client_data"`
}

func (*Entity) TableName() string { return "log_user" }
