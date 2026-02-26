package processing

import (
	"compay/src/helpers"
	"compay/src/models/merchant"
	"compay/src/models/payment"
	"compay/src/models/withdraw"
)

func notifyCreatePayment(paymentModel *payment.Model, merchantModel *merchant.Model) {
	helpers.NotificationCreateInvoice(
		paymentModel.GetIdInt64(),
		int64(paymentModel.MerchantId),
		paymentModel.TrxId,
		string(paymentModel.Address),
		paymentModel.Currency,
		paymentModel.Amount,
		paymentModel.AmountUsd,
		0,
	)

	// сделать логи сюды
	if !merchantModel.Settings.NotificationSettings.DisableCreateInvoice {
		if merchantModel.Settings.NotificationSettings.TelegramChatId != 0 {
			helpers.NotificationCreateInvoice(
				paymentModel.GetIdInt64(),
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

func notifyWithdrawStatus(withdrawModel *withdraw.Model, merchantModel *merchant.Model, isFinish bool) {
	helpers.NotificationWithdrawStatus(
		withdrawModel.GetIdInt64(),
		int64(withdrawModel.MerchantId),
		withdrawModel.TrxId,
		string(withdrawModel.Address),
		withdrawModel.Currency,
		withdrawModel.Amount,
		withdrawModel.AmountUsd,
		isFinish,
		0,
	)

	// сделать логи сюды
	if !merchantModel.Settings.NotificationSettings.DisableWithdraws {
		if merchantModel.Settings.NotificationSettings.TelegramChatId != 0 {
			helpers.NotificationWithdrawStatus(
				withdrawModel.GetIdInt64(),
				int64(withdrawModel.MerchantId),
				withdrawModel.TrxId,
				string(withdrawModel.Address),
				withdrawModel.Currency,
				withdrawModel.Amount,
				withdrawModel.AmountUsd,
				isFinish,
				merchantModel.Settings.NotificationSettings.TelegramChatId,
			)
		}
	}
}

func notifyCreateWithdrawError(withdrawModel *withdraw.Model, merchantModel *merchant.Model, err string) {
	helpers.NotificationWithdrawError(
		withdrawModel.GetIdInt64(),
		int64(withdrawModel.MerchantId),
		withdrawModel.TrxId,
		string(withdrawModel.Address),
		withdrawModel.Currency,
		withdrawModel.Amount,
		withdrawModel.AmountUsd,
		err,
		0,
	)

	// сделать логи сюды
	if !merchantModel.Settings.NotificationSettings.DisableWithdraws {
		if merchantModel.Settings.NotificationSettings.TelegramChatId != 0 {
			helpers.NotificationWithdrawError(
				withdrawModel.GetIdInt64(),
				int64(withdrawModel.MerchantId),
				withdrawModel.TrxId,
				string(withdrawModel.Address),
				withdrawModel.Currency,
				withdrawModel.Amount,
				withdrawModel.AmountUsd,
				err,
				merchantModel.Settings.NotificationSettings.TelegramChatId,
			)
		}
	}
}
