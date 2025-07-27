package merchant

import (
	"compay/application/entities/datatable"
	"compay/application/entities/merchant"
	"compay/application/helpers"
	"compay/application/http/requests"
	"compay/application/service/esvalidator"
	"compay/common/cperrors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestMerchantProcess struct {
	Id              int64    `json:"id" binding:"min=0"`
	Name            string   `json:"name" binding:"required,max=255"`
	Token           string   `json:"token" binding:"required,max=255"`
	IsEnabled       bool     `json:"is_enabled"`
	ProfitPercent   float64  `json:"profit_percent" binding:"required,min=0,max=100"`
	AvailCurrencies []string `json:"avail_currencies"`
}

var merchantsColumnsOrder = []string{"id", "name", "token", "is_enabled", "created_at"}

func MerchantsList(context *gin.Context) {
	requests.DatatableResponse[merchant.Model](context, merchantsColumnsOrder, func(_ int64) int64 {
		return merchant.GetCount()
	}, nil)
}

func MerchantsListFilters(context *gin.Context) {
	helpers.GinSuccessResponse(context, gin.H{"data": []datatable.ResponseFilter{
		{
			Name: "id",
			Type: datatable.FilterTypeNumeric,
		},
		{
			Name: "name",
			Type: datatable.FilterTypeText,
		},
		{
			Name: "token",
			Type: datatable.FilterTypeText,
		},
		{
			Name: "is_enabled",
			Type: datatable.FilterTypeBool,
		},
		{
			Name: "created_at",
			Type: datatable.FilterTypeDate,
		},
	}})
}

func MerchantGet(context *gin.Context) {
	requests.GetRecord(context, merchant.GetModelById, merchant.QueryHelper{})
}

func MerchantProcess(context *gin.Context) {
	var input requestMerchantProcess

	if err := context.ShouldBindJSON(&input); err != nil {
		helpers.GinErrorResponse(context, esvalidator.ValidatorResponse(err))
		return
	}

	merchantModel := &merchant.Model{}
	if input.Id != 0 {
		merchantModel = merchant.GetModel(merchant.QueryHelper{
			Id: input.Id,
		})
		if merchantModel == nil || merchantModel.GetId() == 0 {
			helpers.GinErrorResponse(context, nil, http.StatusNotFound)
			return
		}
	}

	check := merchant.GetModel(merchant.QueryHelper{
		Token: input.Token,
	})
	if check != nil && check.GetId() != 0 && (check.GetId() != input.Id || input.Id == 0) {
		helpers.GinErrorResponse(context, gin.H{"token": cperrors.ErrorAdminMerchantTokenMerchantDuplicate.Error()})
		return
	}

	merchantModel.Name = input.Name
	merchantModel.Token = input.Token
	merchantModel.IsEnabled = input.IsEnabled
	merchantModel.Settings.ProfitPercent = input.ProfitPercent
	merchantModel.Settings.AvailCurrencies = input.AvailCurrencies
	merchantModel.Save(merchantModel)

	helpers.GinSuccessResponse(context, gin.H{"data": merchantModel})
}

func GetMerchants(context *gin.Context) {
	helpers.GinSuccessResponse(context, gin.H{"data": merchant.GetModels(merchant.QueryHelper{})})
}
