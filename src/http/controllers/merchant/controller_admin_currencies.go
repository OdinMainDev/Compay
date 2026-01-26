package merchant

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/currency"
	"efw/services/efwcontroller"
	"efw/services/efwdt"

	"github.com/gin-gonic/gin"
)

type requestAdminCurrencyList struct {
	Id         types.CurrencyId         `json:"id" efwdt:"id" binding:"omitempty,gt=0"`
	CodeAlpha  types.CurrencyCodeAlpha  `json:"code_alpha" efwdt:"code_alpha" binding:"omitempty,gt=0"`
	Title      string                   `json:"title" efwdt:"title" binding:"omitempty,gt=0"`
	Network    types.CurrencyNetwork    `json:"network" efwdt:"network" binding:"omitempty,currencyNetworkExist"`
	IsCoin     *types.CurrencyIsCoin    `json:"is_coin" efwdt:"is_coin" binding:"omitempty"`
	MinorUnits types.CurrencyMinorUnits `json:"minor_units" efwdt:"minor_units" binding:"omitempty"`
	efwdt.WithCreatedAt
}

func AdminCurrencyList(context *gin.Context) {
	efwcontroller.Datable[currency.Model, requestAdminCurrencyList](
		context,
		[]string{"id", "code_alpha", "title", "network", "is_coin", "minor_units", "created_at"},
		nil,
		efwdt.SpecialFilterMap{
			"code_alpha": func(fieldName string, value any) (efwdt.SpecialFilters, bool) {
				v, ok := value.(types.CurrencyCodeAlpha)
				if !ok {
					return efwdt.SpecialFilters{}, false
				}
				return efwdt.SpecialFilters{
					Query: fieldName + " ILIKE ?",
					Args:  []any{"%" + v + "%"},
				}, true
			},
			"title": func(fieldName string, value any) (efwdt.SpecialFilters, bool) {
				v, ok := value.(string)
				if !ok {
					return efwdt.SpecialFilters{}, false
				}
				return efwdt.SpecialFilters{
					Query: fieldName + " ILIKE ?",
					Args:  []any{"%" + v + "%"},
				}, true
			},
		},
	)
}

func AdminCurrencyListFilters(context *gin.Context) {
	efwcontroller.DatatableFilters(context, []efwdt.Filter{
		{Name: "id", Type: types.EFWDT_FilterTypeInteger},
		{Name: "code_alpha", Type: types.EFWDT_FilterTypeText},
		{Name: "title", Type: types.EFWDT_FilterTypeText},
		{Name: "network", Type: types.EFWDT_FilterTypeList, Values: constants.CurrencyNetworks},
		{Name: "is_coin", Type: types.EFWDT_FilterTypeBoolean},
		{Name: "minor_units", Type: types.EFWDT_FilterTypeInteger},
		{Name: "created_at", Type: types.EFWDT_FilterTypeDate},
	})
}

func AdminCurrencyGet(context *gin.Context) {
	efwcontroller.GetModel(
		context,
		currency.GetModel,
		nil,
		func(id int64) currency.QueryHelper {
			return currency.QueryHelper{Id: types.CurrencyId(id)}
		},
	)
}

func AdminCurrencyProcess(context *gin.Context) {
	efwcontroller.GetModel(
		context,
		currency.GetModel,
		nil,
		func(id int64) currency.QueryHelper {
			return currency.QueryHelper{Id: types.CurrencyId(id)}
		},
	)
}
