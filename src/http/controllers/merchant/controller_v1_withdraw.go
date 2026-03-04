package merchant

import (
	"compay/common/cperrors"
	"compay/common/types"
	"compay/src/models/currency"
	"compay/src/models/merchant"
	"compay/src/models/merchantwallet"
	"compay/src/models/withdraw"
	"compay/src/services/processing"
	"efw/components/entity"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
)

type requestV1Withdraw struct {
	WalletId string                   `json:"wallet_id" binding:"required,len=36"`
	Amount   float64                  `json:"amount" binding:"omitempty,min=0"`
	Token    string                   `json:"token" binding:"required,len=36"`
	Currency entity.CurrencyCodeAlpha `json:"currency" binding:"required,max=20"`
	Address  string                   `json:"address" binding:"required,max=255"`
	Sign     string                   `json:"sign" binding:"required,len=128"`
}

func V1Withdraw(context *gin.Context) {
	if !isMerchantStarted() {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorApiV1Maintenance))
		return
	}

	var input requestV1Withdraw
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

	merchantWalletModel := merchantwallet.GetModel(merchantwallet.QueryHelper{
		PaymentTrxId: input.WalletId,
		MerchantId:   merchantModel.GetId(),
	})

	if merchantWalletModel == nil {
		efwginhelpers.ResponseError(context, cperrors.ErrorApiV1InvalidSign)
		return
	}

	handler := processing.NewService(merchantModel, currencyModel)
	withdrawModel, err := handler.StartWithdraw(processing.WithdrawData{
		Address:        types.WithdrawAddress(input.Address),
		Amount:         input.Amount,
		MerchantWallet: merchantWalletModel,
	})
	if err != nil {
		efwginhelpers.ResponseError(context, err)
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": withdraw.DTO(withdrawModel)})
}
