package helpers

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/currency"
	"efw/core/efwnotify"
	"fmt"
)

// 🆕 - новая заявка
// ⚠️ - средства замечены
// ✅ - завершена
// ⛔️ - Не достаточно средств
// ❌ - удалена

// INVOICE NOTIFICATIONS
func NotificationCreatePayment(paymentId, merchantId int64, uuid, address string, codeAlpha types.CurrencyCodeAlpha, amount, amountUsd types.Amount, chatId int64) {
	invoiceType := "Фиксированная сумма"
	if amount == 0 {
		invoiceType = "Произвольная сумма"
	}
	text := "<b>🆕 НОВЫЙ ПЛАТЕЖНЫЙ ИНВОЙС 🆕</b>\n\n"
	text += "<b>Id:</b> <code>" + uuid + "</code>\n"
	if chatId == 0 {
		text += "<b>Id системный:</b> <code>" + fmt.Sprint(paymentId) + "</code>\n"
		text += "<b>Id мерчанта:</b> <code>" + fmt.Sprint(merchantId) + "</code>\n"
	}
	text += "<b>Тип инвойса:</b> <code>" + invoiceType + "</code>\n"
	text += "<b>Валюта:</b> <code>" + string(codeAlpha) + "</code>\n"
	text += "<b>Сумма:</b> <code>" + fmt.Sprint(currency.AmountToFloatByCodeAlpha(amount, codeAlpha)) + "</code>\n"
	text += "<b>Сумма в $:</b> <code>" + fmt.Sprint(currency.AmountToFloat(amountUsd, 2)) + "</code>\n"
	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = constants.NotificationChatLogs
	}

	efwnotify.NotifyTelegram(text, chatId)
}
func NotificationDetectedPayment(paymentId, merchantId int64, uuid, address string, codeAlpha types.CurrencyCodeAlpha, amount, amountUsd types.Amount, txId string, chatId int64) {
	invoiceType := "Фиксированная сумма"
	if amount == 0 {
		invoiceType = "Произвольная сумма"
	}
	text := "<b>⚠️ СРЕДСТВА ЗАМЕЧЕНЫ ⚠️</b>\n\n"
	text += "<b>Id:</b> <code>" + uuid + "</code>\n"
	if chatId == 0 {
		text += "<b>Id системный:</b> <code>" + fmt.Sprint(paymentId) + "</code>\n"
		text += "<b>Id мерчанта:</b> <code>" + fmt.Sprint(merchantId) + "</code>\n"
	}
	text += "<b>Tx Id:</b> <code>" + txId + "</code>\n"
	text += "<b>Тип инвойса:</b> <code>" + invoiceType + "</code>\n"
	text += "<b>Валюта:</b> <code>" + string(codeAlpha) + "</code>\n"
	text += "<b>Сумма:</b> <code>" + fmt.Sprint(currency.AmountToFloatByCodeAlpha(amount, codeAlpha)) + "</code>\n"
	text += "<b>Сумма в $:</b> <code>" + fmt.Sprint(currency.AmountToFloat(amountUsd, 2)) + "</code>\n"
	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = constants.NotificationChatLogs
	}
	efwnotify.NotifyTelegram(text, chatId)
}
func NotificationCompletedPayment(paymentId, merchantId int64, uuid, address string, codeAlpha types.CurrencyCodeAlpha, amount, amountUsd, receivedAmount, receivedAmountUsd types.Amount, chatId int64) {
	invoiceType := "Фиксированная сумма"
	if amount == 0 {
		invoiceType = "Произвольная сумма"
	}
	text := "<b>✅ Заявка завершена ✅</b>\n\n"
	text += "<b>Id:</b> <code>" + uuid + "</code>\n"
	if chatId == 0 {
		text += "<b>Id системный:</b> <code>" + fmt.Sprint(paymentId) + "</code>\n"
		text += "<b>Id мерчанта:</b> <code>" + fmt.Sprint(merchantId) + "</code>\n"
	}
	text += "<b>Тип инвойса:</b> <code>" + invoiceType + "</code>\n"
	text += "<b>Валюта:</b> <code>" + string(codeAlpha) + "</code>\n"
	text += "<b>Сумма:</b> <code>" + fmt.Sprint(currency.AmountToFloatByCodeAlpha(amount, codeAlpha)) + "</code>\n"
	text += "<b>Сумма в $:</b> <code>" + fmt.Sprint(currency.AmountToFloat(amountUsd, 2)) + "</code>\n"
	text += "<b>Полученная сумма:</b> <code>" + fmt.Sprint(currency.AmountToFloatByCodeAlpha(receivedAmount, codeAlpha)) + "</code>\n"
	text += "<b>Полученная сумма в $:</b> <code>" + fmt.Sprint(currency.AmountToFloat(receivedAmountUsd, 2)) + "</code>\n"
	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = constants.NotificationChatLogs
	}
	efwnotify.NotifyTelegram(text, chatId)
}
func NotificationFundsRecivedPayment(paymentId, merchantId int64, uuid, address string, codeAlpha types.CurrencyCodeAlpha, amount, amountUsd, receivedAmount, receivedAmountUsd types.Amount, txId string, chatId int64) {
	invoiceType := "Фиксированная сумма"
	if amount == 0 {
		invoiceType = "Произвольная сумма"
	}
	text := "<b>☑️☑️ Средства пришли на заявку ☑️☑️</b>\n\n"
	text += "<b>Id:</b> <code>" + uuid + "</code>\n"
	if chatId == 0 {
		text += "<b>Id системный:</b> <code>" + fmt.Sprint(paymentId) + "</code>\n"
		text += "<b>Id мерчанта:</b> <code>" + fmt.Sprint(merchantId) + "</code>\n"
	}
	text += "<b>Tx Id:</b> <code>" + txId + "</code>\n"
	text += "<b>Тип инвойса:</b> <code>" + invoiceType + "</code>\n"
	text += "<b>Валюта:</b> <code>" + string(codeAlpha) + "</code>\n"
	text += "<b>Сумма:</b> <code>" + fmt.Sprint(currency.AmountToFloatByCodeAlpha(amount, codeAlpha)) + "</code>\n"
	text += "<b>Сумма в $:</b> <code>" + fmt.Sprint(currency.AmountToFloat(amountUsd, 2)) + "</code>\n"
	text += "<b>Полученная сумма:</b> <code>" + fmt.Sprint(currency.AmountToFloatByCodeAlpha(receivedAmount, codeAlpha)) + "</code>\n"
	text += "<b>Полученная сумма в $:</b> <code>" + fmt.Sprint(currency.AmountToFloat(receivedAmountUsd, 2)) + "</code>\n"
	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = constants.NotificationChatLogs
	}
	efwnotify.NotifyTelegram(text, chatId)
}
func NotificationExpirePayment(paymentId, merchantId int64, uuid, address string, codeAlpha types.CurrencyCodeAlpha, amount, amountUsd types.Amount, chatId int64) {
	invoiceType := "Фиксированная сумма"
	if amount == 0 {
		invoiceType = "Произвольная сумма"
	}
	text := "<b>❌ ИНВОЙС УДАЛЕН ❌</b>\n\n"
	text += "<b>Id:</b> <code>" + uuid + "</code>\n"
	if chatId == 0 {
		text += "<b>Id системный:</b> <code>" + fmt.Sprint(paymentId) + "</code>\n"
		text += "<b>Id мерчанта:</b> <code>" + fmt.Sprint(merchantId) + "</code>\n"
	}
	text += "<b>Тип инвойса:</b> <code>" + invoiceType + "</code>\n"
	text += "<b>Валюта:</b> <code>" + string(codeAlpha) + "</code>\n"
	text += "<b>Сумма:</b> <code>" + fmt.Sprint(currency.AmountToFloatByCodeAlpha(amount, codeAlpha)) + "</code>\n"
	text += "<b>Сумма в $:</b> <code>" + fmt.Sprint(currency.AmountToFloat(amountUsd, 2)) + "</code>\n"
	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = constants.NotificationChatLogs
	}
	efwnotify.NotifyTelegram(text, chatId)
}
func NotificationInsufficientFundsPayment(paymentId, merchantId int64, uuid, address string, codeAlpha types.CurrencyCodeAlpha, amount, amountUsd, receivedAmount, receivedAmountUsd types.Amount, chatId int64) {
	invoiceType := "Фиксированная сумма"
	if amount == 0 {
		invoiceType = "Произвольная сумма"
	}
	text := "<b>⛔️ НЕ ДОСТАТОЧНО СРЕДСТВ В ИНВОЙСЕ ⛔️</b>\n\n"
	text += "<b>Id:</b> <code>" + uuid + "</code>\n"
	if chatId == 0 {
		text += "<b>Id системный:</b> <code>" + fmt.Sprint(paymentId) + "</code>\n"
		text += "<b>Id мерчанта:</b> <code>" + fmt.Sprint(merchantId) + "</code>\n"
	}
	text += "<b>Тип инвойса:</b> <code>" + invoiceType + "</code>\n"
	text += "<b>Валюта:</b> <code>" + string(codeAlpha) + "</code>\n"
	text += "<b>Сумма:</b> <code>" + fmt.Sprint(currency.AmountToFloatByCodeAlpha(amount, codeAlpha)) + "</code>\n"
	text += "<b>Сумма в $:</b> <code>" + fmt.Sprint(currency.AmountToFloat(amountUsd, 2)) + "</code>\n"
	text += "<b>Полученная сумма:</b> <code>" + fmt.Sprint(currency.AmountToFloatByCodeAlpha(receivedAmount, codeAlpha)) + "</code>\n"
	text += "<b>Полученная сумма в $:</b> <code>" + fmt.Sprint(currency.AmountToFloat(receivedAmountUsd, 2)) + "</code>\n"
	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = constants.NotificationChatLogs
	}
	efwnotify.NotifyTelegram(text, chatId)
}

// SYSTEM NOTIFICATIONS
func NotificationDebug(text string) {
	efwnotify.NotifyTelegram(text, constants.NotificationChatDebug)
}

func NotificationError(text string) {
	efwnotify.NotifyTelegram(text, constants.NotificationChatError)
}

func NotificationErrorProcesssingResponse(url, body string) {
	text := "<b>❗️❗️ ОШИБКА ПРИ ЗАПРОСЕ К НОДЕ ❗️❗️</b>\n\n"
	text += "<b>Ссылка:</b> " + url + "\n\n"
	text += "<pre>" + body + "</pre>"

	efwnotify.NotifyTelegram(text, constants.NotificationChatError)
}
