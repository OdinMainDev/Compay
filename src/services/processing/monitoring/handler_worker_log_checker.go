package monitoring

import (
	"compay/common/constants"
	"compay/src/helpers"
	"compay/src/models/logtransaction"
	"efw/core/efwlog"
	"fmt"
	"time"
)

func (that *monitoring) startLoopLogChecker() {
	ticker := time.NewTicker(that.checkLogInterval)
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
	fmt.Println("[Monitoring] Start log loop checker")

	that.logPaymentsMutex.RLock()
	if len(that.activeLogTransactions) == 0 {
		that.logPaymentsMutex.RUnlock()
		return
	}

	for _, logPaymentModel := range that.activeLogTransactions {
		that.logPaymentsJobChan <- logPaymentModel
	}
	that.logPaymentsMutex.RUnlock()
}

func (that *monitoring) processLogChecker(logPaymentModel *logtransaction.Model) {
	var confirmations int64 = 0
	currentBlock := that.btc.currentBlock.Load()
	switch logPaymentModel.Currency {
	case constants.BtcCodeAlpha:
		confirmations = currentBlock - logPaymentModel.StartBlock + 1

	default:
		helpers.NotificationError(fmt.Sprint("BAD CONFIGURATION processLogChecker: ", logPaymentModel.Currency))
		efwlog.Panic("BAD CONFIGURATION processLogChecker: ", logPaymentModel.Currency)
	}

	if confirmations >= int64(logPaymentModel.Data.Conditions.RequiredConfirmations) {
		that.doSetPaymentReceived(logPaymentModel.PaymentId, logPaymentModel.Currency, logPaymentModel.TxId, logPaymentModel.Data.TxData.Amount)

		// Маркируем лог как завершённый
		logtransaction.UpdateStatus(logPaymentModel.GetId())

		// Удаляем из активных
		that.removeLogTransactionAsync(logPaymentModel.TxId, logPaymentModel.Address)
	}
}
