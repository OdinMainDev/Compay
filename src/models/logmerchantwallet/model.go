package logmerchantwallet

import (
	entity "compay/src/entities/logmerchantwallet"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
}
