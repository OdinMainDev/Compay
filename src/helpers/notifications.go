package helpers

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/currency"
	"efw/core/efwnotify"
	"fmt"
	"strconv"
)

// INVOICE NOTIFICATIONS
func NotificationCreateInvoice(paymentId, merchantId int64, uuid, address string, codeAlpha types.CurrencyCodeAlpha, amount, amountUsd types.Amount, chatId int64) {
	invoiceType := "Фиксированная сумма"
	if amount == 0 {
		invoiceType = "Произвольная сумма"
	}
	text := "<b>🆕 Новый платежный инвойс 🆕</b>\n\n"
	text += "<b>Id:</b> <code>" + uuid + "</code>\n"
	if chatId == 0 {
		text += "<b>Id системный:</b> <code>" + fmt.Sprint(paymentId) + "</code>\n"
		text += "<b>Id мерчанта:</b> <code>" + fmt.Sprint(merchantId) + "</code>\n"
	}
	text += "<b>Тип инвойса:</b> <code>" + invoiceType + "</code>\n"
	text += "<b>Валюта:</b> <code>" + string(codeAlpha) + "</code>\n"
	text += "<b>Сумма:</b> <code>" + strconv.FormatFloat(currency.AmountToFloatByCodeAlpha(amount, codeAlpha), 'f', -1, 64) + "</code>\n"
	text += "<b>Сумма в $:</b> <code>" + strconv.FormatFloat(currency.AmountToFloat(amountUsd, 2), 'f', -1, 64) + "</code>\n"
	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = constants.NotificationChatLogs
	}

	efwnotify.NotifyTelegram(text, chatId)
}
func NotificationDetectedTransaction(paymentId, merchantId int64, uuid, address string, codeAlpha types.CurrencyCodeAlpha, amount, amountUsd types.Amount, txId string, chatId int64) {
	invoiceType := "Фиксированная сумма"
	if amount == 0 {
		invoiceType = "Произвольная сумма"
	}
	text := "<b>⚠️ Средства замечены ⚠️</b>\n\n"
	text += "<b>Id:</b> <code>" + uuid + "</code>\n"
	if chatId == 0 {
		text += "<b>Id системный:</b> <code>" + fmt.Sprint(paymentId) + "</code>\n"
		text += "<b>Id мерчанта:</b> <code>" + fmt.Sprint(merchantId) + "</code>\n"
	}
	text += "<b>Tx Id:</b> <code>" + txId + "</code>\n"
	text += "<b>Тип инвойса:</b> <code>" + invoiceType + "</code>\n"
	text += "<b>Валюта:</b> <code>" + string(codeAlpha) + "</code>\n"
	text += "<b>Сумма:</b> <code>" + strconv.FormatFloat(currency.AmountToFloatByCodeAlpha(amount, codeAlpha), 'f', -1, 64) + "</code>\n"
	text += "<b>Сумма в $:</b> <code>" + strconv.FormatFloat(currency.AmountToFloat(amountUsd, 2), 'f', -1, 64) + "</code>\n"
	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = constants.NotificationChatLogs
	}
	efwnotify.NotifyTelegram(text, chatId)
}
func NotificationFundsRecieved(paymentId, merchantId int64, uuid, address string, codeAlpha types.CurrencyCodeAlpha, amount, amountUsd, receivedAmount types.Amount, txId string, chatId int64) {
	invoiceType := "Фиксированная сумма"
	if amount == 0 {
		invoiceType = "Произвольная сумма"
	}
	text := "<b>☑️☑️ Средства пришли на инвойс ☑️☑️</b>\n\n"
	text += "<b>Id:</b> <code>" + uuid + "</code>\n"
	if chatId == 0 {
		text += "<b>Id системный:</b> <code>" + fmt.Sprint(paymentId) + "</code>\n"
		text += "<b>Id мерчанта:</b> <code>" + fmt.Sprint(merchantId) + "</code>\n"
	}
	text += "<b>Tx Id:</b> <code>" + txId + "</code>\n"
	text += "<b>Тип инвойса:</b> <code>" + invoiceType + "</code>\n"
	text += "<b>Валюта:</b> <code>" + string(codeAlpha) + "</code>\n"
	text += "<b>Сумма:</b> <code>" + strconv.FormatFloat(currency.AmountToFloatByCodeAlpha(amount, codeAlpha), 'f', -1, 64) + "</code>\n"
	text += "<b>Сумма в $:</b> <code>" + strconv.FormatFloat(currency.AmountToFloat(amountUsd, 2), 'f', -1, 64) + "</code>\n"
	text += "<b>Полученная сумма:</b> <code>" + strconv.FormatFloat(currency.AmountToFloatByCodeAlpha(receivedAmount, codeAlpha), 'f', -1, 64) + "</code>\n"
	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = constants.NotificationChatLogs
	}
	efwnotify.NotifyTelegram(text, chatId)
}
func NotificationCompletedInvoice(paymentId, merchantId int64, uuid, address string, codeAlpha types.CurrencyCodeAlpha, amount, amountUsd, receivedAmount, receivedAmountUsd types.Amount, chatId int64) {
	invoiceType := "Фиксированная сумма"
	if amount == 0 {
		invoiceType = "Произвольная сумма"
	}
	text := "<b>✅ Инвойс завершен ✅</b>\n\n"
	text += "<b>Id:</b> <code>" + uuid + "</code>\n"
	if chatId == 0 {
		text += "<b>Id системный:</b> <code>" + fmt.Sprint(paymentId) + "</code>\n"
		text += "<b>Id мерчанта:</b> <code>" + fmt.Sprint(merchantId) + "</code>\n"
	}
	text += "<b>Тип инвойса:</b> <code>" + invoiceType + "</code>\n"
	text += "<b>Валюта:</b> <code>" + string(codeAlpha) + "</code>\n"
	text += "<b>Сумма:</b> <code>" + strconv.FormatFloat(currency.AmountToFloatByCodeAlpha(amount, codeAlpha), 'f', -1, 64) + "</code>\n"
	text += "<b>Сумма в $:</b> <code>" + strconv.FormatFloat(currency.AmountToFloat(amountUsd, 2), 'f', -1, 64) + "</code>\n"
	text += "<b>Полученная сумма:</b> <code>" + strconv.FormatFloat(currency.AmountToFloatByCodeAlpha(receivedAmount, codeAlpha), 'f', -1, 64) + "</code>\n"
	text += "<b>Полученная сумма в $:</b> <code>" + strconv.FormatFloat(currency.AmountToFloat(receivedAmountUsd, 2), 'f', -1, 64) + "</code>\n"
	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = constants.NotificationChatLogs
	}
	efwnotify.NotifyTelegram(text, chatId)
}
func NotificationExpireInvoice(paymentId, merchantId int64, uuid, address string, codeAlpha types.CurrencyCodeAlpha, amount, amountUsd types.Amount, chatId int64) {
	invoiceType := "Фиксированная сумма"
	if amount == 0 {
		invoiceType = "Произвольная сумма"
	}
	text := "<b>❌ Инвойс удален ❌</b>\n\n"
	text += "<b>Id:</b> <code>" + uuid + "</code>\n"
	if chatId == 0 {
		text += "<b>Id системный:</b> <code>" + fmt.Sprint(paymentId) + "</code>\n"
		text += "<b>Id мерчанта:</b> <code>" + fmt.Sprint(merchantId) + "</code>\n"
	}
	text += "<b>Тип инвойса:</b> <code>" + invoiceType + "</code>\n"
	text += "<b>Валюта:</b> <code>" + string(codeAlpha) + "</code>\n"
	text += "<b>Сумма:</b> <code>" + strconv.FormatFloat(currency.AmountToFloatByCodeAlpha(amount, codeAlpha), 'f', -1, 64) + "</code>\n"
	text += "<b>Сумма в $:</b> <code>" + strconv.FormatFloat(currency.AmountToFloat(amountUsd, 2), 'f', -1, 64) + "</code>\n"
	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = constants.NotificationChatLogs
	}
	efwnotify.NotifyTelegram(text, chatId)
}
func NotificationInsufficientFunds(paymentId, merchantId int64, uuid, address string, codeAlpha types.CurrencyCodeAlpha, amount, amountUsd, receivedAmount, receivedAmountUsd types.Amount, chatId int64) {
	invoiceType := "Фиксированная сумма"
	if amount == 0 {
		invoiceType = "Произвольная сумма"
	}
	text := "<b>⛔️ Не достаточно средств в инвойсе (завершен) ⛔️</b>\n\n"
	text += "<b>Id:</b> <code>" + uuid + "</code>\n"
	if chatId == 0 {
		text += "<b>Id системный:</b> <code>" + fmt.Sprint(paymentId) + "</code>\n"
		text += "<b>Id мерчанта:</b> <code>" + fmt.Sprint(merchantId) + "</code>\n"
	}
	text += "<b>Тип инвойса:</b> <code>" + invoiceType + "</code>\n"
	text += "<b>Валюта:</b> <code>" + string(codeAlpha) + "</code>\n"
	text += "<b>Сумма:</b> <code>" + strconv.FormatFloat(currency.AmountToFloatByCodeAlpha(amount, codeAlpha), 'f', -1, 64) + "</code>\n"
	text += "<b>Сумма в $:</b> <code>" + strconv.FormatFloat(currency.AmountToFloat(amountUsd, 2), 'f', -1, 64) + "</code>\n"
	text += "<b>Полученная сумма:</b> <code>" + strconv.FormatFloat(currency.AmountToFloatByCodeAlpha(receivedAmount, codeAlpha), 'f', -1, 64) + "</code>\n"
	text += "<b>Полученная сумма в $:</b> <code>" + strconv.FormatFloat(currency.AmountToFloat(receivedAmountUsd, 2), 'f', -1, 64) + "</code>\n"
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
