package processing

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/common/types"
	"compay/src/helpers"
	"compay/src/models/merchant"
	"compay/src/models/payment"
	"compay/src/services/processing/monitoring"
	"efw/components/entity"
	"efw/core/efwrequest"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type responseCreateWallet struct {
	Data    responseCreateWalletData `json:"data"`
	Status  int                      `json:"status"`
	Success bool                     `json:"success"`
	Error   string                   `json:"error"`
	Errors  map[string][]string      `json:"errors"`
}
type responseCreateWalletData struct {
	Address       types.PaymentAddress `json:"address"`
	PrivateKey    string               `json:"private_key"`
	PrivateKeyWif string               `json:"private_key_wif"`
	PublicKey     string               `json:"public_key"`
	PublicKeyFull string               `json:"public_key_full"`
}

func (that *Processing) StartPayment(amountFloat float64) (*payment.Model, error) {
	body, _, err := efwrequest.Request(efwrequest.RequestHelper{
		Url:           that.getRequestUrl(requestCreateWallet),
		Method:        "POST",
		IsLogDisabled: true,
	})
	if err != nil {
		return nil, cperrors.ErrorProcessingStartPaymentRequestError(err.Error())
	}

	var response responseCreateWallet
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, cperrors.ErrorProcessingStartPaymentUnmarshalError(err.Error())
	}
	if !response.Success {
		return nil, cperrors.ErrorProcessingStartPaymentUnmarshalError(fmt.Sprint(response.Error, " ", response.Errors))
	}

	amount, err := that.calculateDepositAmount(amountFloat)
	if err != nil {
		return nil, err
	}

	paymentModel := &payment.Model{}
	paymentModel.TrxId = uuid.NewString()
	paymentModel.MerchantId = that.getMerchant().GetId()
	paymentModel.Status = constants.PaymentStatusCreated
	paymentModel.Currency = that.getCurrency().CodeAlpha
	paymentModel.Address = response.Data.Address
	paymentModel.Amount = amount
	paymentModel.AmountUsd = 0
	paymentModel.RequiredConfirmations = that.getMerchantCurrency().Confirms
	paymentModel.PrivateKey = entity.EncryptedString(response.Data.PrivateKey)
	paymentModel.PrivateKeyWif = entity.EncryptedString(response.Data.PrivateKeyWif)
	paymentModel.PublicKey = entity.EncryptedString(response.Data.PublicKey)
	paymentModel.PublicKeyFull = entity.EncryptedString(response.Data.PublicKeyFull)
	paymentModel.ExpiresAt = time.Now().Add(time.Second * time.Duration(that.getMerchantCurrency().InvoiceTime)).Unix()

	payment.Save(paymentModel, payment.QueryHelper{})

	// Добавляем платеж в пул мониторинга сразу после создания
	go asyncStartPayment(paymentModel, that.getMerchant())

	return paymentModel, nil
}

func asyncStartPayment(paymentModel *payment.Model, merchantModel *merchant.Model) {
	monitoring.AddPayment(paymentModel)
	helpers.NotificationCreatePayment(
		paymentModel.GetIdInt64(),
		int64(paymentModel.MerchantId),
		paymentModel.TrxId,
		string(paymentModel.Address),
		paymentModel.Currency,
		paymentModel.Amount,
		paymentModel.AmountUsd,
		0,
	)

	// сделать логи сюды
}
