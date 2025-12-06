package startup

import (
	"compay/src/entities/user"
)

func StartMerchant() {
	registerMerchantAdminAccounts()
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

	userModel = &user.Entity{
		Username:   "vioriel@developer.com",
		MerchantId: 0,
	}

	user.Save(userModel)
	// user.ResetSequence()
}
