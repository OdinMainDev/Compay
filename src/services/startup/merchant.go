package startup

import (
	"compay/src/models/user"
	"fmt"
)

func StartMerchant() {
	registerMerchantAdminAccounts()
}

func registerMerchantAdminAccounts() {
	userModel2 := &user.Model{}
	fmt.Println("GetSaveType: ", userModel2.GetSaveType())

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
	fmt.Println("GetSaveType: ", userModel.GetSaveType())
	userModel.Username = "vioriel@developer.com"
	userModel.MerchantId = 0
	userModel.SetId(1)
	user.Save(userModel)

	// fmt.Println(userModel.GetSaveType())
	user.ResetSequence()
}
