package merchant

import (
	"compay/common/cperrors"
	"compay/common/types"
	"compay/src/models/merchant"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
)

type requestV1PaymentCreate struct {
	Amount   types.Amount            `json:"amount" binding:"required,gt=0"`
	Token    string                  `json:"token" binding:"required,len=36"`
	Currency types.CurrencyCodeAlpha `json:"currency" binding:"required,currencyExist"`
	Sign     string                  `json:"sign" binding:"required,len=128"`
}

func V1PaymentCreate(context *gin.Context) {
	var input requestV1PaymentCreate
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	merchantModel := merchant.GetModel(merchant.QueryHelper{
		Token: input.Token,
	})
	if merchantModel == nil {
		efwginhelpers.ResponseError(context, cperrors.ErrorApiV1BadMerchantToken)
		return
	}

	if !signVerify(merchantModel, signData{input.Token, input.Currency, input.Sign}) {
		efwginhelpers.ResponseError(context, cperrors.ErrorApiV1InvalidSign)
		return
	}

	// 1) проверяем валюты мерчанта, у него должны быть заполнены конфирмы, мин и макс сумма
	// 1.1) если не заполнены - сначала создаем их, потом используем
	// 2)

	efwginhelpers.ResponseSuccess(context, gin.H{
		"address": "testaddress",
		"amount":  input.Amount,
	})
}
