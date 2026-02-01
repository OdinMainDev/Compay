package monitoring

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/helpers"
	"compay/src/models/payment"
	"efw/components/model"
	"efw/core/efwlog"
	"fmt"
	"time"
)

func (that *monitoring) loadPayments() {
	now := time.Now().Unix()

	count := 0
	for i := range 100000000 {
		// Получаем только активные платежи которые еще не истекли
		paymentModels := payment.GetModels(payment.QueryHelper{
			Statuses: []types.PaymentStatus{
				constants.PaymentStatusCreated,
				constants.PaymentStatusFundsDetected,
			},
			DontCheckMerchantId: true,
			QueryHelper: model.QueryHelper{
				HasPagination: model.HasPagination{
					HasPagination: true,
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

			that.addPayment(paymentModel)
			count++
		}
	}

	fmt.Printf("[Monitoring] Loaded %d active payments from DB\n", count)
}

// isPaymentActive проверяет нужно ли проверять платеж
func (that *monitoring) isPaymentActive(paymentModel *payment.Model, now int64) bool {
	// если истек, то мы его отбрасываем и возвращаем в processing для закрытия
	if paymentModel.ExpiresAt < now {
		// that.closePaymentChan <- ClosePaymentChan{
		// 	PaymentModel: paymentModel,
		// 	Status:       constants.PaymentStatusExpired,
		// }
		return false
	}

	return true
}

// addPaymentAsync добавляет новый платеж в пул обработки
func (that *monitoring) addPaymentAsync(paymentModel *payment.Model) {
	that.paymentsMutex.Lock()
	that.addPayment(paymentModel)
	that.paymentsMutex.Unlock()

	fmt.Printf("[Monitoring] Payment with wallet %s added to monitoring pool\n", paymentModel.Address)
}

func (that *monitoring) addPayment(paymentModel *payment.Model) {
	switch paymentModel.Currency {
	case constants.BtcCodeAlpha:
		that.btc.activePayments[paymentModel.Address] = paymentModel

	default:
		helpers.NotificationError(fmt.Sprint("BAD CONFIGURATION addPayment: ", paymentModel.Currency))
		efwlog.Panic("BAD CONFIGURATION addPayment: ", paymentModel.Currency)
	}
}

// removePaymentAsync удаляет платеж из пула
func (that *monitoring) removePaymentAsync(address string, codeAlpha types.CurrencyCodeAlpha) {
	that.paymentsMutex.Lock()
	that.removePayment(address, codeAlpha)
	that.paymentsMutex.Unlock()

	fmt.Printf("[Monitoring] Payment with wallet %s added to monitoring pool\n", address)
}

func (that *monitoring) removePayment(address string, codeAlpha types.CurrencyCodeAlpha) {
	switch codeAlpha {
	case constants.BtcCodeAlpha:
		delete(that.btc.activePayments, address)

	default:
		helpers.NotificationError(fmt.Sprint("BAD CONFIGURATION removePayment: ", codeAlpha))
		efwlog.Panic("BAD CONFIGURATION removePayment: ", codeAlpha)
	}
}
