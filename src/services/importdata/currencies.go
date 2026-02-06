package importdata

import (
	"compay/common/constants"
	"compay/src/models/currency"
	"efw/core/efwamount"
	"sync"
	"time"
)

var isImportCurrenciesWorking = false
var mutex = sync.Mutex{}

func ImportCurrencies() bool {
	if isImportCurrenciesWorking {
		return true
	}
	mutex.Lock()
	isImportCurrenciesWorking = true
	mutex.Unlock()
	go func() {
		currencies := []*currency.Model{}
		currencies = append(currencies, addBtcCurrency())

		for _, item := range currencies {
			currencyModel := currency.GetModel(currency.QueryHelper{CodeAlpha: item.CodeAlpha})
			if currencyModel == nil {
				currencyModel = &currency.Model{}
			}

			currencyModel.CodeAlpha = item.CodeAlpha
			currencyModel.IsEnabled = true
			currencyModel.Title = item.Title
			currencyModel.Network = item.Network
			currencyModel.IsCoin = item.IsCoin
			currencyModel.Image = item.Image
			currencyModel.MinorUnits = item.MinorUnits

			currency.Save(currencyModel, currency.QueryHelper{})
			efwamount.SetMinorUnits(currencyModel.CodeAlpha, currencyModel.MinorUnits)
		}
		time.Sleep(30 * time.Second)
		mutex.Lock()
		isImportCurrenciesWorking = false
		mutex.Unlock()
	}()

	return false
}

func addBtcCurrency() *currency.Model {
	resp := &currency.Model{}
	resp.CodeAlpha = constants.BtcCodeAlpha
	resp.Title = "Bitcoin"
	resp.Network = constants.CurrencyNetworkBtc
	resp.IsCoin = false
	resp.MinorUnits = 8

	return resp
}
