package datatable

import "reflect"

func (request RequestDatatable) ResponseEmpty(totalRecords int64) Response {
	return Response{
		Data:         []any{},
		RecordsTotal: totalRecords,
	}
}

func (request RequestDatatable) Response(totalRecords int64, data interface{}) Response {
	return Response{
		Data:            data,
		RecordsTotal:    totalRecords,
		RecordsFiltered: recordsCount(data),
	}
}

func recordsCount(col interface{}) int {
	if col == nil {
		return 0
	}
	v := reflect.ValueOf(col)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		return v.Len()
	}

	return 0
}
