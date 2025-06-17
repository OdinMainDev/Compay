package escache

import (
	"reflect"
)

func Delete[Model any](model Model, args ...int64) {
	key := createKeyFor(model, args...)
	dto := getDtoByKey(key)
	if dto == nil {
		return
	}

	keyId := any(model).(CachableModel).GetId()
	index, exists := dto.Indexes[keyId]
	if !exists || reflect.TypeOf(dto.Data).Kind() != reflect.Slice {
		return
	}

	v := reflect.ValueOf(dto.Data)
	if index >= v.Len() {
		return
	}

	dto.Data = reflect.AppendSlice(v.Slice(0, index), v.Slice(index+1, v.Len())).Interface()
	delete(dto.Indexes, keyId)
	for key, value := range dto.Indexes {
		if value > index {
			dto.Indexes[key] = value - 1
		}
	}
	setDtoByKey(key, *dto)
}
