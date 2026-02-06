package merchantcurrency

import (
	entity "compay/src/entities/merchantcurrency"

	"github.com/gin-gonic/gin"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
}

func DTO(model *Model) gin.H { return entity.DTO(&model.Entity) }
