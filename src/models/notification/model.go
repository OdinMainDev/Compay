package notification

import (
	entity "compay/src/entities/notification"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
}
