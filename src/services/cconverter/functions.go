package cconverter

import (
	"compay/common/cperrors"
	"compay/src/helpers"
	"efw/components/entity"
	"efw/core/efwamount"
)

func Convert(amount entity.Amount, currencyFrom entity.CurrencyCodeAlpha) entity.Amount {
	resp, err := convert(amount, currencyFrom)
	if err != nil {
		helpers.NotificationError("[CConverter] Ошибка при конвертации: " + err.Error())
	}
	return resp
}

func convert(amount entity.Amount, currencyFrom entity.CurrencyCodeAlpha) (entity.Amount, error) {
	service.mutex.RLock()
	defer service.mutex.RUnlock()

	currencyRates, exists := service.rates[string(currencyFrom)]
	if !exists {
		return 0, cperrors.ErrorCConverterCurrencyNotExist(string(currencyFrom))
	}

	rate, exists := currencyRates["USDT"]
	if !exists {
		return 0, cperrors.ErrorCConverterUsdForCurrencyNotExist(string(currencyFrom))
	}

	return efwamount.FloatToAmount(
		efwamount.AmountToFloatByCodeAlpha(amount, currencyFrom)*rate,
		2,
	), nil
}
