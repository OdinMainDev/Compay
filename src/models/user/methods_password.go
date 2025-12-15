package user

import (
	"compay/common/cperrors"

	"github.com/alexedwards/argon2id"
)

func (that *Model) ValidatePassword(password string) error {
	if ok, err := argon2id.ComparePasswordAndHash(password, that.Password); err != nil || !ok {
		if err != nil {
			return err
		}
		if !ok {
			return cperrors.ErrorAuthPasswordDoNotMatch
		}
	}

	return nil
}
