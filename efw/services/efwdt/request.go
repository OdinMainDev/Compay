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
	OrderColumn    string `json:"order_column" binding:"required,min=1"`
	OrderDirection string `json:"order_direction" binding:"required,eq=desc|eq=asc"`
	Filters        T      `json:"filters"`
}

type WithCreatedAt struct {
	CreatedAtFrom int64 `json:"created_at_from" efwdt:"created_at_from" binding:"min=0"`
	CreatedAtTo   int64 `json:"created_at_to" efwdt:"created_at_to" binding:"min=0"`
}

type WithUpdatedAt struct {
	UpdatedAtFrom int64 `json:"updated_at_from" efwdt:"updated_at_from" binding:"min=0"`
	UpdatedAtTo   int64 `json:"updated_at_to" efwdt:"updated_at_to" binding:"min=0"`
}

type WithExpiresAt struct {
	ExpiresAtFrom int64 `json:"expires_at_from" efwdt:"expires_at_from" binding:"min=0"`
	ExpiresAtTo   int64 `json:"expires_at_to" efwdt:"expires_at_to" binding:"min=0"`
}

func (request *Request[T]) ValidateOrderColumn(availColumns []string) error {
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
