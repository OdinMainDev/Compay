package monitoring

import (
	"compay/common/types"
	"compay/src/models/logtransaction"
	"fmt"
)

// addLogTransactionAsync добавляет новый лог платежа в пул обработки
func (that *monitoring) addLogTransactionAsync(logPaymentModel *logtransaction.Model) {
	that.logPaymentsMutex.Lock()
	that.addLogTransaction(logPaymentModel)
	that.logPaymentsMutex.Unlock()

	fmt.Printf("[Monitoring] LogTransaction with tx_id %s added to monitoring pool\n", logPaymentModel.TxId)
}

func (that *monitoring) addLogTransaction(logPaymentModel *logtransaction.Model) {
	that.activeLogTransactions[toTxId(logPaymentModel.TxId, logPaymentModel.Address)] = logPaymentModel
}

// removeLogTransactionAsync удаляет лог платежа из пула
func (that *monitoring) removeLogTransactionAsync(txId types.PaymentTxId, address types.PaymentAddress) {
	that.logPaymentsMutex.Lock()
	that.removeLogTransaction(txId, address)
	that.logPaymentsMutex.Unlock()

	fmt.Printf("[Monitoring] LogTransaction with tx_id %s removed from monitoring pool\n", txId)
}

func (that *monitoring) removeLogTransaction(txId types.PaymentTxId, address types.PaymentAddress) {
	delete(that.activeLogTransactions, toTxId(txId, address))
}
