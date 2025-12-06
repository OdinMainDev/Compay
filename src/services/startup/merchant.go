package startup

import (
	"compay/src/models/user"
	"fmt"
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
		fmt.Println("ВСЕ НАЙДЕНО ЗБС")
		return
	}

	userModel = &user.Model{}
	userModel.Username = "vioriel@developer.com"
	userModel.MerchantId = 0
	userModel.SetId(1)
	user.Save(userModel)
	// user.ResetSequence()
}
