package user

import (
	"compay/src/entities/merchant"
	entity "compay/src/entities/user"
	"context"
	"efw"
	"efw/components/model"
	"efw/utils"
	"strings"

	"github.com/alexedwards/argon2id"
	"gorm.io/gorm"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
	NewPassword   string           `json:"-" gorm:"-"`
	Merchant      *merchant.Entity `json:"merchant,omitempty" gorm:"foreignKey:MerchantId;references:Id"`
}

var (
	GetModel = model.GetModel[*Model, QueryHelper]
	Save     = model.Save[*Model]
	Delete   = model.Delete[*Model]
	GetCount = model.GetCount[*Model, QueryHelper]
)

func ResetSequence() {
	efw.DBTransaction[any]().Exec(context.Background(), `SELECT setval(pg_get_serial_sequence('users', 'id'), (SELECT MAX(id) FROM users))`)
}

func (that *Model) BeforeSave(*gorm.DB) error {
	if that.Username != "" {
		that.Username = strings.ToLower(that.Username)
	}

	if that.SecretToken == "" {
		that.SecretToken = utils.GenerateRandomHash()
	}

	if that.NewPassword == "" {
		return nil
	}

	passwordHash, err := argon2id.CreateHash(that.NewPassword, argon2id.DefaultParams)
	if err != nil {
		return err
	}
	that.Password = string(passwordHash)
	return nil
}
