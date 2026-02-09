package monitoring

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/helpers"
	"compay/src/models/logtransaction"
	"compay/src/models/payment"
	"efw/components/entity"
	"efw/components/model"
	"efw/core/efwlog"
	"fmt"
	"time"
)

func (that *monitoring) loadPayments() {
	now := time.Now().Unix()

	count := 0
	countLogs := 0
	for i := range 100000000 {
		// Получаем только активные платежи которые еще не истекли
		paymentModels := payment.GetModels(payment.QueryHelper{
			Statuses: []types.PaymentStatus{
				constants.PaymentStatusCreated,
				constants.PaymentStatusFundsDetected,
				constants.PaymentStatusFundsReceived,
			},
			DontCheckMerchantId: true,
			QueryHelper: model.QueryHelper{
				HasPagination: model.HasPagination{
					HasPagination: true,
					PageNumber:    i + 1,
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
				that.doSetPaymentExpired(paymentModel.GetId())
				continue
			}

			that.addPayment(paymentModel)
			count++

			logTransactionModels := logtransaction.GetModels(logtransaction.QueryHelper{
				PaymentId: paymentModel.GetId(),
				QueryHelper: model.QueryHelper{
					HasPagination: model.HasPagination{
						HasLimit:      true,
						ElementsCount: 10,
					},
				},
			})
			for _, logTransactionModel := range logTransactionModels {
				if !logTransactionModel.IsComplete {
					that.addLogTransactionAsync(logTransactionModel)
					countLogs++
				}
			}
		}
	}

	helpers.NotificationDebug(fmt.Sprintf("[Monitoring] Загружено %d активных платежей", count))
	helpers.NotificationDebug(fmt.Sprintf("[Monitoring] Загружено %d непроверенных транзакций", countLogs))
}

// isPaymentActive проверяет нужно ли проверять платеж
func (that *monitoring) isPaymentActive(paymentModel *payment.Model, now int64) bool {
	// если истек, то мы его отбрасываем и возвращаем в processing для закрытия
	if paymentModel.ExpiresAt < now {
		return false
	}

	return true
}

// addPaymentAsync добавляет новый платеж в пул обработки
func (that *monitoring) addPaymentAsync(paymentModel *payment.Model) {
	switch paymentModel.Currency {
	case constants.BtcCodeAlpha:
		that.btc.mutex.Lock()
		defer that.btc.mutex.Unlock()

	default:
		helpers.NotificationError(fmt.Sprint("BAD CONFIGURATION addPaymentAsync: ", paymentModel.Currency))
		efwlog.Panic("BAD CONFIGURATION addPaymentAsync: ", paymentModel.Currency)
	}
	that.addPayment(paymentModel)

	that.log(fmt.Sprintf("[Monitoring] Payment with wallet %s added to monitoring pool", paymentModel.Address))
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
func (that *monitoring) removePaymentAsync(address types.PaymentAddress, codeAlpha entity.CurrencyCodeAlpha) {
	switch codeAlpha {
	case constants.BtcCodeAlpha:
		that.btc.mutex.Lock()
		defer that.btc.mutex.Unlock()

	default:
		helpers.NotificationError(fmt.Sprint("BAD CONFIGURATION removePaymentAsync: ", codeAlpha))
		efwlog.Panic("BAD CONFIGURATION removePaymentAsync: ", codeAlpha)
	}
	that.removePayment(address, codeAlpha)

	that.log(fmt.Sprintf("[Monitoring] Payment with wallet %s removed from monitoring pool", address))
}

func (that *monitoring) removePayment(address types.PaymentAddress, codeAlpha entity.CurrencyCodeAlpha) {
	switch codeAlpha {
	case constants.BtcCodeAlpha:
		delete(that.btc.activePayments, address)

	default:
		helpers.NotificationError(fmt.Sprint("BAD CONFIGURATION removePayment: ", codeAlpha))
		efwlog.Panic("BAD CONFIGURATION removePayment: ", codeAlpha)
	}
}
