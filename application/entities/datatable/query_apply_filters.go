package datatable

import (
	"compay/core/utils"
	"reflect"
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
		case "id", "position":
			tx = tx.Where("concat("+quoted+") LIKE ?", strings.ToLower(val)+"%")
		case "ids":
			tx = tx.Where("id in (?)", strings.Split(val, ","))
		case "casino_ids":
			tx = tx.Where("casino_id in (?)", strings.Split(val, ","))
		case "session_id", "user_id", "casino_id", "payment_id", "withdraw_id", "type", "minor_units", "admin_id", "withdraw_type":
			tx = tx.Where(column+" = ?", strings.ToLower(val))

		case "is_enabled", "is_viewed", "is_approved", "is_premium", "is_active":
			switch reflect.TypeOf(model).String() {
			case "[]*game.GameModel":
				if request.CasinoId == 0 {
					tx = tx.Where(column+" = ?", getBoolValue(val))
				}
			default:
				tx = tx.Where(column+" = ?", getBoolValue(val))
			}

		default:
			if column == "type" && reflect.TypeOf(model).String() == "[]*chatmessage.Model" {
				if val == "file" {
					tx = tx.Where("data->>'file_url' IS NOT NULL AND data->>'file_url' != ''")
				}
				break
			}
			tx = tx.Where("lower("+quoted+") LIKE ?", "%"+strings.ToLower(val)+"%")
		}
	}
	return tx
}
