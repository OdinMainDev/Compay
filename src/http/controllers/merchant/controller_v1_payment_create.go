package merchant

import (
	"compay/common/cperrors"
	"compay/src/models/currency"
	"compay/src/models/merchant"
	"compay/src/models/payment"
	"compay/src/services/processing"
	"efw/components/entity"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
)

type requestV1PaymentCreate struct {
	Amount   float64                  `json:"amount" binding:"omitempty,min=0"`
	Token    string                   `json:"token" binding:"required,len=36"`
	Currency entity.CurrencyCodeAlpha `json:"currency" binding:"required,max=20"`
	Sign     string                   `json:"sign" binding:"required,len=128"`
}

func V1PaymentCreate(context *gin.Context) {
	if !isMerchantStarted() {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorApiV1Maintenance))
		return
	}

	var input requestV1PaymentCreate
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	currencyModel := currency.GetModel(currency.QueryHelper{
		CodeAlpha: input.Currency,
	})
	if currencyModel == nil {
		efwginhelpers.ResponseError(context, cperrors.ErrorApiV1BadCurrency)
		return
	}
	if !currencyModel.IsEnabled {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorApiV1CurrencyOnMaintenance))
		return
	}

	merchantModel := merchant.GetModel(merchant.QueryHelper{
		Token: input.Token,
	})
	if merchantModel == nil {
		efwginhelpers.ResponseError(context, cperrors.ErrorApiV1BadMerchantToken)
		return
	}

	if !signVerify(string(merchantModel.SecretKey), signData{input.Token, input.Currency, input.Sign}) {
		efwginhelpers.ResponseError(context, cperrors.ErrorApiV1InvalidSign)
		return
	}

	handler := processing.NewService(merchantModel, currencyModel)
	paymentModel, err := handler.StartPayment(input.Amount)
	if err != nil {
		efwginhelpers.ResponseError(context, err)
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": payment.DTO(paymentModel)})
}
