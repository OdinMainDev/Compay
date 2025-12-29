package efwdt

import (
	"efw/utils"
	"fmt"
)

var (
	ErrorRequestNilAvailColumnOrder = fmt.Errorf("%s", "[EFW: DATATABLE REQUEST] NilAvailColumnOrder")
	ErrorRequestBadColumnOrder      = fmt.Errorf("%s", "[EFW: DATATABLE REQUEST] BadColumnOrder")
)

type Request[T any] struct {
	Start          int64  `json:"start" binding:"min=0"`
	Length         int    `json:"length" binding:"required,max=250"`
	OrderColumn    string `json:"order_column" binding:"required"`
	OrderDirection string `json:"order_direction" binding:"required,eq=desc|eq=asc"`
	Conditions     T      `json:"conditions" binding:"omitempty,dive"`
}

func (request Request[T]) ValidateOrderColumn(availColumns []string) error {
	if request.OrderColumn == "" {
		return nil
	}

	if len(availColumns) == 0 {
		return ErrorRequestNilAvailColumnOrder
	}

	if utils.InArray(request.OrderColumn, availColumns) {
		return nil
	}

	return ErrorRequestBadColumnOrder
}
