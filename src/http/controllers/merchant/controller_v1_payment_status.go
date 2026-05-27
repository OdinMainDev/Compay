package merchant

import (
	"compay/common/cperrors"
	"compay/src/models/merchant"
	"compay/src/models/payment"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestV1PaymentStatus struct {
	Token string `json:"token" binding:"required,len=36"`
	Sign  string `json:"sign" binding:"required,len=128"`
	TrxId string `json:"trx_id" gorm:"column:trx_id" binding:"required,len=36"`
}

func V1PaymentStatus(context *gin.Context) {
	if !isMerchantStarted() {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorApiV1Maintenance))
		return
	}

	var input requestV1PaymentStatus
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

	paymentModel := payment.GetModel(payment.QueryHelper{
		TrxId:      input.TrxId,
		MerchantId: merchantModel.GetId(),
	})
	if paymentModel == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": payment.DTO(paymentModel)})
}
