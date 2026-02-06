package merchant

import (
	entity "compay/src/entities/merchant"

	"github.com/gin-gonic/gin"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
	// UserAdmin     *user.Entity `json:"user_admin,omitempty" gorm:"foreignKey:AdminId;references:Id"`
}

func DTO(model *Model) gin.H { return entity.DTO(&model.Entity) }
