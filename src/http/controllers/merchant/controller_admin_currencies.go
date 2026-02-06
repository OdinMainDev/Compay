package merchant

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/common/types"
	"compay/src/models/currency"
	"compay/src/models/user"
	"compay/src/services/importdata"
	"efw/components/entity"
	"efw/components/model"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/services/efwcontroller"
	"efw/services/efwdt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestAdminCurrenciesList struct {
	Id         types.CurrencyId          `json:"id" efwdt:"id" binding:"omitempty,gt=0"`
	CodeAlpha  entity.CurrencyCodeAlpha  `json:"code_alpha" efwdt:"code_alpha" binding:"omitempty,gt=0"`
	Title      string                    `json:"title" efwdt:"title" binding:"omitempty,gt=0"`
	Network    types.CurrencyNetwork     `json:"network" efwdt:"network" binding:"omitempty,currencyNetworkExist"`
	IsCoin     *types.CurrencyIsCoin     `json:"is_coin" efwdt:"is_coin" binding:"omitempty"`
	MinorUnits entity.CurrencyMinorUnits `json:"minor_units" efwdt:"minor_units" binding:"omitempty"`
	efwdt.WithCreatedAt
}

type requestAdminCurrenciesProcess struct {
	Id    types.CurrencyId `json:"id" binding:"required,gt=0"`
	Title string           `json:"title" efwdt:"title" binding:"required,min=1"`
}

type requestAdminCurrenciesAction struct {
	Id types.CurrencyId `json:"id" binding:"required,gt=0"`
}

func AdminCurrenciesList(context *gin.Context) {
	efwcontroller.Datable[currency.Model, requestAdminCurrenciesList](
		context,
		[]string{"id", "code_alpha", "title", "network", "is_coin", "minor_units", "created_at"},
		nil,
		efwdt.SpecialFilterMap{
			"code_alpha": func(fieldName string, value any) (efwdt.SpecialFilters, bool) {
				v, ok := value.(entity.CurrencyCodeAlpha)
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
		nil,
	)
}

func AdminCurrenciesListFilters(context *gin.Context) {
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

func AdminCurrenciesGet(context *gin.Context) {
	efwcontroller.GetModel(
		context,
		currency.GetModel,
		nil,
		func(id int64) currency.QueryHelper {
			return currency.QueryHelper{Id: types.CurrencyId(id)}
		},
	)
}

func AdminCurrenciesProcess(context *gin.Context) {
	var input requestAdminCurrenciesProcess
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	currencyModel := currency.GetModel(currency.QueryHelper{
		Id: input.Id,
	})
	if currencyModel == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	currencyModel.Title = input.Title
	currency.Save(currencyModel, currency.QueryHelper{})

	createLogUser(context, adminUserModel, constants.LogUserAdminCurrencyProcess)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": currencyModel})
}

func AdminCurrenciesSearch(context *gin.Context) {
	var input requestSearch
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": currency.GetModels(currency.QueryHelper{
		Search: input.Search,
		QueryHelper: model.QueryHelper{
			HasPagination: model.HasPagination{
				HasLimit: true,
			},
		},
	})})
}

func AdminCurrenciesActionImport(context *gin.Context) {
	isWorking := importdata.ImportCurrencies()
	if isWorking {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorCurrencyImportIsAlreadyStarted))
		return
	}

	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	createLogUser(context, adminUserModel, constants.LogUserAdminCurrencyImport)
	efwginhelpers.ResponseSuccess(context, nil)
}

func AdminCurrenciesActionDisable(context *gin.Context) {
	adminCurrenciesActionDisableOrNot(context, constants.LogUserAdminCurrencyDisable)
}

func AdminCurrenciesActionEnable(context *gin.Context) {
	adminCurrenciesActionDisableOrNot(context, constants.LogUserAdminCurrencyEnable)
}

func adminCurrenciesActionDisableOrNot(context *gin.Context, action types.LogUserAction) {
	var input requestAdminCurrenciesAction
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	currencyModel := currency.GetModel(currency.QueryHelper{Id: input.Id})
	if currencyModel == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	actionSuccess := false

	if action == constants.LogUserAdminCurrencyDisable && currencyModel.IsEnabled {
		currencyModel.IsEnabled = false
		actionSuccess = true
	}

	if action == constants.LogUserAdminCurrencyEnable && !currencyModel.IsEnabled {
		currencyModel.IsEnabled = true
		actionSuccess = true
	}

	if actionSuccess {
		currency.Save(currencyModel, currency.QueryHelper{})
		createLogUser(context, adminUserModel, action)
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": currencyModel})
}
