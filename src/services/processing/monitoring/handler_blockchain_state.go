package monitoring

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/helpers"
	"compay/src/models/blockchainstate"
	"compay/src/models/currency"
	"efw/core/efwlog"
	"fmt"
)

func (that *monitoring) loadBlockchainStates() {
	currencyModels := currency.GetModels(currency.QueryHelper{})
	for _, currencyModel := range currencyModels {
		blockchainStateModel := blockchainstate.GetModel(blockchainstate.QueryHelper{
			Currency: currencyModel.CodeAlpha,
		})

		if blockchainStateModel == nil {
			blockchainStateModel = that.createDefaultMerchantState(currencyModel.CodeAlpha, 0, nil)
		}

		switch currencyModel.CodeAlpha {
		case constants.BtcCodeAlpha:
			that.btc.blockchainStateModel = blockchainStateModel

		default:
			helpers.NotificationError(fmt.Sprint("BAD CONFIGURATION loadBlockchainStates: ", currencyModel.CodeAlpha))
			efwlog.Panic("BAD CONFIGURATION loadBlockchainStates: ", currencyModel.CodeAlpha)
		}
	}
}

func (that *monitoring) createDefaultMerchantState(codeAlpha types.CurrencyCodeAlpha, attempt int, err error) *blockchainstate.Model {
	resp := &blockchainstate.Model{}
	resp.Currency = codeAlpha

	if attempt >= 5 {
		helpers.NotificationError(fmt.Sprint("FATAL ERROR ON createDefaultMerchantState: ", codeAlpha, " ", err.Error()))
		efwlog.Panic("FATAL ERROR ON createDefaultMerchantState: ", codeAlpha, " ", err.Error())
	}

	switch codeAlpha {
	case constants.BtcCodeAlpha:
		var response getBlockCountBtcResponse
		err := that.doRequest(requestGetBlockCount, codeAlpha, nil, &response)
		if err != nil {
			return that.createDefaultMerchantState(codeAlpha, attempt+1, err)
		}
		resp.CurrentBlock = response.Data

	default:
		helpers.NotificationError(fmt.Sprint("BAD CONFIGURATION createDefaultMerchantState: ", codeAlpha))
		efwlog.Panic("BAD CONFIGURATION createDefaultMerchantState: ", codeAlpha)
	}

	blockchainstate.Save(resp, blockchainstate.QueryHelper{})
	return resp
}
