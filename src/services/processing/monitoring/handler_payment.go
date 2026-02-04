package monitoring

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/helpers"
	"compay/src/models/logpayment"
	"compay/src/models/merchant"
	"compay/src/models/payment"
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
				if !logPaymentModel.IsComplete {
					that.addLogPaymentAsync(logPaymentModel)
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
func (that *monitoring) removePaymentAsync(address types.PaymentAddress, codeAlpha types.CurrencyCodeAlpha) {
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

func (that *monitoring) removePayment(address types.PaymentAddress, codeAlpha types.CurrencyCodeAlpha) {
	switch codeAlpha {
	case constants.BtcCodeAlpha:
		delete(that.btc.activePayments, address)

	default:
		helpers.NotificationError(fmt.Sprint("BAD CONFIGURATION removePayment: ", codeAlpha))
		efwlog.Panic("BAD CONFIGURATION removePayment: ", codeAlpha)
	}
}

// Ставим статус заявки "Обнаружен платеж"
func (that *monitoring) doSetPaymentDetected(paymentId types.PaymentId, txId types.PaymentTxId) {
	paymentModel := payment.GetModel(payment.QueryHelper{Id: paymentId, DontCheckMerchantId: true})
	if paymentModel.Status == constants.PaymentStatusCreated {
		payment.UpdateStatus(paymentId, constants.PaymentStatusFundsDetected)
	}

	go that.afterSetPaymentDetected(paymentModel, txId)
}

func (that *monitoring) afterSetPaymentDetected(paymentModel *payment.Model, txId types.PaymentTxId) {
	helpers.NotificationDetectedTransaction(paymentModel.GetIdInt64(),
		int64(paymentModel.MerchantId),
		paymentModel.TrxId,
		string(paymentModel.Address),
		paymentModel.Currency,
		paymentModel.Amount,
		paymentModel.AmountUsd,
		string(txId),
		0,
	)

	merchantModel := merchant.GetModel(merchant.QueryHelper{Id: paymentModel.MerchantId})
	if merchantModel != nil && !merchantModel.Settings.NotificationSettings.DisableDetectedTransaction {
		if merchantModel.Settings.NotificationSettings.TelegramChatId != 0 {
			helpers.NotificationDetectedTransaction(paymentModel.GetIdInt64(),
				int64(paymentModel.MerchantId),
				paymentModel.TrxId,
				string(paymentModel.Address),
				paymentModel.Currency,
				paymentModel.Amount,
				paymentModel.AmountUsd,
				string(txId),
				merchantModel.Settings.NotificationSettings.TelegramChatId,
			)
		}
	}
}

// Обновляем сумму платежа
func (that *monitoring) doSetPaymentReceived(paymentId types.PaymentId, txId types.PaymentTxId, amount types.Amount) {
	payment.IncreaseAmount(paymentId, amount)
	paymentModel := payment.GetModel(payment.QueryHelper{Id: paymentId, DontCheckMerchantId: true})

	if paymentModel.Amount != 0 && paymentModel.RecievedAmount >= paymentModel.Amount {
		that.doSetPaymentComplete(paymentModel)
		return
	}

	if paymentModel.Status != constants.PaymentStatusFundsReceived {
		payment.UpdateStatus(paymentId, constants.PaymentStatusFundsReceived)
	}

	go that.afterSetPaymentReceived(paymentModel, amount, txId)
}

func (that *monitoring) afterSetPaymentReceived(paymentModel *payment.Model, amount types.Amount, txId types.PaymentTxId) {
	helpers.NotificationFundsRecieved(paymentModel.GetIdInt64(),
		int64(paymentModel.MerchantId),
		paymentModel.TrxId,
		string(paymentModel.Address),
		paymentModel.Currency,
		paymentModel.Amount,
		paymentModel.AmountUsd,
		amount,
		string(txId),
		0,
	)

	merchantModel := merchant.GetModel(merchant.QueryHelper{Id: paymentModel.MerchantId})
	if merchantModel != nil && !merchantModel.Settings.NotificationSettings.DisableFundsRecieved {
		if merchantModel.Settings.NotificationSettings.TelegramChatId != 0 {
			helpers.NotificationFundsRecieved(paymentModel.GetIdInt64(),
				int64(paymentModel.MerchantId),
				paymentModel.TrxId,
				string(paymentModel.Address),
				paymentModel.Currency,
				paymentModel.Amount,
				paymentModel.AmountUsd,
				amount,
				string(txId),
				merchantModel.Settings.NotificationSettings.TelegramChatId,
			)
		}
	}
}

// Обновляем сумму платежа
func (that *monitoring) doSetPaymentComplete(paymentModel *payment.Model) {
	payment.UpdateStatus(paymentModel.GetId(), constants.PaymentStatusCompleted)
	that.removePaymentAsync(paymentModel.Address, paymentModel.Currency)

	go that.afterSetPaymentComplete(paymentModel)
}

func (that *monitoring) afterSetPaymentComplete(paymentModel *payment.Model) {
	helpers.NotificationCompletedInvoice(paymentModel.GetIdInt64(),
		int64(paymentModel.MerchantId),
		paymentModel.TrxId,
		string(paymentModel.Address),
		paymentModel.Currency,
		paymentModel.Amount,
		paymentModel.AmountUsd,
		paymentModel.RecievedAmount,
		paymentModel.RecievedAmountUsd,
		0,
	)

	merchantModel := merchant.GetModel(merchant.QueryHelper{Id: paymentModel.MerchantId})
	if merchantModel != nil && !merchantModel.Settings.NotificationSettings.DisableCompletedInvoice {
		if merchantModel.Settings.NotificationSettings.TelegramChatId != 0 {
			helpers.NotificationCompletedInvoice(paymentModel.GetIdInt64(),
				int64(paymentModel.MerchantId),
				paymentModel.TrxId,
				string(paymentModel.Address),
				paymentModel.Currency,
				paymentModel.Amount,
				paymentModel.AmountUsd,
				paymentModel.RecievedAmount,
				paymentModel.RecievedAmountUsd,
				merchantModel.Settings.NotificationSettings.TelegramChatId,
			)
		}
	}
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

	status := constants.PaymentStatusExpired
	if paymentModel.RecievedAmount != 0 {
		status = constants.PaymentStatusInsufficientFunds
	}
	payment.UpdateStatus(paymentId, status)
	that.removePaymentAsync(paymentModel.Address, paymentModel.Currency)

	go that.afterSetPaymentExpired(paymentModel)
}

func (that *monitoring) afterSetPaymentExpired(paymentModel *payment.Model) {
	merchantModel := merchant.GetModel(merchant.QueryHelper{Id: paymentModel.MerchantId})

	if paymentModel.RecievedAmount == 0 {
		helpers.NotificationExpireInvoice(paymentModel.GetIdInt64(),
			int64(paymentModel.MerchantId),
			paymentModel.TrxId,
			string(paymentModel.Address),
			paymentModel.Currency,
			paymentModel.Amount,
			paymentModel.AmountUsd,
			0,
		)

		if merchantModel != nil && !merchantModel.Settings.NotificationSettings.DisableExpireInvoice {
			if merchantModel.Settings.NotificationSettings.TelegramChatId != 0 {
				helpers.NotificationExpireInvoice(paymentModel.GetIdInt64(),
					int64(paymentModel.MerchantId),
					paymentModel.TrxId,
					string(paymentModel.Address),
					paymentModel.Currency,
					paymentModel.Amount,
					paymentModel.AmountUsd,
					merchantModel.Settings.NotificationSettings.TelegramChatId,
				)
			}
		}
	} else {
		helpers.NotificationInsufficientFunds(paymentModel.GetIdInt64(),
			int64(paymentModel.MerchantId),
			paymentModel.TrxId,
			string(paymentModel.Address),
			paymentModel.Currency,
			paymentModel.Amount,
			paymentModel.AmountUsd,
			paymentModel.RecievedAmount,
			paymentModel.RecievedAmountUsd,
			0,
		)

		if merchantModel != nil && !merchantModel.Settings.NotificationSettings.DisableInsufficientFunds {
			if merchantModel.Settings.NotificationSettings.TelegramChatId != 0 {
				helpers.NotificationInsufficientFunds(paymentModel.GetIdInt64(),
					int64(paymentModel.MerchantId),
					paymentModel.TrxId,
					string(paymentModel.Address),
					paymentModel.Currency,
					paymentModel.Amount,
					paymentModel.AmountUsd,
					paymentModel.RecievedAmount,
					paymentModel.RecievedAmountUsd,
					merchantModel.Settings.NotificationSettings.TelegramChatId,
				)
			}
		}
	}
}
