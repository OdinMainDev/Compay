package importdata

import (
	"compay/common/constants"
	"compay/src/models/currency"
	"time"
)

var isImportCurrenciesWorking = false

func ImportCurrencies() bool {
	if isImportCurrenciesWorking {
		return false
	}
	isImportCurrenciesWorking = true
	go func() {
		currencies := []*currency.Model{}
		currencies = append(currencies, addBtcCurrency())

		for _, item := range currencies {
			currencyModel := currency.GetModel(currency.QueryHelper{CodeAlpha: item.CodeAlpha})
			if currencyModel == nil {
				currencyModel = &currency.Model{}
			}

			currencyModel.CodeAlpha = item.CodeAlpha
			currencyModel.Title = item.Title
			currencyModel.Network = item.Network
			currencyModel.IsCoin = item.IsCoin
			currencyModel.Image = item.Image
			currencyModel.MinorUnits = item.MinorUnits

			currency.Save(currencyModel, currency.QueryHelper{})
		}
		time.Sleep(30 * time.Second)
		isImportCurrenciesWorking = false
	}()

	return true
}

func addBtcCurrency() *currency.Model {
	resp := &currency.Model{}
	resp.CodeAlpha = "BTC"
	resp.Title = "Bitcoin"
	resp.Network = constants.CurrencyNetworkBtc
	resp.IsCoin = false
	resp.MinorUnits = 8

	return resp
}
