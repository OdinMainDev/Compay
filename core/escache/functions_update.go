package escache

import (
	"reflect"
	"sort"
)

func Update[Model any](model Model, args ...int64) {
	key := createKeyFor(model, args...)
	dto := getDtoByKey(key)
	if dto == nil {
		return
	}

	keyId := any(model).(CachableModel).GetId()
	index, exists := dto.Indexes[keyId]
	dtoDataValue := reflect.ValueOf(dto.Data)
	if !exists {
		dtoDataValue = reflect.Append(dtoDataValue, reflect.ValueOf(model))
		dto.Indexes[any(model).(CachableModel).GetId()] = len(dto.Indexes)
	} else {
		dtoDataElement := dtoDataValue.Index(index)
		dtoDataElement.Set(reflect.ValueOf(model))
	}

	dto.Data = dtoDataValue.Interface()

	sortDataIfNeeded(dto)
	setDtoByKey(key, *dto)
}

func sortDataIfNeeded(dto *cacheDTO) {
	dataVal := reflect.ValueOf(dto.Data)
	if dataVal.Kind() != reflect.Slice || dataVal.Len() == 0 {
		return
	}

	elem := dataVal.Index(0)
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	if elem.Kind() != reflect.Struct {
		return
	}
	posField := elem.FieldByName("Position")
	if !posField.IsValid() || posField.Kind() != reflect.Int {
		return
	}

	sliceInterface := dto.Data
	sort.SliceStable(sliceInterface, func(i, j int) bool {
		sliceVal := reflect.ValueOf(sliceInterface)
		vi := sliceVal.Index(i)
		vj := sliceVal.Index(j)

		if vi.Kind() == reflect.Ptr {
			vi = vi.Elem()
		}
		if vj.Kind() == reflect.Ptr {
			vj = vj.Elem()
		}
		return vi.FieldByName("Position").Int() < vj.FieldByName("Position").Int()
	})
}
