package efwcache

import (
	"reflect"
	"sort"
	"strconv"
)

func sortDataIfNeeded(items []CacheItem) {
	if len(items) == 0 {
		return
	}

	firstElem := reflect.ValueOf(items[0].Data)
	if !firstElem.IsValid() {
		return
	}
	for firstElem.Kind() == reflect.Pointer && !firstElem.IsNil() {
		firstElem = firstElem.Elem()
	}
	if firstElem.Kind() != reflect.Struct {
		return
	}

	posField := firstElem.FieldByName("Position")
	if !posField.IsValid() || posField.Kind() != reflect.Int {
		return
	}

	sort.SliceStable(items, func(i, j int) bool {
		posI := extractPosition(items[i].Data)
		posJ := extractPosition(items[j].Data)
		return posI < posJ
	})
}

func extractPosition(data any) int64 {
	if data == nil {
		return 0
	}

	val := reflect.ValueOf(data)
	for val.Kind() == reflect.Pointer && !val.IsNil() {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return 0
	}

	field := val.FieldByName("Position")
	if !field.IsValid() || field.Kind() != reflect.Int {
		return 0
	}
	return field.Int()
}

func constructKey[T any](args ...int64) string {
	modelType := reflect.TypeOf((*T)(nil)).Elem()
	if modelType.Kind() == reflect.Pointer {
		modelType = modelType.Elem()
	}

	key := modelType.String()
	for _, arg := range args {
		key += "-" + strconv.FormatInt(arg, 10)
	}

	return key
}

func constructKeyFromModel(model any, args ...int64) string {
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Pointer {
		modelType = modelType.Elem()
	}

	key := modelType.String()
	for _, arg := range args {
		key += "-" + strconv.FormatInt(arg, 10)
	}

	return key
}
