package user

import (
	"compay/common/types"

	"gorm.io/gorm"
)

type QueryHelper struct {
	Id                  types.UserId
	DontCheckMerchantId bool
	MerchantId          types.MerchantId
	Username            string
	WithMerchant        bool
}

func (helper QueryHelper) CacheAdditionalKeys() []int64 {
	return nil
}

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	if !helper.DontCheckMerchantId {
		chain = chain.Where("merchant_id = ?", helper.MerchantId)
	}
	if helper.Username != "" {
		chain = chain.Where("username = ?", helper.Username)
	}
	if helper.WithMerchant {
		chain = chain.Preload("Merchant", func(db gorm.PreloadBuilder) error { return nil })
	}
	return chain.Order("id DESC")
}

func (helper QueryHelper) ApplyCacheFilters(entity *Model) bool {
	return true
}
