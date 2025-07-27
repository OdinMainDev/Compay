package datatable

import (
	"compay/core/utils"
	"errors"
)

type FilterType string

const (
	FilterTypeNumeric  FilterType = "numeric"
	FilterTypeBool     FilterType = "bool"
	FilterTypeText     FilterType = "text"
	FilterTypeList     FilterType = "list"
	FilterTypeUserId   FilterType = "user_id"
	FilterTypeDate     FilterType = "date"
	FilterTypeCurrency FilterType = "currency"
)

type ResponseFilter struct {
	Name   string     `json:"name"`
	Type   FilterType `json:"type"`
	Values any        `json:"values"`
}

type RequestDatatable struct {
	Start                  int64    `json:"start" form:"start" binding:"min=0"`
	Length                 int      `json:"length" form:"length" binding:"required,max=250"`
	OrderColumn            string   `json:"order_column" form:"order_column" binding:"required"`
	OrderDirection         string   `json:"order_direction" form:"order_direction" binding:"required,eq=desc|eq=asc"`
	FilterName             []string `json:"filter_name" form:"filter_name[]" binding:"max=255"`
	FilterValue            []string `json:"filter_value" form:"filter_value[]" binding:"max=255"`
	MerchantId             int64    `json:"-" form:"-"`
	UserVerificationStatus *int     `json:"-" form:"-"`
	CheckUpdatedAtMore     int64    `json:"-" form:"-"`
	CheckUpdatedAtLess     int64    `json:"-" form:"-"`
	CheckChatViewed        bool     `json:"-" form:"-"`
	CheckChatUnViewed      bool     `json:"-" form:"-"`
}

func (request RequestDatatable) ValidateOrderColumn(availColumns []string) error {
	if len(availColumns) == 0 {
		return nil
	}
	if utils.InArray(request.OrderColumn, availColumns) {
		return nil
	}

	return errors.New("BadOrderColumns")
}

func (request RequestDatatable) ValidateUserIdRequire() error {
	if len(request.FilterName) == 0 {
		return errors.New("BadRequest")
	}
	if !utils.InArray("user_id", request.FilterName) {
		return errors.New("BadRequest")
	}

	return nil
}
