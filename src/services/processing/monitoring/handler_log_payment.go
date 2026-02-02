package monitoring

import (
	"compay/src/models/logpayment"
	"fmt"
)

// addLogPaymentAsync добавляет новый лог платежа в пул обработки
func (that *monitoring) addLogPaymentAsync(logPaymentModel *logpayment.Model) {
	that.logPaymentsMutex.Lock()
	that.addLogPayment(logPaymentModel)
	that.logPaymentsMutex.Unlock()

	fmt.Printf("[Monitoring] LogPayment with tx_id %s added to monitoring pool\n", logPaymentModel.TxId)
}

func (that *monitoring) addLogPayment(logPaymentModel *logpayment.Model) {
	that.activeLogPayments[logPaymentModel.TxId] = logPaymentModel
}

// removeLogPaymentAsync удаляет лог платежа из пула
func (that *monitoring) removeLogPaymentAsync(txId string) {
	that.logPaymentsMutex.Lock()
	that.removeLogPayment(txId)
	that.logPaymentsMutex.Unlock()

	fmt.Printf("[Monitoring] LogPayment with tx_id %s removed from monitoring pool\n", txId)
}

func (that *monitoring) removeLogPayment(txId string) {
	delete(that.activeLogPayments, txId)
}
