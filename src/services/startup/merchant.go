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
		Username: "test",
	})

	fmt.Println(userModel)
	fmt.Println(userModel == nil)
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
