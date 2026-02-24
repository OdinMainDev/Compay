package logmerchantwallet

import (
	entity "compay/src/entities/logmerchantwallet"

	"github.com/gin-gonic/gin"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
}

func DTO(model *Model) gin.H { return entity.DTO(&model.Entity) }
