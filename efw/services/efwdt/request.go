package efwdt

import (
	"efw/utils"
	"fmt"
)

var (
	ErrorRequestBadColumnOrder = fmt.Errorf("%s", "[EFW: DATATABLE REQUEST] BadColumnOrder")
)

type Request[T any] struct {
	Start          int64  `json:"start" form:"start" binding:"min=0"`
	Length         int    `json:"length" form:"length" binding:"required,max=250"`
	OrderColumn    string `json:"order_column" form:"order_column" binding:"required"`
	OrderDirection string `json:"order_direction" form:"order_direction" binding:"required,eq=desc|eq=asc"`
	Conditions     T      `json:"conditions" form:"conditions" binding:"omitempty,dive"`
}

func (request Request[T]) ValidateOrderColumn(availColumns []string) error {
	if request.OrderColumn == "" {
		return nil
	}
	if len(availColumns) == 0 {
		return nil
	}
	if utils.InArray(request.OrderColumn, availColumns) {
		return nil
	}

	return ErrorRequestBadColumnOrder
}
