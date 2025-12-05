package user

import (
	"efw/components/entity"
	"efw/components/model"
)

func (Entity) TableName() string {
	return "users"
}

func (Entity) GetSaveType() entity.SaveType {
	return entity.SAVE_TYPE_DB
}

func (Entity) CacheArgs() []int64 {
	return nil
}

var (
	GetModel = model.GetModel[Entity, QueryHelper]
)
