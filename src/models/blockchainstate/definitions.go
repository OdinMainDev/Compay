package blockchainstate

import (
	"efw/components/entity"
	"efw/components/model"
)

func (*Model) GetSaveType() entity.SaveType { return entity.SAVE_TYPE_DB_CACHE }

var (
	GetModel  = model.GetModel[*Model, QueryHelper]
	GetModels = model.GetModels[*Model, QueryHelper]
	Save      = model.Save[*Model, QueryHelper]
	Delete    = model.Delete[*Model, QueryHelper]
)
