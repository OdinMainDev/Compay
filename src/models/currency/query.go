package currency

import (
	"compay/common/types"
	"efw/components/model"
	"efw/utils"

	"gorm.io/gorm"
)

type QueryHelper struct {
	model.QueryHelper
	Id         types.CurrencyId
	CodeAlpha  types.CurrencyCodeAlpha
	IsEnabled  bool
	IsDisabled bool
	Search     string
}

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	helper.Search = utils.ToLower(helper.Search)

	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	if helper.CodeAlpha != "" {
		chain = chain.Where("code_alpha = ?", helper.CodeAlpha)
	}
	if helper.IsEnabled {
		chain = chain.Where("is_enabled = ?", true)
	}
	if helper.IsDisabled {
		chain = chain.Where("is_enabled = ?", false)
	}
	if helper.Search != "" {
		chain = chain.Where(
			"(code_alpha ILIKE ? OR title ILIKE ?)",
			"%"+helper.Search+"%",
			"%"+helper.Search+"%",
		)
	}
	return chain.Order("id DESC")
}
