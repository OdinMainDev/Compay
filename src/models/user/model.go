package user

import (
	entity "compay/src/entities/user"
	"efw/components/model"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
}

var (
	GetModel = model.GetModel[*Model, QueryHelper]
	Save     = model.Save[*Model]
	Delete   = model.Delete[*Model]
)
