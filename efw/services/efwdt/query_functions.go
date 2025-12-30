package efwdt

import (
	"maps"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

func applyAdditionalFilters[T any](chain gorm.ChainInterface[T], filters []ChainFilter[T]) gorm.ChainInterface[T] {
	for _, filter := range filters {
		if filter == nil {
			continue
		}
		chain = filter(chain)
	}
	return chain
}

func applyRequestFilters[TModel, TRequest any](chain gorm.ChainInterface[TModel], request Request[TRequest]) gorm.ChainInterface[TModel] {
	filters, specialFilters := buildFiltersMap(reflect.ValueOf(request.Filters))
	if len(filters) == 0 && len(specialFilters) == 0 {
		return chain
	}
	if len(filters) > 0 {
		chain = chain.Where(filters)
	}
	for _, condition := range specialFilters {
		chain = chain.Where(condition.Query, condition.Args...)
	}
	return chain
}

func buildFiltersMap(value reflect.Value) (map[string]any, []SpecialFilters) {
	if value.Kind() != reflect.Struct {
		return nil, nil
	}

	result := make(map[string]any)
	specialFilters := make([]SpecialFilters, 0)
	valueType := value.Type()
	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		if field.PkgPath != "" {
			continue
		}

		fieldValue := value.Field(i)
		if field.Anonymous && fieldValue.Kind() == reflect.Struct {
			r, sF := buildFiltersMap(fieldValue)
			maps.Copy(result, r)
			specialFilters = append(specialFilters, sF...)
			continue
		}

		fieldName := efwdtTagName(field)
		if fieldName == "-" {
			continue
		}

		fieldValueInterface, ok := normalizeConditionValue(fieldValue)
		if !ok {
			continue
		}

		if applied := applySpecialFilters(fieldName, fieldValueInterface, &specialFilters); applied {
			continue
		}

		result[fieldName] = fieldValueInterface
	}
	if len(result) == 0 && len(specialFilters) == 0 {
		return nil, nil
	}

	if len(result) == 0 {
		return nil, specialFilters
	}
	return result, specialFilters
}

func efwdtTagName(field reflect.StructField) string {
	tag := field.Tag.Get("efwdt")
	if tag == "-" {
		return "-"
	}
	if tag == "" {
		return field.Name
	}
	name := strings.Split(tag, ",")[0]
	if name == "" {
		return field.Name
	}
	return name
}

func applySpecialFilters(fieldName string, value any, specialFilters *[]SpecialFilters) bool {
	valueCheck := reflect.ValueOf(value)
	if valueCheck.IsValid() && valueCheck.Kind() == reflect.Int64 && valueCheck.Int() == 0 {
		return true
	}

	switch fieldName {
	case "created_at_from":
		*specialFilters = append(*specialFilters, SpecialFilters{Query: "created_at > ?", Args: []any{value}})
		return true
	case "created_at_to":
		*specialFilters = append(*specialFilters, SpecialFilters{Query: "created_at < ?", Args: []any{value}})
		return true
	case "updated_at_from":
		*specialFilters = append(*specialFilters, SpecialFilters{Query: "updated_at > ?", Args: []any{value}})
		return true
	case "updated_at_to":
		*specialFilters = append(*specialFilters, SpecialFilters{Query: "updated_at < ?", Args: []any{value}})
		return true
	default:
		return false
	}
}

func normalizeConditionValue(value reflect.Value) (any, bool) {
	if !value.IsValid() {
		return nil, false
	}

	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return nil, false
		}
		return value.Elem().Interface(), true
	}

	if (value.Kind() == reflect.Slice || value.Kind() == reflect.Array) && value.Len() == 0 {
		return nil, false
	}

	if value.IsZero() {
		return nil, false
	}

	return value.Interface(), true
}

func buildDataQuery[TModel, TRequest any](dbTrx gorm.ChainInterface[TModel], request Request[TRequest]) gorm.ChainInterface[TModel] {
	return dbTrx.Order(request.OrderColumn + " " + request.OrderDirection).Offset(int(request.Start)).Limit(request.Length)
}

func calcRecordsCount(col any) int {
	if col == nil {
		return 0
	}
	value := reflect.ValueOf(col)
	for value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
		if value.IsNil() {
			return 0
		}
		value = value.Elem()
	}
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		return value.Len()
	}
	return 0
}
