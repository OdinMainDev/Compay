package merchant

import (
	"compay/application/entities/currency"
	"compay/application/entities/datatable"
	"compay/application/helpers"
	"compay/application/http/requests"
	"compay/application/service/esvalidator"
	"compay/application/service/start"
	"compay/common/cperrors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestCurrencyProcess struct {
	Id         int64  `json:"id" binding:"min=0"`
	Title      string `json:"title" binding:"required,max=255"`
	Image      string `json:"image" binding:"required,max=255"`
	Network    string `json:"network" binding:"required,max=255"`
	IsCoin     bool   `json:"is_coin"`
	CodeAlpha  string `json:"code_alpha" binding:"required,max=255"`
	MinorUnits int    `json:"minor_units" binding:"required,min=0"`
}

var currenciesColumnsOrder = []string{"id", "title", "code_alpha", "minor_units", "created_at"}
var currenciesImportStarted bool = false

func CurrenciesList(context *gin.Context) {
	requests.DatatableResponse[currency.Model](context, currenciesColumnsOrder, func(_ int64) int64 {
		return currency.GetCount()
	}, nil)
}

func CurrenciesListFilters(context *gin.Context) {
	helpers.GinSuccessResponse(context, gin.H{"data": []datatable.ResponseFilter{
		{
			Name: "id",
			Type: datatable.FilterTypeNumeric,
		},
		{
			Name: "title",
			Type: datatable.FilterTypeText,
		},
		{
			Name: "code_alpha",
			Type: datatable.FilterTypeText,
		},
		{
			Name: "minor_units",
			Type: datatable.FilterTypeNumeric,
		},
		{
			Name: "created_at",
			Type: datatable.FilterTypeDate,
		},
	}})
}

func CurrencyGet(context *gin.Context) {
	requests.GetRecord(context, currency.GetModelById, currency.QueryHelper{})
}

func CurrencyProcess(context *gin.Context) {
	var input requestCurrencyProcess

	if err := context.ShouldBindJSON(&input); err != nil {
		helpers.GinErrorResponse(context, esvalidator.ValidatorResponse(err))
		return
	}

	currencyModel := &currency.Model{}
	if input.Id != 0 {
		currencyModel = currency.GetModel(currency.QueryHelper{
			Id: input.Id,
		})
		if currencyModel == nil || currencyModel.GetId() == 0 {
			helpers.GinErrorResponse(context, nil, http.StatusNotFound)
			return
		}
	}

	check := currency.GetModel(currency.QueryHelper{
		CodeAlpha: input.CodeAlpha,
	})
	if check != nil && check.GetId() != 0 && (check.GetId() != input.Id || input.Id == 0) {
		helpers.GinErrorResponse(context, gin.H{"code_alpha": cperrors.ErrorAdminCurrencyCodeAlphaDuplicate.Error()})
		return
	}

	currencyModel.Title = input.Title
	currencyModel.Image = input.Image
	currencyModel.Network = input.Network
	currencyModel.IsCoin = input.IsCoin
	currencyModel.CodeAlpha = input.CodeAlpha
	currencyModel.MinorUnits = input.MinorUnits
	currencyModel.Save(currencyModel)

	helpers.GinSuccessResponse(context, gin.H{"data": currencyModel})
}

func CurrencyDelete(context *gin.Context) {
	requests.DeleteRecord(context, currency.GetModelById, currency.QueryHelper{})
}

func CurrencyImport(context *gin.Context) {
	if currenciesImportStarted {
		helpers.GinErrorResponse(context, cperrors.ErrorAdminCurrenciesImportAlreadyStarted)
		return
	}
	currenciesImportStarted = true
	go func() {
		start.FillCurrencies()
		currenciesImportStarted = false
	}()
	helpers.GinSuccessResponse(context, nil)
}

func GetCurrencies(context *gin.Context) {
	helpers.GinSuccessResponse(context, gin.H{"data": currency.GetModels(currency.QueryHelper{})})
}
