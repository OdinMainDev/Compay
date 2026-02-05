package startup

import (
	"compay/src/models/blockchainstate"
	"compay/src/models/currency"
	"compay/src/models/merchant"
	"compay/src/models/merchantcurrency"
	"compay/src/models/user"
	"efw/core/efwcache"
)

func StartMerchant() {
	registerMerchantAdminAccounts()
	registerMerchantCacheTypes()
}

func registerMerchantAdminAccounts() {
	userModel := user.GetModel(user.QueryHelper{
		Id:                  1,
		DontCheckMerchantId: true,
	})

	if userModel != nil {
		if userModel.MerchantId != 0 {
			panic("BAD ROOT ADMIN MODEL")
		}
		return
	}

	userModel = &user.Model{}
	userModel.Username = "vioriel@developer.com"
	userModel.NewPassword = "vioriel@developer.com"
	userModel.SetId(1)
	user.Save(userModel, user.QueryHelper{})
	user.ResetSequence()
}

func registerMerchantCacheTypes() {
	efwcache.RegisterType(&currency.Model{})
	efwcache.RegisterType(&merchant.Model{})
	efwcache.RegisterType(&blockchainstate.Model{})
	efwcache.RegisterType(&merchantcurrency.Model{})
}
