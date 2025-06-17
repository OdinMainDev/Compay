package basemodel

import (
	"compay/core"
	"compay/core/eslog"
	"errors"

	"gorm.io/gorm"
)

type QueryHelper struct {
	Order    string
	Preloads map[string][]any
}

const (
	OrderDefault    = "id desc"
	OrderDefaultAsc = "id asc"
	OrderPosition   = "position asc"
)

func GetAllFromDbQuery(value any, helper QueryHelper) any {
	if helper.Order == "" {
		helper.Order = OrderDefault
	}
	query := core.GetInstance().GetDB()

	for preload, args := range helper.Preloads {
		query = query.Preload(preload, args)
	}

	tx := query.Order(helper.Order).Find(&value)
	if tx.Error == nil {
		return value
	}

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	eslog.LogError(tx.Error)
	return nil
}
