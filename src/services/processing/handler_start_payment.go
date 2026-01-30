package processing

import (
	"compay/common/constants"
	"compay/src/models/payment"
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
	Address       string `json:"address"`
	PrivateKey    string `json:"private_key"`
	PrivateKeyWif string `json:"private_key_wif"`
	PublicKey     string `json:"public_key"`
	PublicKeyFull string `json:"public_key_full"`
}

func (that *Processing) StartPayment(amountFloat float64) (*payment.Model, error) {
	body, _, err := efwrequest.Request(efwrequest.RequestHelper{
		Url:           that.getRequestUrl(requestCreateWallet),
		Method:        "POST",
		IsLogDisabled: true,
	})
	if err != nil {
		return nil, ErrorStartPaymentRequestError(err.Error())
	}

	var response responseCreateWallet
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, ErrorStartPaymentUnmarshalError(err.Error())
	}
	if !response.Success {
		return nil, ErrorStartPaymentUnmarshalError(fmt.Sprint(response.Error, " ", response.Errors))
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
	paymentModel.PrivateKey = entity.EncryptedString(response.Data.PrivateKey)
	paymentModel.PrivateKeyWif = entity.EncryptedString(response.Data.PrivateKeyWif)
	paymentModel.PublicKey = entity.EncryptedString(response.Data.PublicKey)
	paymentModel.PublicKeyFull = entity.EncryptedString(response.Data.PublicKeyFull)
	paymentModel.IsActive = true
	paymentModel.ExpiresAt = time.Now().Add(time.Hour * 3).Unix()

	payment.Save(paymentModel, payment.QueryHelper{})

	return paymentModel, nil
}
