package payment

import (
	entity "compay/src/entities/payment"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
}
