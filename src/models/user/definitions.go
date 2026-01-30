package user

import (
	"efw/components/entity"
	"efw/components/model"
)

func (*Model) GetSaveType() entity.SaveType { return entity.SAVE_TYPE_DB }

var (
	GetModel  = model.GetModel[*Model, QueryHelper]
	GetModels = model.GetModels[*Model, QueryHelper]
	Save      = model.Save[*Model, QueryHelper]
	GetCount  = model.GetCount[*Model, QueryHelper]
)
