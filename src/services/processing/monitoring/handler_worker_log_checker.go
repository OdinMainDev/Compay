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
			that.log("[Monitoring] STOP loop log checker")
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
			that.log(fmt.Sprint("[Monitoring] STOP worker ", id, " log checker"))
			return
		case data := <-that.logTransactionsJobChan:
			that.processLogChecker(data)
		}
	}
}

func (that *monitoring) handleLoopLogChecker() {
	that.log("[Monitoring] Start log loop checker")

	that.logTransactionsMutex.RLock()
	if len(that.activeLogTransactions) == 0 {
		that.logTransactionsMutex.RUnlock()
		return
	}

	for _, logTransactionModel := range that.activeLogTransactions {
		that.logTransactionsJobChan <- logTransactionModel
	}
	that.logTransactionsMutex.RUnlock()
}

func (that *monitoring) processLogChecker(logTransactionModel *logtransaction.Model) {
	var confirmations int64 = 0
	currentBlock := that.btc.currentBlock.Load()
	switch logTransactionModel.Currency {
	case constants.BtcCodeAlpha:
		confirmations = currentBlock - logTransactionModel.StartBlock + 1

	default:
		helpers.NotificationError(fmt.Sprint("BAD CONFIGURATION processLogChecker: ", logTransactionModel.Currency))
		efwlog.Panic("BAD CONFIGURATION processLogChecker: ", logTransactionModel.Currency)
	}

	if confirmations >= int64(logTransactionModel.Data.Conditions.RequiredConfirmations) {
		that.doSetPaymentReceived(logTransactionModel.PaymentId, logTransactionModel.Currency, logTransactionModel.TxId, logTransactionModel.Data.TxData.Amount)

		// Маркируем лог как завершённый
		logtransaction.UpdateStatus(logTransactionModel.GetId())

		// Удаляем из активных
		that.removeLogTransactionAsync(logTransactionModel.TxId, logTransactionModel.Address)
	}
}
