package wallet

import "compay/application/entities/basemodel"

type Model struct {
	basemodel.BaseModel
	UserId     int64                     `json:"user_id" gorm:"column:user_id"`
	Mnemonic   basemodel.EncryptedString `json:"mnemonic" gorm:"column:mnemonic"`
	PrivateKey basemodel.EncryptedString `json:"private_key" gorm:"column:private_key"`
}

func (Model) TableName() string {
	return "wallets"
}
