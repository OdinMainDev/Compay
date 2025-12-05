package startup

import (
	"compay/src/entities/user"
	"fmt"
)

func StartMerchant() {
	registerMerchantAdminAccounts()
}

func registerMerchantAdminAccounts() {
	userModel := user.GetModel(user.QueryHelper{
		Username: "test1",
	})

	// fmt.Println(userModel)
	// fmt.Println(userModel == nil)

	userModel = &user.Entity{
		Username: "test",
	}

	fmt.Printf("Ссылка 0: %p\n", &userModel)
	user.Save(userModel)

	// "vioriel@developer.com"
	// check := user.GetModel(user.QueryHelper{
	// 	Id:                constants.DevIds[i],
	// 	DontCheckCasinoId: true,
	// })
	// if check != nil && check.GetId() != 0 {
	// 	continue
	// }
	// userModel := &user.Model{
	// 	Email:       email,
	// 	NewPassword: email,
	// }
	// userModel.Id = constants.DevIds[i]
	// userModel.SetCasinoId(0)
	// userModel.Save(userModel)
	// userRoleModel := &rbac.UserRole{
	// 	UserId: userModel.GetId(),
	// 	RoleId: roleModel.GetId(),
	// }
	// userRoleModel.Save(userRoleModel)
	// user.ResetSequence()
}
