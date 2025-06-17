package escache

import (
	"fmt"
	"reflect"
)

func RegisterType(instance any) {
	t := reflect.TypeOf(instance)
	esredis.RegisterType(t)
	types[t] = true
}

func IsCachableModel(model any) bool {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Slice {
		t = t.Elem()
	}
	return types[t]
}

func createKey[Model any](args ...int64) string {
	modelType := reflect.TypeOf((*Model)(nil)).Elem()
	if modelType.Kind() == reflect.Pointer {
		modelType = modelType.Elem()
	}

	key := fmt.Sprint(modelType)
	if len(args) > 0 {
		key += "-" + fmt.Sprint(args)
	}

	return key
}

func createKeyFor(model any, args ...int64) string {
	modelType := reflect.TypeOf(model).Elem()
	if modelType.Kind() == reflect.Pointer {
		modelType = modelType.Elem()
	}

	key := fmt.Sprint(modelType)
	if len(args) > 0 {
		key += "-" + fmt.Sprint(args)
	}

	return key
}

func ClearCache() {
	if manager != nil {
		manager.Clear(ctx)
	}
}

type cacheDTO struct {
	Data    any
	Indexes map[int64]int
}

func modelsToCacheDTO[Model any](models []Model) cacheDTO {
	dto := cacheDTO{
		Data:    models,
		Indexes: make(map[int64]int, len(models)),
	}

	for index, model := range models {
		dto.Indexes[any(model).(CachableModel).GetId()] = index
	}

	return dto
}
