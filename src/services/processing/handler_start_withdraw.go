package processing

import (
	"compay/common/cperrors"
	"compay/common/types"
	"compay/src/models/logpayment"
	"compay/src/models/merchantwallet"
	"compay/src/models/payment"
	"efw/core/efwamount"
	"efw/core/efwrequest"
	"fmt"
)

type WithdrawData struct {
	Address        types.PaymentAddress
	Amount         float64
	MerchantWallet *merchantwallet.Model
}

func (that *Processing) StartWithdraw(data WithdrawData) (*payment.Model, error) {
	amount := efwamount.ToAmount(data.Amount, that.getCurrency().MinorUnits)
	if amount < that.getMerchantCurrency().MinAmountWithdraw {
		return nil, cperrors.ErrorApiV1WithdrawAmountLessThanMinimal
	}
	if amount > that.getMerchantCurrency().MaxAmountWithdraw {
		return nil, cperrors.ErrorApiV1WithdrawAmountGreaterThanMaximum
	}

	logPaymentModels := logpayment.GetModels(logpayment.QueryHelper{PaymentId: data.MerchantWallet.PaymentId})

	requestBody, err := that.constructWithdrawRequestBody(string(data.MerchantWallet.PrivateKeyWif), string(data.Address), amount, logPaymentModels)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(requestBody))

	body, _, err := efwrequest.Request(efwrequest.RequestHelper{
		Url:           that.getRequestUrl(requestWithdraw),
		Method:        "POST",
		IsLogDisabled: true,
		Body:          requestBody,
	})
	if err != nil {
		return nil, cperrors.ErrorProcessingStartPaymentRequestError(err.Error())
	}

	fmt.Println(string(body))
	return nil, nil
}
