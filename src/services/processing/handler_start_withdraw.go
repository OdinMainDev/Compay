package processing

import (
	"compay/common/cperrors"
	"compay/common/types"
	"compay/src/models/logtransaction"
	"compay/src/models/merchantwallet"
	"compay/src/models/payment"
	"efw/core/efwamount"
	"efw/core/efwrequest"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type WithdrawData struct {
	Address        types.PaymentAddress
	Amount         float64
	MerchantWallet *merchantwallet.Model
}

type responseWithdraw struct {
	Data    types.TxWithdrawData `json:"data"`
	Status  int                  `json:"status"`
	Success bool                 `json:"success"`
	Error   string               `json:"error"`
	Errors  map[string][]string  `json:"errors"`
}

var mutexWithdraws sync.Mutex

func (that *Processing) StartWithdraw(data WithdrawData) (*payment.Model, error) {
	amount := efwamount.ToAmount(data.Amount, that.getCurrency().MinorUnits)
	if amount < that.getMerchantCurrency().MinAmountWithdraw {
		return nil, cperrors.ErrorProcessingWithdrawAmountLessThanMinimal
	}
	if amount > that.getMerchantCurrency().MaxAmountWithdraw {
		return nil, cperrors.ErrorProcessingWithdrawAmountGreaterThanMaximum
	}

	mutexWithdraws.Lock()

	mutexWithdraws.Unlock()

	logTransactionModels := logtransaction.GetModels(logtransaction.QueryHelper{PaymentId: data.MerchantWallet.PaymentId})

	requestBody, err := that.constructWithdrawRequestBody(string(data.MerchantWallet.PrivateKeyWif), string(data.Address), amount, logTransactionModels)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(requestBody))

	body, _, err := efwrequest.Request(efwrequest.RequestHelper{
		Url:            that.getRequestUrl(requestWithdraw),
		Method:         "POST",
		IsLogDisabled:  true,
		Body:           requestBody,
		RequestTimeout: 10 * time.Second,
	})
	if err != nil {
		return nil, cperrors.ErrorProcessingStartPaymentRequestError(err.Error())
	}

	fmt.Println(string(body))

	var response responseWithdraw
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, cperrors.ErrorProcessingWithdrawUnmarshalError(err.Error())
	}
	if !response.Success {
		return nil, cperrors.ErrorProcessingWithdrawResponseError(fmt.Sprint(response.Error, " ", response.Errors))
	}

	// logWithdrawModel :=
	return nil, nil
}
