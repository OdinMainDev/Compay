package processing

import (
	"compay/common/cperrors"
	"compay/common/types"
	entityLogtransaction "compay/src/entities/logtransaction"
	"compay/src/models/logtransaction"
	"compay/src/models/merchantwallet"
	"compay/src/models/withdraw"
	"compay/src/services/cconverter"
	"efw/components/entity"
	"efw/core/efwamount"
	"efw/core/efwrequest"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/google/uuid"
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

func (that *Processing) StartWithdraw(data WithdrawData) error {
	amount := efwamount.ToAmount(data.Amount, that.getCurrency().MinorUnits)
	if amount < that.getMerchantCurrency().MinAmountWithdraw {
		return cperrors.ErrorProcessingWithdrawAmountLessThanMinimal
	}
	if amount > that.getMerchantCurrency().MaxAmountWithdraw {
		return cperrors.ErrorProcessingWithdrawAmountGreaterThanMaximum
	}
	if amount > data.MerchantWallet.Amount {
		return cperrors.ErrorProcessingWithdrawInsufficientFunds
	}

	mutexWithdraws.Lock()

	withdrawModel := withdraw.GetModel(withdraw.QueryHelper{
		MerchantWalletId: data.MerchantWallet.GetId(),
		IsInComplete:     true,
	})
	if withdrawModel != nil {
		mutexWithdraws.Unlock()
		return cperrors.ErrorProcessingWithdrawHasUncompletedWithdraw
	}
	withdrawModel = &withdraw.Model{}
	withdrawModel.TrxId = uuid.NewString()
	withdrawModel.MerchantId = that.getMerchant().GetId()
	withdrawModel.MerchantWalletId = data.MerchantWallet.GetId()
	withdrawModel.Currency = that.getCurrency().CodeAlpha
	withdrawModel.IsComplete = false
	withdrawModel.Address = types.WithdrawAddress(data.Address)
	withdrawModel.Amount = amount
	withdrawModel.AmountUsd = cconverter.Convert(amount, that.getCurrency().CodeAlpha)
	withdraw.Save(withdrawModel, withdraw.QueryHelper{})

	mutexWithdraws.Unlock()

	go that.asyncStartWithdraw(data, withdrawModel, amount)
	return nil
}

func (that *Processing) asyncStartWithdraw(data WithdrawData, withdrawModel *withdraw.Model, amount entity.Amount) (rE error) {
	defer func() {
		if rE != nil {
			withdrawModel.Error = rE.Error()
			withdrawModel.IsComplete = true
			withdraw.Save(withdrawModel, withdraw.QueryHelper{})
		}
	}()
	logTransactionModels := logtransaction.GetModels(logtransaction.QueryHelper{MerchantWalletId: data.MerchantWallet.GetId(), IsComplete: true, IsUnSpent: true})

	fmt.Println(logTransactionModels)
	requestBody, err := that.constructWithdrawRequestBody(string(data.MerchantWallet.PrivateKeyWif), string(data.Address), amount, logTransactionModels)
	if err != nil {
		return err
	}
	fmt.Println(string(requestBody))

	body, _, err := efwrequest.Request(efwrequest.RequestHelper{
		Url:            that.getRequestUrl(requestWithdraw),
		Method:         "POST",
		IsLogDisabled:  true,
		Body:           requestBody,
		RequestTimeout: 10,
	})
	if err != nil {
		return cperrors.ErrorProcessingStartPaymentRequestError(err.Error())
	}

	fmt.Println(string(body))

	var response responseWithdraw
	if err := json.Unmarshal(body, &response); err != nil {
		return cperrors.ErrorProcessingWithdrawUnmarshalError(err.Error())
	}
	if !response.Success {
		return cperrors.ErrorProcessingWithdrawResponseError(fmt.Sprint(response.Error, " ", response.Errors))
	}

	// Помечаем только использованные UTXO как потраченные
	for _, usedUTXO := range response.Data.UsedUTXOs {
		for _, logTransactionModel := range logTransactionModels {
			if string(logTransactionModel.TxId) == usedUTXO.TxID &&
				logTransactionModel.Data.TxData.VoutPosition == int(usedUTXO.Vout) {
				logTransactionModel.IsSpent = true
				logtransaction.Save(logTransactionModel, logtransaction.QueryHelper{})
				break
			}
		}
	}

	logTransactionModel := &logtransaction.Model{}
	logTransactionModel.WithdrawId = withdrawModel.GetId()
	logTransactionModel.WithdrawTrxId = withdrawModel.TrxId
	logTransactionModel.MerchantWalletId = data.MerchantWallet.GetId()
	logTransactionModel.MerchantId = that.getMerchant().GetId()
	logTransactionModel.Currency = that.getCurrency().CodeAlpha
	logTransactionModel.IsComplete = true
	logTransactionModel.IsSpent = response.Data.ChangeAmount == 0
	logTransactionModel.TxId = response.Data.TxId
	logTransactionModel.TxHash = response.Data.TxHash
	logTransactionModel.Address = response.Data.Address
	logTransactionModel.Data = entityLogtransaction.Data{
		TxWithdrawData: response.Data,
	}
	logtransaction.Save(logTransactionModel, logtransaction.QueryHelper{})

	withdrawModel.IsComplete = true
	withdraw.Save(withdrawModel, withdraw.QueryHelper{})

	merchantwallet.IncreaseSentAmount(data.MerchantWallet.GetId(), amount)
	return nil
}
