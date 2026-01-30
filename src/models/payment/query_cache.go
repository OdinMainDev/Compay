package payment

import "gorm.io/gorm"

func (helper QueryHelper) CacheAdditionalKeys(model *Model) []int64 { return nil }
func (helper QueryHelper) ApplyCacheFilters(entity *Model) bool     { return true }
func (helper QueryHelper) ApplyCacheDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	return nil
}
