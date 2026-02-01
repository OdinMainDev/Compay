package monitoring

import (
	"compay/common/constants"
	"compay/src/models/logpayment"
	"compay/src/models/merchantcurrency"
	"efw/components/model"
	"fmt"
)

func (that *monitoring) startWorkerJob(id int) {
	for {
		select {
		case <-that.stopChan:
			fmt.Println("[Monitoring] STOP worker ", id, " job")
			return
		case data := <-that.jobsChan:
			that.processJob(data)
		}
	}
}

func (that *monitoring) processJob(data job) {
	// обновления полей модели платежа
	paymentStatus := data.PaymentModel.Status

	// данные транзакций, мерчанта валюты
	var merchantCurrencyModel *merchantcurrency.Model
	logPaymentModels := logpayment.GetModels(logpayment.QueryHelper{
		PaymentId: data.PaymentModel.GetId(),
		QueryHelper: model.QueryHelper{
			HasPagination: model.HasPagination{
				HasLimit:      true,
				ElementsCount: 10,
			},
		},
	})

	// уведомления
	needFundsDetectedNotify := false

	// проверяем транзакцию и в случае чего выставляем статус
	isTransactionCompleted := false
	if data.TxData.Confirmations > 0 {
		if paymentStatus != constants.PaymentStatusFundsDetected {
			paymentStatus = constants.PaymentStatusFundsDetected
			needFundsReceivedNotify = true
		}
		merchantCurrencyModel = that.getMerchantCurrencyModel(merchantCurrencyModel, data.PaymentModel.MerchantId, data.PaymentModel.Currency)
		if merchantCurrencyModel.MinConfirms >= data.TxData.Confirmations {
			isTransactionCompleted = true
		}
	}

	if isTransactionCompleted {

	}
}
