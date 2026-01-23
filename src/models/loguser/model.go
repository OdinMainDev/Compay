package loguser

import (
	entity "compay/src/entities/loguser"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
}
