package logauth

import (
	entity "compay/src/entities/logauth"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
}
