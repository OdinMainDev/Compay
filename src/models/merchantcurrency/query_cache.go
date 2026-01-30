package merchantcurrency

import (
	"gorm.io/gorm"
)

func (helper QueryHelper) CacheAdditionalKeys(model *Model) []int64 {
	if model != nil {
		return []int64{int64(model.MerchantId)}
	}
	if helper.MerchantId == 0 {
		panic("BAD CONFIGURATION CacheAdditionalKeys merchantcurrency")
	}

	return []int64{int64(helper.MerchantId)}
}

func (helper QueryHelper) ApplyCacheDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	if helper.MerchantId == 0 {
		panic("BAD CONFIGURATION ApplyCacheDBFilters merchantcurrency")
	}

	return chain.Where("merchant_id ? ", helper.MerchantId).Order("id DESC")
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
