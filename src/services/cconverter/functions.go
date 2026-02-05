package cconverter

import (
	"compay/common/cperrors"
	"compay/common/types"
	"compay/src/helpers"
	"compay/src/models/currency"
)

func Convert(amount types.Amount, currencyFrom types.CurrencyCodeAlpha) types.Amount {
	resp, err := convert(amount, currencyFrom)
	if err != nil {
		helpers.NotificationError("[CConverter] Ошибка при конвертации: " + err.Error())
	}
	return resp
}

func convert(amount types.Amount, currencyFrom types.CurrencyCodeAlpha) (types.Amount, error) {
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

	return currency.FloatToAmount(
		currency.AmountToFloatByCodeAlpha(amount, currencyFrom)*rate,
		2,
	), nil
}
