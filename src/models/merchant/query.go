package merchant

import (
	"compay/common/types"
	"efw/components/model"
	"efw/utils"
	"strings"

	"gorm.io/gorm"
)

type QueryHelper struct {
	model.QueryHelper
	Id     types.MerchantId
	Search string
}

func (helper QueryHelper) CacheAdditionalKeys() []int64 { return nil }

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	helper.Search = utils.ToLower(helper.Search)

	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}

	if helper.Search != "" {
		chain = chain.Where("name ILIKE ?", "%"+helper.Search+"%")
	}
	return chain.Order("id DESC")
}

func (helper QueryHelper) ApplyCacheFilters(entity *Model) bool {
	helper.Search = utils.ToLower(helper.Search)

	if helper.Id != 0 {
		if entity.GetId() != helper.Id {
			return false
		}
	}

	if helper.Search != "" {
		if !strings.Contains(entity.Name, helper.Search) {
			return false
		}
	}
	return true
}
