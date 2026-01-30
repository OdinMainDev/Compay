package merchant

import (
	"compay/common/types"
	"efw/components/entity"
	"efw/components/model"
	"efw/utils"

	"gorm.io/gorm"
)

type QueryHelper struct {
	model.QueryHelper
	Id        types.MerchantId
	CreatorId entity.UserId
	Token     string
	Search    string
}

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	helper.Search = utils.ToLower(helper.Search)

	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	if helper.CreatorId != 0 {
		chain = chain.Where("creator_id = ?", helper.CreatorId)
	}
	if helper.Token != "" {
		chain = chain.Where("token = ?", helper.Token)
	}
	if helper.Search != "" {
		chain = chain.Where("name ILIKE ?", "%"+helper.Search+"%")
	}
	return chain.Order("id DESC")
}
