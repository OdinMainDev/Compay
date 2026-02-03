package monitoring

import (
	"compay/common/constants"
	"compay/src/helpers"
	"compay/src/models/logpayment"
	"efw/core/efwlog"
	"fmt"
	"maps"
	"time"
)

func (that *monitoring) startLoopLogChecker() {
	ticker := time.NewTicker(that.checkInterval)
	// Первая проверка сразу при старте
	that.handleLoopLogChecker()

	for {
		select {
		case <-that.stopChan:
			fmt.Println("[Monitoring] STOP loop log checker")
			return
		case <-ticker.C:
			that.handleLoopLogChecker()
		}
	}
}

func (that *monitoring) startWorkerLogChecker(id int) {
	for {
		select {
		case <-that.stopChan:
			fmt.Println("[Monitoring] STOP worker ", id, " log checker")
			return
		case data := <-that.logPaymentsJobChan:
			that.processLogChecker(data)
		}
	}
}

func (that *monitoring) handleLoopLogChecker() {
	fmt.Println("[Monitoring] Start loop checker")

	that.logPaymentsMutex.RLock()
	activeLogs := make(map[string]*logpayment.Model, len(that.activeLogPayments))
	maps.Copy(activeLogs, that.activeLogPayments)
	that.logPaymentsMutex.RUnlock()

	if len(activeLogs) == 0 {
		return
	}

	for _, logPaymentModel := range activeLogs {
		that.logPaymentsJobChan <- logPaymentModel
	}
}

func (that *monitoring) processLogChecker(logPaymentModel *logpayment.Model) {
	var confirmations int64 = 0
	switch logPaymentModel.Currency {
	case constants.BtcCodeAlpha:
		currentBlock := that.btc.currentBlock.Load()
		confirmations = currentBlock - logPaymentModel.StartBlock + 1

	default:
		helpers.NotificationError(fmt.Sprint("BAD CONFIGURATION processLogChecker: ", logPaymentModel.Currency))
		efwlog.Panic("BAD CONFIGURATION processLogChecker: ", logPaymentModel.Currency)
	}

	if confirmations >= int64(logPaymentModel.Data.Conditions.RequiredConfirmations) {
		that.doSetPaymentReceived(logPaymentModel.PaymentId, logPaymentModel.TxId, logPaymentModel.Data.TxData.Amount)

		// Маркируем лог как завершённый
		logPaymentModel.Lock()
		logPaymentModel.IsComplete = true
		logpayment.Save(logPaymentModel, logpayment.QueryHelper{})
		logPaymentModel.Unlock()

		// Удаляем из активных
		that.removeLogPaymentAsync(logPaymentModel.TxId)
	}
}
