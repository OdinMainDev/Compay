package merchantuser

import (
	"compay/application/entities/basemodel"
	"compay/core/utils"
	"fmt"
	"reflect"
	"runtime"

	"gorm.io/gorm"
)

type UserModel struct {
	basemodel.BaseModelWithMerchantId
	TgId           int64                     `json:"tg_id" gorm:"column:tg_id"`
	FirstName      string                    `json:"first_name" gorm:"column:first_name"`
	LastName       string                    `json:"last_name" gorm:"column:last_name"`
	Username       string                    `json:"username" gorm:"column:username"`
	IsPremium      bool                      `json:"is_premium" gorm:"column:is_premium"`
	IsActive       bool                      `json:"is_active" gorm:"column:is_active"`
	IsApproved     bool                      `json:"is_approved" gorm:"column:is_approved"`
	SecretToken    basemodel.EncryptedString `json:"secret_token" gorm:"column:secret_token"`
	TotpToken      basemodel.EncryptedString `json:"totp_token" gorm:"column:totp_token"`
	IsTwofaEnabled bool                      `json:"is_twofa_enabled" gorm:"column:is_twofa_enabled"`
}

func (UserModel) TableName() string {
	return "users"
}

func (user *UserModel) BeforeSave(trx *gorm.DB) error {
	if user.SecretToken == "" {
		user.SecretToken = basemodel.EncryptedString(utils.GenerateRandomSalt())
	}

	return nil
}

func init() {
	// получаем тип указателя на модель
	t := reflect.TypeOf(&UserModel{})

	fmt.Printf("Inspecting methods of %s (pkg %q)\n", t.Elem().Name(), t.Elem().PkgPath())
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fn := runtime.FuncForPC(m.Func.Pointer())
		fmt.Printf(
			"  reflect.Name = %-16q  runtime.Name = %q\n",
			m.Name,
			fn.Name(),
		)
	}
}
