package merchant

import (
	entity "compay/src/entities/merchant"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
	// UserAdmin     *user.Entity `json:"user_admin,omitempty" gorm:"foreignKey:AdminId;references:Id"`
}
