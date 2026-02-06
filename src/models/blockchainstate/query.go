package blockchainstate

import (
	"compay/common/types"
	"efw/components/entity"
	"efw/components/model"

	"gorm.io/gorm"
)

type QueryHelper struct {
	model.QueryHelper
	Id       types.BlockchainStateId
	Currency entity.CurrencyCodeAlpha
}

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	if helper.Currency != "" {
		chain = chain.Where("currency = ?", helper.Currency)
	}
	return chain.Order("id DESC")
}
