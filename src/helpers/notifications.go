package helpers

import (
	"compay/common/config"
	"efw/components/entity"
	"efw/core/efwamount"
	"efw/core/efwnotify"
	"fmt"
	"strconv"
)

// INVOICE NOTIFICATIONS
func NotificationCreateInvoice(paymentId, merchantId int64, uuid, address string, codeAlpha entity.CurrencyCodeAlpha, amount, amountUsd entity.Amount, chatId int64) {
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
	if amount != 0 {
		amountFormatted := strconv.FormatFloat(efwamount.ToFloatByCodeAlpha(amount, codeAlpha), 'f', -1, 64)
		amountUsdFormatted := strconv.FormatFloat(efwamount.ToFloat(amountUsd, 2), 'f', -1, 64)
		text += "<b>Сумма:</b> <code>" + amountFormatted
		text += " (" + amountUsdFormatted + " $)</code>\n"
	}
	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = config.NotificationChatLogs
	}

	fmt.Println(efwnotify.NotifyTelegram(text, chatId))
}
func NotificationDetectedTransaction(paymentId, merchantId int64, uuid, address string, codeAlpha entity.CurrencyCodeAlpha, amount, amountUsd entity.Amount, txId string, chatId int64) {
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
	if amount != 0 {
		amountFormatted := strconv.FormatFloat(efwamount.ToFloatByCodeAlpha(amount, codeAlpha), 'f', -1, 64)
		amountUsdFormatted := strconv.FormatFloat(efwamount.ToFloat(amountUsd, 2), 'f', -1, 64)
		text += "<b>Сумма:</b> <code>" + amountFormatted
		text += " (" + amountUsdFormatted + " $)</code>\n"
	}
	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = config.NotificationChatLogs
	}
	efwnotify.NotifyTelegram(text, chatId)
}
func NotificationFundsReceived(paymentId, merchantId int64, uuid, address string, codeAlpha entity.CurrencyCodeAlpha, amount, amountUsd, receivedAmount, receivedAmountUsd entity.Amount, txId string, chatId int64) {
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

	if amount != 0 {
		amountFormatted := strconv.FormatFloat(efwamount.ToFloatByCodeAlpha(amount, codeAlpha), 'f', -1, 64)
		amountUsdFormatted := strconv.FormatFloat(efwamount.ToFloat(amountUsd, 2), 'f', -1, 64)
		text += "<b>Сумма:</b> <code>" + amountFormatted
		text += " (" + amountUsdFormatted + " $)</code>\n"
	}
	receivedAmountFormatted := strconv.FormatFloat(efwamount.ToFloatByCodeAlpha(receivedAmount, codeAlpha), 'f', -1, 64)
	receivedAmountUsdFormatted := strconv.FormatFloat(efwamount.ToFloat(receivedAmountUsd, 2), 'f', -1, 64)
	text += "<b>Полученная сумма:</b> <code>" + receivedAmountFormatted
	text += " (" + receivedAmountUsdFormatted + " $)</code>\n"

	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = config.NotificationChatLogs
	}
	efwnotify.NotifyTelegram(text, chatId)
}
func NotificationCompletedInvoice(paymentId, merchantId int64, uuid, address string, codeAlpha entity.CurrencyCodeAlpha, amount, amountUsd, receivedAmount, receivedAmountUsd entity.Amount, chatId int64) {
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

	if amount != 0 {
		amountFormatted := strconv.FormatFloat(efwamount.ToFloatByCodeAlpha(amount, codeAlpha), 'f', -1, 64)
		amountUsdFormatted := strconv.FormatFloat(efwamount.ToFloat(amountUsd, 2), 'f', -1, 64)
		text += "<b>Сумма:</b> <code>" + amountFormatted
		text += " (" + amountUsdFormatted + " $)</code>\n"
	}
	receivedAmountFormatted := strconv.FormatFloat(efwamount.ToFloatByCodeAlpha(receivedAmount, codeAlpha), 'f', -1, 64)
	receivedAmountUsdFormatted := strconv.FormatFloat(efwamount.ToFloat(receivedAmountUsd, 2), 'f', -1, 64)
	text += "<b>Полученная сумма:</b> <code>" + receivedAmountFormatted
	text += " (" + receivedAmountUsdFormatted + " $)</code>\n"

	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = config.NotificationChatLogs
	}
	efwnotify.NotifyTelegram(text, chatId)
}
func NotificationExpireInvoice(paymentId, merchantId int64, uuid, address string, codeAlpha entity.CurrencyCodeAlpha, amount, amountUsd, receivedAmount, receivedAmountUsd entity.Amount, chatId int64) {
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

	if amount != 0 {
		amountFormatted := strconv.FormatFloat(efwamount.ToFloatByCodeAlpha(amount, codeAlpha), 'f', -1, 64)
		amountUsdFormatted := strconv.FormatFloat(efwamount.ToFloat(amountUsd, 2), 'f', -1, 64)
		text += "<b>Сумма:</b> <code>" + amountFormatted
		text += " (" + amountUsdFormatted + " $)</code>\n"
	}
	if receivedAmount != 0 {
		receivedAmountFormatted := strconv.FormatFloat(efwamount.ToFloatByCodeAlpha(receivedAmount, codeAlpha), 'f', -1, 64)
		receivedAmountUsdFormatted := strconv.FormatFloat(efwamount.ToFloat(receivedAmountUsd, 2), 'f', -1, 64)
		text += "<b>Полученная сумма:</b> <code>" + receivedAmountFormatted
		text += " (" + receivedAmountUsdFormatted + " $)</code>\n"
	}

	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = config.NotificationChatLogs
	}
	efwnotify.NotifyTelegram(text, chatId)
}
func NotificationInsufficientFunds(paymentId, merchantId int64, uuid, address string, codeAlpha entity.CurrencyCodeAlpha, amount, amountUsd, receivedAmount, receivedAmountUsd entity.Amount, chatId int64) {
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
	text += "<b>Сумма:</b> <code>" + strconv.FormatFloat(efwamount.ToFloatByCodeAlpha(amount, codeAlpha), 'f', -1, 64) + "</code>\n"
	text += "<b>Сумма в $:</b> <code>" + strconv.FormatFloat(efwamount.ToFloat(amountUsd, 2), 'f', -1, 64) + "</code>\n"
	text += "<b>Полученная сумма:</b> <code>" + strconv.FormatFloat(efwamount.ToFloatByCodeAlpha(receivedAmount, codeAlpha), 'f', -1, 64) + "</code>\n"
	text += "<b>Полученная сумма в $:</b> <code>" + strconv.FormatFloat(efwamount.ToFloat(receivedAmountUsd, 2), 'f', -1, 64) + "</code>\n"
	text += "<b>Адрес поступления средств:</b> <code>" + address + "</code>\n"

	if chatId == 0 {
		chatId = config.NotificationChatLogs
	}
	efwnotify.NotifyTelegram(text, chatId)
}

// SYSTEM NOTIFICATIONS
func NotificationDebug(text string) {
	efwnotify.NotifyTelegram(text, config.NotificationChatDebug)
}

func NotificationError(text string) {
	efwnotify.NotifyTelegram(text, config.NotificationChatError)
}

func NotificationErrorProcesssingResponse(url, body string) {
	text := "<b>❗️❗️ ОШИБКА ПРИ ЗАПРОСЕ К НОДЕ ❗️❗️</b>\n\n"
	text += "<b>Ссылка:</b> " + url + "\n\n"
	text += "<pre>" + body + "</pre>"

	efwnotify.NotifyTelegram(text, config.NotificationChatError)
}
