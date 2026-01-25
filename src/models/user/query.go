package user

import (
	"compay/common/types"
	"efw/components/entity"
	"efw/components/model"
	"efw/utils"

	"gorm.io/gorm"
)

type QueryHelper struct {
	Id                  entity.UserId
	DontCheckMerchantId bool
	MerchantId          types.MerchantId
	Username            string
	WithMerchant        bool
	NotBanned           bool
	Search              string
	model.HasPagination
}

func (helper QueryHelper) CacheAdditionalKeys() []int64         { return nil }
func (helper QueryHelper) ApplyCacheFilters(entity *Model) bool { return true }
func (helper QueryHelper) GetPagination() *model.HasPagination  { return &helper.HasPagination }

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	helper.Search = utils.ToLower(helper.Search)

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
	if helper.NotBanned {
		chain = chain.Where("is_banned = ?", false)
	}
	if helper.Search != "" {
		chain = chain.Where("username ILIKE ?", "%"+helper.Search+"%")
	}
	return chain.Order("id DESC")
}
