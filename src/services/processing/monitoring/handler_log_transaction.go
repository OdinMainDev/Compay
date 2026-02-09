package monitoring

import (
	"compay/common/types"
	"compay/src/models/logtransaction"
	"fmt"
)

// addLogTransactionAsync добавляет новый лог платежа в пул обработки
func (that *monitoring) addLogTransactionAsync(logTransactionModel *logtransaction.Model) {
	that.logTransactionsMutex.Lock()
	that.addLogTransaction(logTransactionModel)
	that.logTransactionsMutex.Unlock()

	that.log(fmt.Sprintf("[Monitoring] LogTransaction with tx_id %s added to monitoring pool", logTransactionModel.TxId))
}

func (that *monitoring) addLogTransaction(logTransactionModel *logtransaction.Model) {
	that.activeLogTransactions[toTxId(logTransactionModel.TxId, logTransactionModel.Address)] = logTransactionModel
}

// removeLogTransactionAsync удаляет лог платежа из пула
func (that *monitoring) removeLogTransactionAsync(txId types.PaymentTxId, address types.PaymentAddress) {
	that.logTransactionsMutex.Lock()
	that.removeLogTransaction(txId, address)
	that.logTransactionsMutex.Unlock()

	that.log(fmt.Sprintf("[Monitoring] LogTransaction with tx_id %s removed from monitoring pool", txId))
}

func (that *monitoring) removeLogTransaction(txId types.PaymentTxId, address types.PaymentAddress) {
	delete(that.activeLogTransactions, toTxId(txId, address))
}
