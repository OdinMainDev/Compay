package withdraw

import (
	entity "compay/src/entities/withdraw"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
}

// func DTO(model *Model) gin.H         { return entity.DTO(&model.Entity) }
// func DTOAdmin(model *Model) gin.H    { return entity.DTOAdmin(&model.Entity) }
// func DTOMerchant(model *Model) gin.H { return entity.DTOMerchant(&model.Entity) }
