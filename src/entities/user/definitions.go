package user

import (
	"efw/components/entity"
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
