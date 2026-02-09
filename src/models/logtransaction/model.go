package logtransaction

import (
	entity "compay/src/entities/logtransaction"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
}
