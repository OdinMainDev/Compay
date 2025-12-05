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

func (helper QueryHelper) BucketArgs() []int64 {
	return nil
}

func (helper QueryHelper) ApplyBucketQuery(chain gorm.ChainInterface[*Entity]) gorm.ChainInterface[*Entity] {
	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	if !helper.DontCheckMerchantId {
		chain = chain.Where("merchant_id = ?", helper.MerchantId)
	}
	return chain.Order("id DESC")
}

func (helper QueryHelper) Matches(entity *Entity) bool {
	return false
}
