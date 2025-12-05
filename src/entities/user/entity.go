package user

import (
	"compay/common/types"
	"compay/src/entities"
)

type Entity struct {
	*entities.BaseEntity[types.UserId] `gorm:"embedded"`
	MerchantId                         types.MerchantId `json:"-" gorm:"column:merchant_id"`
	UserID                             int64            `json:"user_id" gorm:"column:user_id"`
	SecretToken                        string           `json:"-" gorm:"column:secret_token"`
	TotpToken                          string           `json:"-" gorm:"column:totp_token"`
	IsTwofaEnabled                     bool             `json:"is_twofa_enabled" gorm:"column:is_twofa_enabled"`
}
