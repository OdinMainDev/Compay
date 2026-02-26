package processing

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/common/types"
	"compay/src/models/merchant"
	"compay/src/models/merchantwallet"
	"compay/src/models/payment"
	"compay/src/services/cconverter"
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
	paymentModel.AmountUsd = cconverter.Convert(amount, that.getCurrency().CodeAlpha)
	paymentModel.RequiredConfirmations = that.getMerchantCurrency().Confirms
	paymentModel.ExpiresAt = time.Now().Add(time.Second * time.Duration(that.getMerchantCurrency().InvoiceTime)).Unix()
	payment.Save(paymentModel, payment.QueryHelper{})

	merchantWalletModel := &merchantwallet.Model{}
	merchantWalletModel.MerchantId = that.getMerchant().GetId()
	merchantWalletModel.PaymentId = paymentModel.GetId()
	merchantWalletModel.Currency = that.getCurrency().CodeAlpha
	merchantWalletModel.Address = response.Data.Address
	merchantWalletModel.Amount = 0
	merchantWalletModel.ReceivedAmount = 0
	merchantWalletModel.SentAmount = 0
	merchantWalletModel.PrivateKey = entity.EncryptedString(response.Data.PrivateKey)
	merchantWalletModel.PrivateKeyWif = entity.EncryptedString(response.Data.PrivateKeyWif)
	merchantWalletModel.PublicKey = entity.EncryptedString(response.Data.PublicKey)
	merchantWalletModel.PublicKeyFull = entity.EncryptedString(response.Data.PublicKeyFull)
	merchantwallet.Save(merchantWalletModel, merchantwallet.QueryHelper{})

	paymentModel.MerchantWalletId = merchantWalletModel.GetId()
	payment.Save(paymentModel, payment.QueryHelper{})

	// Добавляем платеж в пул мониторинга сразу после создания
	go asyncStartPayment(paymentModel, that.getMerchant())

	return paymentModel, nil
}

func asyncStartPayment(paymentModel *payment.Model, merchantModel *merchant.Model) {
	monitoring.AddPayment(paymentModel)
	notifyCreatePayment(paymentModel, merchantModel)
}
