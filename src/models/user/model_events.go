package user

import (
	"efw/utils"

	"github.com/alexedwards/argon2id"
	"gorm.io/gorm"
)

func (that *Model) BeforeSave(*gorm.DB) error {
	that.Username = utils.ToLower(that.Username)

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
