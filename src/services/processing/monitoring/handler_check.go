package monitoring

import (
	"compay/src/models/payment"
	"fmt"
)

// processPayment обрабатывает один платеж
func (that *Monitoring) processPayment(paymentModel *payment.Model) {
	fmt.Println("START CHECK PAYMENT")
	// // Проверяем статус платежа через BTC ноду
	// statusInfo, err := that.btcService.CheckPaymentStatus(paymentModel)
	// if err != nil {
	// 	fmt.Printf("[Monitoring] Error checking payment %d: %v\n", paymentModel.GetIdInt64(), err)
	// 	return
	// }

	// // Обновляем платеж
	// needsUpdate := false
	// oldStatus := paymentModel.Status

	// // Обновляем полученную сумму
	// if statusInfo.TotalReceived != paymentModel.RecievedAmount {
	// 	paymentModel.RecievedAmount = statusInfo.TotalReceived
	// 	needsUpdate = true
	// }

	// // Обновляем статус если изменился
	// if statusInfo.SuggestedStatus != paymentModel.Status {
	// 	paymentModel.Status = statusInfo.SuggestedStatus
	// 	needsUpdate = true
	// 	log.Printf("[Monitoring] Payment %d status changed: %s -> %s (received: %d/%d satoshi, txs: %d)",
	// 		job.PaymentID, oldStatus, statusInfo.SuggestedStatus,
	// 		statusInfo.TotalReceived, statusInfo.RequiredAmount, statusInfo.TransactionCount)
	// }

	// // Помечаем неактивным если завершен
	// if paymentModel.Status == constants.PaymentStatusCompleted || paymentModel.Status == constants.PaymentStatusExpired {
	// 	paymentModel.IsActive = false
	// 	needsUpdate = true
	// 	log.Printf("[Monitoring] Payment %d completed (status: %s)", job.PaymentID, paymentModel.Status)
	// }

	// // Сохраняем изменения
	// if needsUpdate {
	// 	payment.Save(paymentModel, payment.QueryHelper{})
	// }

	// // Убираем из пула если обработка завершена
	// if !statusInfo.ShouldContinueMonitoring || !paymentModel.IsActive {
	// 	that.removePayment(job.PaymentID)
	// 	log.Printf("[Monitoring] Payment %d removed from monitoring pool", job.PaymentID)
	// }
}
