package logauth

import (
	"compay/common/types"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[types.LogAuthId] `gorm:"embedded"`
	UserId                             entity.UserId       `json:"user_id" gorm:"column:user_id"`
	MerchantId                         types.MerchantId    `json:"merchant_id" gorm:"column:merchant_id"`
	Action                             types.LogAuthAction `json:"action" gorm:"column:action"`
	Url                                entity.Url          `json:"url" gorm:"column:url"`
	FingerPrint                        entity.FingerPrint  `json:"fingerprint" gorm:"column:fingerprint"`
	Ip                                 entity.Ip           `json:"ip" gorm:"column:ip"`
	UserAgent                          entity.UserAgent    `json:"user_agent" gorm:"column:user_agent"`
	ClientData                         entity.ClientData   `json:"client_data" gorm:"column:client_data"`
}

func (*Entity) TableName() string { return "log_auth" }
