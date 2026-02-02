package monitoring

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/helpers"
	"compay/src/models/logpayment"
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
				constants.PaymentStatusFundsReceived,
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
				that.doSetPaymentExpired(paymentModel.GetId())
				continue
			}

			that.addPayment(paymentModel)
			count++

			logPaymentModels := logpayment.GetModels(logpayment.QueryHelper{
				PaymentId: paymentModel.GetId(),
				QueryHelper: model.QueryHelper{
					HasPagination: model.HasPagination{
						HasLimit:      true,
						ElementsCount: 10,
					},
				},
			})
			for _, logPaymentModel := range logPaymentModels {
				that.addLogPaymentAsync(logPaymentModel)
			}
		}
	}

	fmt.Printf("[Monitoring] Loaded %d active payments from DB\n", count)
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
	switch codeAlpha {
	case constants.BtcCodeAlpha:
		that.btc.mutex.Lock()
		defer that.btc.mutex.Unlock()

	default:
		helpers.NotificationError(fmt.Sprint("BAD CONFIGURATION removePaymentAsync: ", codeAlpha))
		efwlog.Panic("BAD CONFIGURATION removePaymentAsync: ", codeAlpha)
	}
	that.removePayment(address, codeAlpha)

	fmt.Printf("[Monitoring] Payment with wallet %s removed from monitoring pool\n", address)
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

// Ставим статус заявки "Обнаружен платеж"
func (that *monitoring) doSetPaymentDetected(paymentId types.PaymentId, txId string) {
	paymentModel := payment.GetModel(payment.QueryHelper{Id: paymentId, DontCheckMerchantId: true})
	paymentModel.Lock()
	paymentModel.Status = constants.PaymentStatusFundsDetected
	payment.Save(paymentModel, payment.QueryHelper{})
	paymentModel.Unlock()

	go that.afterSetPaymentDetected(paymentModel, txId)
}

func (that *monitoring) afterSetPaymentDetected(paymentModel *payment.Model, txId string) {
	helpers.NotificationDetectedPayment(paymentModel.GetIdInt64(),
		int64(paymentModel.MerchantId),
		paymentModel.TrxId,
		paymentModel.Address,
		paymentModel.Currency,
		paymentModel.Amount,
		paymentModel.AmountUsd,
		txId,
		0,
	)
}

// Обновляем сумму платежа
func (that *monitoring) doSetPaymentReceived(paymentId types.PaymentId, txId string, amount types.Amount) {
	paymentModel := payment.GetModel(payment.QueryHelper{Id: paymentId, DontCheckMerchantId: true})
	paymentModel.Lock()
	paymentModel.RecievedAmount += amount
	paymentModel.Status = constants.PaymentStatusFundsReceived

	if paymentModel.Amount != 0 && paymentModel.RecievedAmount > paymentModel.Amount {
		paymentModel.Unlock()
		that.doSetPaymentComplete(paymentModel)
		return
	}

	payment.Save(paymentModel, payment.QueryHelper{})
	paymentModel.Unlock()

	go that.afterSetPaymentReceived(paymentModel, txId)
}

func (that *monitoring) afterSetPaymentReceived(paymentModel *payment.Model, txId string) {
	helpers.NotificationFundsRecivedPayment(paymentModel.GetIdInt64(),
		int64(paymentModel.MerchantId),
		paymentModel.TrxId,
		paymentModel.Address,
		paymentModel.Currency,
		paymentModel.Amount,
		paymentModel.AmountUsd,
		paymentModel.RecievedAmount,
		paymentModel.RecievedAmountUsd,
		txId,
		0,
	)
}

// Обновляем сумму платежа
func (that *monitoring) doSetPaymentComplete(paymentModel *payment.Model) {
	paymentModel.Lock()
	paymentModel.Status = constants.PaymentStatusCompleted
	payment.Save(paymentModel, payment.QueryHelper{})
	paymentModel.Unlock()
	that.removePaymentAsync(paymentModel.Address, paymentModel.Currency)

	go that.afterSetPaymentComplete(paymentModel)
}

func (that *monitoring) afterSetPaymentComplete(paymentModel *payment.Model) {
	helpers.NotificationCompletedPayment(paymentModel.GetIdInt64(),
		int64(paymentModel.MerchantId),
		paymentModel.TrxId,
		paymentModel.Address,
		paymentModel.Currency,
		paymentModel.Amount,
		paymentModel.AmountUsd,
		paymentModel.RecievedAmount,
		paymentModel.RecievedAmountUsd,
		0,
	)
}

// Удаляем заявку (истекло время)
func (that *monitoring) doSetPaymentExpired(paymentId types.PaymentId) {
	paymentModel := payment.GetModel(payment.QueryHelper{Id: paymentId, DontCheckMerchantId: true})
	if paymentModel.RecievedAmount != 0 {
		if paymentModel.RecievedAmount >= paymentModel.Amount {
			that.doSetPaymentComplete(paymentModel)
			return
		}
	}

	paymentModel.Lock()
	if paymentModel.RecievedAmount == 0 {
		paymentModel.Status = constants.PaymentStatusExpired
	} else {
		paymentModel.Status = constants.PaymentStatusInsufficientFunds
	}
	payment.Save(paymentModel, payment.QueryHelper{})
	paymentModel.Unlock()
	that.removePaymentAsync(paymentModel.Address, paymentModel.Currency)

	go that.afterSetPaymentExpired(paymentModel)
}

func (that *monitoring) afterSetPaymentExpired(paymentModel *payment.Model) {
	if paymentModel.RecievedAmount == 0 {
		helpers.NotificationExpirePayment(paymentModel.GetIdInt64(),
			int64(paymentModel.MerchantId),
			paymentModel.TrxId,
			paymentModel.Address,
			paymentModel.Currency,
			paymentModel.Amount,
			paymentModel.AmountUsd,
			0,
		)
	} else {
		helpers.NotificationInsufficientFundsPayment(paymentModel.GetIdInt64(),
			int64(paymentModel.MerchantId),
			paymentModel.TrxId,
			paymentModel.Address,
			paymentModel.Currency,
			paymentModel.Amount,
			paymentModel.AmountUsd,
			paymentModel.RecievedAmount,
			paymentModel.RecievedAmountUsd,
			0,
		)
	}
}
