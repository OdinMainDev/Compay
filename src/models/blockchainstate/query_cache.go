package blockchainstate

import (
	"gorm.io/gorm"
)

func (helper QueryHelper) CacheAdditionalKeys(model *Model) []int64 { return nil }
func (helper QueryHelper) ApplyCacheDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	return chain.Order("id DESC")
}

func (helper QueryHelper) ApplyCacheFilters(entity *Model) bool {
	if helper.Id != 0 {
		if entity.GetId() != helper.Id {
			return false
		}
	}
	if helper.Currency != "" {
		if entity.Currency != helper.Currency {
			return false
		}
	}
	return true
}
