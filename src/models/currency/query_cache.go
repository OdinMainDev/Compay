package currency

import (
	"efw/utils"
	"strings"

	"gorm.io/gorm"
)

func (helper QueryHelper) CacheAdditionalKeys(model *Model) []int64 { return nil }
func (helper QueryHelper) ApplyCacheDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	return chain.Order("id DESC")
}

func (helper QueryHelper) ApplyCacheFilters(entity *Model) bool {
	helper.Search = utils.ToLower(helper.Search)

	if helper.Id != 0 {
		if entity.GetId() != helper.Id {
			return false
		}
	}
	if helper.CodeAlpha != "" {
		if entity.CodeAlpha != helper.CodeAlpha {
			return false
		}
	}
	if helper.IsEnabled {
		if !entity.IsEnabled {
			return false
		}
	}
	if helper.IsDisabled {
		if entity.IsEnabled {
			return false
		}
	}
	if helper.Search != "" {
		if !strings.Contains(utils.ToLower(entity.Title), helper.Search) {
			if !strings.Contains(utils.ToLower(string(entity.CodeAlpha)), helper.Search) {
				return false
			}
		}
	}
	return true
}
