package merchant

import (
	entity "compay/src/entities/merchant"
	"efw/components/model"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
	// UserAdmin     *user.Entity `json:"user_admin,omitempty" gorm:"foreignKey:AdminId;references:Id"`
}

var (
	GetModel  = model.GetModel[*Model, QueryHelper]
	GetModels = model.GetModels[*Model, QueryHelper]
	Save      = model.Save[*Model]
	Delete    = model.Delete[*Model]
)
