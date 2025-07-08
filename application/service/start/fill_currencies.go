package start

import "compay/application/entities/currency"

func FillCurrencies() {
	currencies := []*currency.Model{
		{
			CodeAlpha:  "BTC",
			Title:      "Bitcoin",
			Network:    "BTC",
			Image:      "btc",
			IsCoin:     true,
			MinorUnits: 8,
		},
		{
			CodeAlpha:  "ETH",
			Title:      "Ethereum",
			Network:    "Ethereum",
			Image:      "eth",
			IsCoin:     true,
			MinorUnits: 18,
		},
		{
			CodeAlpha:  "TRX",
			Title:      "Tron",
			Network:    "Tron",
			Image:      "trx",
			IsCoin:     true,
			MinorUnits: 6,
		},
		{
			CodeAlpha:  "SOL",
			Title:      "Solana",
			Network:    "Solana",
			Image:      "sol",
			IsCoin:     true,
			MinorUnits: 9,
		},
		{
			CodeAlpha:  "LTC",
			Title:      "Litecoin",
			Network:    "Litecoin",
			Image:      "ltc",
			IsCoin:     true,
			MinorUnits: 8,
		},
		{
			CodeAlpha:  "USDTTRC20",
			Title:      "USDT",
			Network:    "Tron",
			Image:      "ltc",
			IsCoin:     false,
			MinorUnits: 6,
		},
	}

	for _, item := range currencies {
		currencyModel := currency.GetModel(currency.QueryHelper{
			CodeAlpha: item.CodeAlpha,
		})

		if currencyModel != nil && currencyModel.GetId() != 0 {
			continue
		}

		currencyModel = &currency.Model{
			CodeAlpha:  item.CodeAlpha,
			Title:      item.Title,
			Network:    item.Network,
			Image:      item.Image,
			IsCoin:     item.IsCoin,
			MinorUnits: item.MinorUnits,
		}

		currencyModel.Save(currencyModel)
	}
}
