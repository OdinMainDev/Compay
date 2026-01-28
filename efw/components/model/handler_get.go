package model

import (
	"efw/core/efwcache"
)

func getModels[T Model, F Filter[T]](filter F) ([]T, error) {
	records, isCachedModel, err := doGetModels(filter)
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return []T{}, nil
	}

	if !isCachedModel {
		return records, nil
	}

	result := make([]T, 0, len(records))
	for i := range records {
		if filter.ApplyCacheFilters(records[i]) {
			result = append(result, records[i])
		}
	}
	return result, nil
}

func doGetModels[T Model](filter Filter[T]) ([]T, bool, error) {
	var m T
	if isCacheModel(m) {
		cached, err := efwcache.Get[T](filter.CacheAdditionalKeys()...)
		if err != nil {
			return nil, true, err
		}
		if len(cached) > 0 {
			return cached, true, nil
		}
	}

	if !isDatabaseModel(m) {
		return nil, true, nil
	}

	records, err := fetchFromDb(filter)
	if err != nil {
		return nil, false, err
	}

	if isCacheModel(m) {
		if err := efwcache.Set(records, filter.CacheAdditionalKeys()...); err != nil {
			return records, false, err
		}
	}

	return records, false, nil
}

func getModel[T Model, F Filter[T]](filter F) (T, error) {
	records, err := getModels(filter)
	if len(records) > 0 && err == nil {
		return records[0], nil
	}

	var zero T
	return zero, err
}

func getCount[T Model, F Filter[T]](filter F) (int64, error) {
	records, err := getModels(filter)
	if err != nil {
		return 0, err
	}
	return int64(len(records)), nil
}
