package user

import (
	"compay/common/types"

	"gorm.io/gorm"
)

type QueryHelper struct {
	Id                  types.UserId
	DontCheckMerchantId bool
	MerchantId          types.MerchantId
}

func (helper QueryHelper) CacheAdditionalKeys() []int64 {
	return nil
}

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[Entity]) gorm.ChainInterface[Entity] {
	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	if !helper.DontCheckMerchantId {
		chain = chain.Where("merchant_id = ?", helper.MerchantId)
	}
	return chain.Order("id DESC")
}

func (helper QueryHelper) ApplyCacheFilters(entity *Entity) bool {
	if helper.Id != 0 && entity.ID != helper.Id {
		return false
	}
	if !helper.DontCheckMerchantId && helper.MerchantId != 0 && entity.MerchantId != helper.MerchantId {
		return false
	}
	return true
}
