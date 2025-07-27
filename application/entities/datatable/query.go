package datatable

import (
	"compay/core"
	"compay/core/eslog"
	"errors"
	"strings"

	"gorm.io/gorm"
)

func DataTableQuery(model interface{}, req RequestDatatable, availCols []string) interface{} {
	db := core.GetInstance().GetDB()

	baseTx := prepareBaseQuery(db, req)

	filteredTx := applyFilters(baseTx, db, model, req, availCols)

	result, err := applySpecialQueries(filteredTx, db, model, req)
	if err != nil {
		eslog.LogError(err)
		return nil
	}
	return result
}

func prepareBaseQuery(db *gorm.DB, request RequestDatatable) *gorm.DB {
	return db.Order(request.OrderColumn + " " + request.OrderDirection).Offset(int(request.Start)).Limit(request.Length)
}

func applySpecialQueries(tx *gorm.DB, _ *gorm.DB, model interface{}, request RequestDatatable) (interface{}, error) {
	return defaultQuery(tx, model, request)
}

func defaultQuery(tx *gorm.DB, model interface{}, request RequestDatatable) (interface{}, error) {
	if request.MerchantId != 0 {
		tx = tx.Where("merchant_id", request.MerchantId)
	}

	if resp := tx.Find(&model); resp.Error != nil {
		if errors.Is(resp.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, resp.Error
	}
	return model, nil
}

func getBoolValue(val string) string {
	boolStr := "TRUE"
	if strings.ToLower(val) == "false" {
		boolStr = "FALSE"
	}
	return boolStr
}
