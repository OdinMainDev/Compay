package helpers

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/currency"
	"efw/core/efwnotify"
	"fmt"
)

// INVOICE NOTIFICATIONS
func NotificationCreatePayment(paymentId, merchantId int64, trxId, address string, codeAlpha types.CurrencyCodeAlpha, amount, amountUsd types.Amount) {
	invoiceType := "Фиксированная сумма"
	if amount == 0 {
		invoiceType = "Произвольная сумма"
	}
	text := "<b>🟡 НОВЫЙ ПЛАТЕЖНЫЙ ИНВОЙС 🟡</b>\n\n"
	text += "<b>Id:</b> " + fmt.Sprint(paymentId) + "\n"
	text += "<b>Id внутренний:</b> " + trxId + "\n"
	text += "<b>Id мерчанта:</b> " + fmt.Sprint(merchantId) + "\n"
	text += "<b>Тип инвойса:</b> " + invoiceType + "\n"
	text += "<b>Валюта:</b> " + string(codeAlpha) + "\n"
	text += "<b>Сумма:</b> " + fmt.Sprint(currency.AmountToFloatByCodeAlpha(amount, codeAlpha)) + "\n"
	text += "<b>Сумма в $:</b> " + fmt.Sprint(amountUsd) + "\n"
	text += "<b>Адрес поступления средств:</b> " + address + "\n"

	efwnotify.NotifyTelegram(text, constants.NotificationChatLogs)
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
