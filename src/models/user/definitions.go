package user

import (
	"efw/components/entity"
)

func (*Model) GetSaveType() entity.SaveType {
	return entity.SAVE_TYPE_DB
}

func (*Model) CacheArgs() []int64 {
	return nil
}
