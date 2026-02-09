package monitoring

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/helpers"
	"compay/src/models/logmerchantwallet"
	"compay/src/models/merchant"
	"compay/src/models/merchantwallet"
	"compay/src/models/payment"
	"compay/src/services/cconverter"
	"efw/components/entity"
)

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
func (that *monitoring) doSetPaymentReceived(paymentId types.PaymentId, currency entity.CurrencyCodeAlpha, txId types.PaymentTxId, amount entity.Amount) {
	amountUsd := cconverter.Convert(amount, currency)
	payment.IncreaseReceivedAmount(paymentId, amount, amountUsd)
	paymentModel := payment.GetModel(payment.QueryHelper{Id: paymentId, DontCheckMerchantId: true})

	if paymentModel.Amount != 0 && paymentModel.ReceivedAmount >= paymentModel.Amount {
		that.doSetPaymentComplete(paymentModel)
		return
	}

	if paymentModel.Status != constants.PaymentStatusFundsReceived {
		payment.UpdateStatus(paymentId, constants.PaymentStatusFundsReceived)
	}

	go that.afterSetPaymentReceived(paymentModel, amount, amountUsd, txId)
}

func (that *monitoring) afterSetPaymentReceived(paymentModel *payment.Model, amount, amountUsd entity.Amount, txId types.PaymentTxId) {
	helpers.NotificationFundsReceived(paymentModel.GetIdInt64(),
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
	if merchantModel != nil && !merchantModel.Settings.NotificationSettings.DisableFundsReceived {
		if merchantModel.Settings.NotificationSettings.TelegramChatId != 0 {
			helpers.NotificationFundsReceived(paymentModel.GetIdInt64(),
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

	merchantwallet.IncreaseReceiveAmount(paymentModel.MerchantWalletId, amount)

	logmerchantwallet.FastCreate(
		paymentModel.MerchantId,
		paymentModel.GetId(),
		0,
		paymentModel.TrxId,
		"",
		paymentModel.Currency,
		amount,
		amountUsd,
	)
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
		paymentModel.ReceivedAmount,
		paymentModel.ReceivedAmountUsd,
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
				paymentModel.ReceivedAmount,
				paymentModel.ReceivedAmountUsd,
				merchantModel.Settings.NotificationSettings.TelegramChatId,
			)
		}
	}

	merchantWalletModel := merchantwallet.GetModel(merchantwallet.QueryHelper{PaymentId: paymentModel.GetId(), MerchantId: paymentModel.MerchantId})
	merchantWalletModel.Amount += paymentModel.ReceivedAmount
	merchantWalletModel.ReceivedAmount += paymentModel.ReceivedAmount
	merchantwallet.Save(merchantWalletModel, merchantwallet.QueryHelper{})

	if paymentModel.Amount != 0 {
		logmerchantwallet.FastCreate(
			paymentModel.MerchantId,
			paymentModel.GetId(),
			0,
			paymentModel.TrxId,
			"",
			paymentModel.Currency,
			paymentModel.ReceivedAmount,
			paymentModel.ReceivedAmountUsd,
		)
	}
}

// Удаляем заявку (истекло время)
func (that *monitoring) doSetPaymentExpired(paymentId types.PaymentId) {
	paymentModel := payment.GetModel(payment.QueryHelper{Id: paymentId, DontCheckMerchantId: true})
	if paymentModel.ReceivedAmount != 0 {
		if paymentModel.ReceivedAmount >= paymentModel.Amount {
			that.doSetPaymentComplete(paymentModel)
			return
		}
	}

	status := constants.PaymentStatusExpired
	if paymentModel.ReceivedAmount != 0 {
		status = constants.PaymentStatusInsufficientFunds
	}
	payment.UpdateStatus(paymentId, status)
	that.removePaymentAsync(paymentModel.Address, paymentModel.Currency)

	if paymentModel.ReceivedAmount == 0 {
		go that.afterSetPaymentExpired(paymentModel)
	} else {
		go that.afterSetPaymentInsufficientFunds(paymentModel)
	}
}

func (that *monitoring) afterSetPaymentExpired(paymentModel *payment.Model) {
	merchantModel := merchant.GetModel(merchant.QueryHelper{Id: paymentModel.MerchantId})

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
}

func (that *monitoring) afterSetPaymentInsufficientFunds(paymentModel *payment.Model) {
	merchantModel := merchant.GetModel(merchant.QueryHelper{Id: paymentModel.MerchantId})

	helpers.NotificationInsufficientFunds(paymentModel.GetIdInt64(),
		int64(paymentModel.MerchantId),
		paymentModel.TrxId,
		string(paymentModel.Address),
		paymentModel.Currency,
		paymentModel.Amount,
		paymentModel.AmountUsd,
		paymentModel.ReceivedAmount,
		paymentModel.ReceivedAmountUsd,
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
				paymentModel.ReceivedAmount,
				paymentModel.ReceivedAmountUsd,
				merchantModel.Settings.NotificationSettings.TelegramChatId,
			)
		}
	}

	merchantWalletModel := merchantwallet.GetModel(merchantwallet.QueryHelper{PaymentId: paymentModel.GetId(), MerchantId: paymentModel.MerchantId})
	merchantWalletModel.Amount += paymentModel.ReceivedAmount
	merchantWalletModel.ReceivedAmount += paymentModel.ReceivedAmount
	merchantwallet.Save(merchantWalletModel, merchantwallet.QueryHelper{})

	logmerchantwallet.FastCreate(
		paymentModel.MerchantId,
		paymentModel.GetId(),
		0,
		paymentModel.TrxId,
		"",
		paymentModel.Currency,
		paymentModel.ReceivedAmount,
		paymentModel.ReceivedAmountUsd,
	)
}
