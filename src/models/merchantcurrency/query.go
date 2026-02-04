package merchantcurrency

import (
	"compay/common/types"
	"efw/components/entity"
	"efw/components/model"

	"gorm.io/gorm"
)

type QueryHelper struct {
	model.QueryHelper
	Id         types.MerchantCurrencyId
	CreatorId  entity.UserId
	Currency   types.CurrencyCodeAlpha
	MerchantId types.MerchantId
}

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	chain = chain.Where("merchant_id = ?", helper.MerchantId)

	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	if helper.Currency != "" {
		chain = chain.Where("currency = ?", helper.Currency)
	}
	if helper.CreatorId != 0 {
		chain = chain.Where("creator_id = ?", helper.CreatorId)
	}
	return chain.Order("id DESC")
}
