package merchant

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/merchantcurrency"
	"compay/src/models/user"
	"efw/core/efwamount"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/services/efwcontroller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requetsAdminMerchantCurrencyGet struct {
	MerchantId types.MerchantId `uri:"merchantId" binding:"required,min=0"`
}

type requetsAdminMerchantCurrencyProcess struct {
	Id                types.MerchantCurrencyId `json:"id" binding:"required,gt=0"`
	MerchantId        types.MerchantId         `json:"merchant_id" binding:"required,gt=0"`
	MinAmount         float64                  `json:"min_amount" binding:"omitempty,min=0"`
	MaxAmount         float64                  `json:"max_amount" binding:"required,gt=0"`
	Confirms          types.Confirmation       `json:"confirms" binding:"required,gt=0"`
	InvoiceTime       int64                    `json:"invoice_time" binding:"required,gt=0"`
	MinAmountWithdraw float64                  `json:"min_amount_withdraw" binding:"required,gt=0"`
	MaxAmountWithdraw float64                  `json:"max_amount_withdraw" binding:"required,gt=0"`
	FeeType           types.FeeType            `json:"fee_type" binding:"required,feeTypeExist"`
	FeeFixed          float64                  `json:"fee_fixed" binding:"required,min=0"`
	FeePercent        float64                  `json:"fee_percent" binding:"required,min=0,max=100"`
}

func AdminMerchantCurrenciesGet(context *gin.Context) {
	var input requetsAdminMerchantCurrencyGet
	if !efwgin.ShouldBindUri(context, &input) {
		return
	}

	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	efwcontroller.GetModel(
		context,
		merchantcurrency.GetModel,
		merchantcurrency.DTO,
		func(id int64) merchantcurrency.QueryHelper {
			qH := merchantcurrency.QueryHelper{Id: types.MerchantCurrencyId(id), MerchantId: input.MerchantId}
			if !isUserRootAdmin(adminUserModel.GetId()) {
				qH.CreatorId = adminUserModel.GetId()
			}
			return qH
		},
	)
}

func AdminMerchantCurrenciesGetAll(context *gin.Context) {
	var input efwcontroller.RequestWithId
	if !efwgin.ShouldBindUri(context, &input) {
		return
	}

	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	qH := merchantcurrency.QueryHelper{MerchantId: types.MerchantId(input.Id)}
	if !isUserRootAdmin(adminUserModel.GetId()) {
		qH.CreatorId = adminUserModel.GetId()
	}

	merchantCurrencyModels := merchantcurrency.GetModels(qH)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": merchantCurrencyModels})
}

func AdminMerchantCurrenciesProcess(context *gin.Context) {
	var input requetsAdminMerchantCurrencyProcess
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	qH := merchantcurrency.QueryHelper{Id: input.Id, MerchantId: input.MerchantId}
	if !isUserRootAdmin(adminUserModel.GetId()) {
		qH.CreatorId = adminUserModel.GetId()
	}

	merchantCurrencyModel := merchantcurrency.GetModel(qH)
	if merchantCurrencyModel == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	merchantCurrencyModel.MinAmount = efwamount.FloatToAmountByCodeAlpha(input.MinAmount, merchantCurrencyModel.Currency)
	merchantCurrencyModel.MaxAmount = efwamount.FloatToAmountByCodeAlpha(input.MaxAmount, merchantCurrencyModel.Currency)
	merchantCurrencyModel.MinAmountWithdraw = efwamount.FloatToAmountByCodeAlpha(input.MinAmountWithdraw, merchantCurrencyModel.Currency)
	merchantCurrencyModel.MaxAmountWithdraw = efwamount.FloatToAmountByCodeAlpha(input.MaxAmountWithdraw, merchantCurrencyModel.Currency)
	merchantCurrencyModel.Confirms = input.Confirms
	merchantCurrencyModel.InvoiceTime = input.InvoiceTime
	merchantCurrencyModel.FeeType = input.FeeType
	merchantCurrencyModel.FeeFixed = types.FeeAmount(efwamount.FloatToAmountByCodeAlpha(input.FeeFixed, merchantCurrencyModel.Currency))
	merchantCurrencyModel.FeePercent = types.FeePercentAmount(efwamount.FloatToAmount(input.FeePercent, 2))
	merchantcurrency.Save(merchantCurrencyModel, merchantcurrency.QueryHelper{})

	createLogUser(context, adminUserModel, constants.LogUserAdminMerchantCurrencyProcess)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": merchantcurrency.DTO(merchantCurrencyModel)})
}
