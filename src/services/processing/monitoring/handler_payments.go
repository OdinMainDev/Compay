package monitoring

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/payment"
	"efw/components/model"
	"fmt"
	"time"
)

// loadPayments загружает активные платежи
func (that *Monitoring) loadPayments() {
	now := time.Now().Unix()

	that.mutex.Lock()
	defer that.mutex.Unlock()

	count := 0
	for i := range 100000000 {
		// Получаем только активные платежи которые еще не истекли
		paymentModels := payment.GetModels(payment.QueryHelper{
			Statuses:            []types.PaymentStatus{constants.PaymentStatusCreated, constants.PaymentStatusFundsReceived},
			DontCheckMerchantId: true,
			QueryHelper: model.QueryHelper{
				HasPagination: model.HasPagination{
					PageNumber:    i,
					ElementsCount: 100,
				},
			},
		})
		if len(paymentModels) == 0 {
			break
		}

		for _, paymentModel := range paymentModels {
			// Проверяем платеж
			if !that.isPaymentActive(paymentModel, now) {
				continue
			}

			that.activePayments[paymentModel.GetIdInt64()] = paymentModel
			count++

			// Импортируем адрес в ноду
			go func(address string, id int64) {
				if err := that.btcService.ImportPaymentAddress(address, fmt.Sprintf("%d", id)); err != nil {
					fmt.Printf("[Monitoring] Failed to import address %s: %v\n", address, err)
				}
			}(paymentModel.Address, paymentModel.GetIdInt64())
		}
	}

	fmt.Printf("[Monitoring] Loaded %d active payments from DB\n", count)
}

// isPaymentActive проверяет нужно ли проверять платеж
func (that *Monitoring) isPaymentActive(paymentModel *payment.Model, now int64) bool {
	// если истек, то мы его отбрасываем и возвращаем в processing для закрытия
	if paymentModel.ExpiresAt < now {
		that.closePaymentChan <- ClosePaymentChan{
			PaymentModel: paymentModel,
			Status:       constants.PaymentStatusExpired,
		}
		return false
	}

	return true
}

// AddPayment добавляет новый платеж в пул обработки
func (that *Monitoring) addPayment(paymentModel *payment.Model) {
	that.mutex.Lock()
	defer that.mutex.Unlock()

	that.activePayments[paymentModel.GetIdInt64()] = paymentModel

	// Импортируем адрес в ноду для отслеживания
	go func(address string, trxId string) {
		if err := that.btcService.ImportPaymentAddress(address, trxId); err != nil {
			fmt.Printf("[Monitoring] Failed to import address %s for payment %s: %v\n", address, trxId, err)
		}
		fmt.Printf("[Monitoring] Payment %s added to monitoring pool\n", paymentModel.TrxId)
	}(paymentModel.Address, paymentModel.TrxId)
}

// removePayment удаляет платеж из пула
func (that *Monitoring) removePayment(paymentId int64) {
	that.mutex.Lock()
	defer that.mutex.Unlock()
	delete(that.activePayments, paymentId)
}
