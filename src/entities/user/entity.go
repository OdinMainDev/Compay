package user

import (
	"compay/common/types"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[entity.UserId] `gorm:"embedded"`
	MerchantId                       types.MerchantId       `json:"merchant_id" gorm:"column:merchant_id"`
	CreatorId                        entity.UserId          `json:"creator_id" gorm:"column:creator_id"`
	Username                         string                 `json:"username" gorm:"column:username"`
	Password                         string                 `json:"-" gorm:"column:password"`
	SecretToken                      string                 `json:"-" gorm:"column:secret_token"`
	TotpToken                        entity.EncryptedString `json:"-" gorm:"column:totp_token"`
	IsTwofaEnabled                   bool                   `json:"is_twofa_enabled" gorm:"column:is_twofa_enabled"`
	IsBanned                         bool                   `json:"is_banned" gorm:"column:is_banned"`
}

func (*Entity) TableName() string {
	return "users"
}
