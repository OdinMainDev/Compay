package user

import (
	"compay/src/entities/merchant"
	entity "compay/src/entities/user"
	"efw/components/model"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
	Merchant      *merchant.Entity `json:"merchant,omitempty" gorm:"foreignKey:MerchantId;references:Id"`
}

var (
	GetModel = model.GetModel[*Model, QueryHelper]
	Save     = model.Save[*Model]
	Delete   = model.Delete[*Model]
)
