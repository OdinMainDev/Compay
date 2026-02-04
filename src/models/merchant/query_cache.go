package merchant

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
	if helper.CreatorId != 0 {
		if entity.CreatorId != helper.CreatorId {
			return false
		}
	}
	if helper.Token != "" {
		if entity.Token != helper.Token {
			return false
		}
	}
	if helper.Search != "" {
		if !strings.Contains(utils.ToLower(entity.Name), helper.Search) {
			return false
		}
	}
	return true
}
