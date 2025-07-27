package datatable

import (
	"compay/core/utils"
	"strings"

	"gorm.io/gorm"
)

func applyFilters(tx *gorm.DB, db *gorm.DB, model interface{}, request RequestDatatable, availFilterColumns []string) *gorm.DB {
	for index, column := range request.FilterName {
		if !utils.InArray(column, availFilterColumns) || len(request.FilterValue) < index {
			continue
		}

		val := request.FilterValue[index]
		var buf strings.Builder
		db.QuoteTo(&buf, column)
		quoted := buf.String()

		switch column {
		case "id":
			tx = tx.Where("concat("+quoted+") LIKE ?", strings.ToLower(val)+"%")
		case "ids":
			tx = tx.Where("id in (?)", strings.Split(val, ","))
		case "is_enabled":
			tx = tx.Where(column+" = ?", getBoolValue(val))
		default:
			tx = tx.Where("lower("+quoted+") LIKE ?", "%"+strings.ToLower(val)+"%")
		}
	}
	return tx
}
