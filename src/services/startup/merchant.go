package startup

import (
	"compay/common/constants"
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
	for i, id := range constants.AdminIds {
		userModel := user.GetModel(user.QueryHelper{
			Id:                  id,
			DontCheckMerchantId: true,
		})

		if userModel != nil {
			if userModel.MerchantId != 0 {
				panic("BAD ROOT ADMIN MODEL")
			}
			return
		}
		userModel = &user.Model{}
		userModel.Username = constants.AdminUsernames[i]
		userModel.NewPassword = constants.AdminUsernames[i]
		userModel.SetId(id)
		user.Save(userModel, user.QueryHelper{})
	}
	user.ResetSequence()
}

func registerMerchantCacheTypes() {
	efwcache.RegisterType(&currency.Model{})
	efwcache.RegisterType(&merchant.Model{})
	efwcache.RegisterType(&blockchainstate.Model{})
	efwcache.RegisterType(&merchantcurrency.Model{})
}
