package logauth

import (
	"compay/common/types"
	"efw/components/model"

	"gorm.io/gorm"
)

type QueryHelper struct {
	Id types.LogAuthId
}

func (helper QueryHelper) CacheAdditionalKeys() []int64         { return nil }
func (helper QueryHelper) ApplyCacheFilters(entity *Model) bool { return true }
func (helper QueryHelper) GetPagination() *model.HasPagination  { return nil }

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	return chain.Order("id DESC")
}
