package monitoring

import (
	"compay/common/constants"
	entitylogpayment "compay/src/entities/logpayment"
	"compay/src/models/logpayment"
	"fmt"
	"time"
)

func (that *monitoring) startWorkerTxParser(id int) {
	for {
		select {
		case <-that.stopChan:
			fmt.Println("[Monitoring] STOP worker ", id, " tx parser")
			return
		case data := <-that.txParserJobChan:
			that.processTxParse(data)
		}
	}
}

func (that *monitoring) processTxParse(data txParserJob) {
	// проверяем наличие транзакции в проверках
	logPaymentModel := logpayment.GetModel(logpayment.QueryHelper{
		PaymentId: data.PaymentModel.GetId(),
		TxId:      data.TxData.TxId,
	})

	// если она есть то дальше не обрабатываем тут
	if logPaymentModel != nil {
		return
	}

	if !that.isPaymentActive(data.PaymentModel, time.Now().Unix()) {
		that.doSetPaymentExpired(data.PaymentModel.GetId())
		return
	}

	// если у платежа статус транзакции "Создана" то тогда обновляем и отправляем уведомление
	if data.PaymentModel.Status == constants.PaymentStatusCreated {
		that.doSetPaymentDetected(data.PaymentModel.GetId(), data.TxData.TxId)
	}

	logPaymentModel = &logpayment.Model{}
	logPaymentModel.PaymentId = data.PaymentModel.GetId()
	logPaymentModel.PaymentTrxId = data.PaymentModel.TrxId
	logPaymentModel.MerchantId = data.PaymentModel.MerchantId
	logPaymentModel.Currency = data.PaymentModel.Currency
	logPaymentModel.IsComplete = false
	logPaymentModel.StartBlock = data.TxData.BlockHeight
	logPaymentModel.TxId = data.TxData.TxId
	logPaymentModel.TxHash = data.TxData.TxHash
	logPaymentModel.Address = data.TxData.Address
	logPaymentModel.Data = entitylogpayment.Data{
		TxData: *data.TxData,
		Conditions: entitylogpayment.Conditions{
			RequiredConfirmations: data.PaymentModel.RequiredConfirmations,
		},
	}
	logpayment.Save(logPaymentModel, logpayment.QueryHelper{})

	that.addLogPaymentAsync(logPaymentModel)
}
