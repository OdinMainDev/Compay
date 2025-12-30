package merchant

import (
	"compay/common/types"
	"strings"

	"gorm.io/gorm"
)

type QueryHelper struct {
	Id     types.UserId
	Search string
}

func (helper QueryHelper) CacheAdditionalKeys() []int64 {
	return nil
}

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}

	if helper.Search != "" {
		helper.Search = strings.ToLower(helper.Search)
		chain = chain.Where("name ILIKE ?", "%"+helper.Search+"%")
	}
	return chain.Order("id DESC")
}

func (helper QueryHelper) ApplyCacheFilters(entity *Model) bool {
	return true
}
