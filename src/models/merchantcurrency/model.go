package merchantcurrency

import (
	entity "compay/src/entities/merchantcurrency"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
}
