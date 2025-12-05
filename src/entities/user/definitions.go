package user

import (
	"efw/components/entity"
	"efw/components/model"
)

func (*Entity) TableName() string {
	return "users"
}

func (*Entity) GetSaveType() entity.SaveType {
	return entity.SAVE_TYPE_DB
}

func (*Entity) CacheArgs() []int64 {
	return nil
}

func (e *Entity) GetId() int64 {
	return e.GetId()
}

func (e *Entity) GetCreatedAt() int64 {
	return e.GetCreatedAt()
}

func (e *Entity) GetUpdatedAt() int64 {
	return e.GetUpdatedAt()
}

func (e *Entity) Lock() {
	e.Lock()
}

func (e *Entity) Unlock() {
	e.Unlock()
}

func (e *Entity) UpdateTimestamps() {
	e.TouchTimestamps()
}

var (
	GetModel = model.GetModel[Entity, QueryHelper]
)
