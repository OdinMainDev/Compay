package currency

import (
	entity "compay/src/entities/currency"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
}
