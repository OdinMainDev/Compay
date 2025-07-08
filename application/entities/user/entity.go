package user

import "compay/application/entities/basemodel"

type Model struct {
	basemodel.BaseModel
	Token       basemodel.EncryptedString `json:"token" gorm:"column:token"`
	PinCode     basemodel.EncryptedString `json:"pin_code" gorm:"column:pin_code"`
	IsBanned    bool                      `json:"is_banned" gorm:"column:is_banned"`
	SecretToken string                    `json:"secret_token" gorm:"column:secret_token"`
}

func (Model) TableName() string {
	return "users"
}
