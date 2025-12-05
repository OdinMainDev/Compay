package user

import (
	"compay/common/types"
	"efw/components/entity"
)

type Entity struct {
	*entity.BaseEntity[types.UserId] `gorm:"embedded"`
	MerchantId                       types.MerchantId `json:"-" gorm:"column:merchant_id"`
	Username                         string           `json:"-" gorm:"column:username"`
	Password                         string           `json:"-" gorm:"column:password"`
	SecretToken                      string           `json:"-" gorm:"column:secret_token"`
	TotpToken                        string           `json:"-" gorm:"column:totp_token"`
	IsTwofaEnabled                   bool             `json:"is_twofa_enabled" gorm:"column:is_twofa_enabled"`
}
