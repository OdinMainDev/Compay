package monitoring

import (
	"compay/common/constants"
	entityLogtransaction "compay/src/entities/logtransaction"
	"compay/src/models/logtransaction"
	"fmt"
	"time"
)

func (that *monitoring) startWorkerTxParser(id int) {
	for {
		select {
		case <-that.stopChan:
			that.log(fmt.Sprint("[Monitoring] STOP worker ", id, " tx parser"))
			return
		case data := <-that.txParserJobChan:
			that.processTxParse(data)
		}
	}
}

func (that *monitoring) processTxParse(data txParserJob) {
	// проверяем наличие транзакции в проверках
	logTransactionModel := logtransaction.GetModel(logtransaction.QueryHelper{
		PaymentId: data.PaymentModel.GetId(),
		TxId:      data.TxData.TxId,
	})

	// если она есть то дальше не обрабатываем тут
	if logTransactionModel != nil {
		return
	}

	if !that.isPaymentActive(data.PaymentModel, time.Now().Unix()) {
		that.doSetPaymentExpired(data.PaymentModel.GetId())
		return
	}

	// если у платежа статус транзакции "Создана" то тогда обновляем и отправляем уведомление
	if data.PaymentModel.Status == constants.PaymentStatusCreated {
		that.doSetPaymentDetected(data.PaymentModel.GetId(), data.TxData.TxId, data.TxData.Amount)
	}

	logTransactionModel = &logtransaction.Model{}
	logTransactionModel.PaymentId = data.PaymentModel.GetId()
	logTransactionModel.PaymentTrxId = data.PaymentModel.TrxId
	logTransactionModel.MerchantId = data.PaymentModel.MerchantId
	logTransactionModel.MerchantWalletId = data.PaymentModel.MerchantWalletId
	logTransactionModel.Currency = data.PaymentModel.Currency
	logTransactionModel.IsComplete = false
	logTransactionModel.StartBlock = data.TxData.BlockHeight
	logTransactionModel.TxId = data.TxData.TxId
	logTransactionModel.TxHash = data.TxData.TxHash
	logTransactionModel.Address = data.TxData.Address
	logTransactionModel.Data = entityLogtransaction.Data{
		TxData: *data.TxData,
		Conditions: entityLogtransaction.Conditions{
			RequiredConfirmations: data.PaymentModel.RequiredConfirmations,
		},
	}
	logtransaction.Save(logTransactionModel, logtransaction.QueryHelper{})

	that.addLogTransactionAsync(logTransactionModel)
}
