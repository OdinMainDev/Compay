package model

import "efw/core/efwcache"

func getModels[T Model, F Filter[T]](filter F, additionalCacheKeys ...int64) ([]*T, error) {
	records, err := doGetModels(filter, additionalCacheKeys...)
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return []*T{}, nil
	}

	result := make([]*T, 0, len(records))
	for i := range records {
		if filter.ApplyCacheFilters(&records[i]) {
			result = append(result, &records[i])
		}
	}
	return result, nil
}

func doGetModels[T Model](filter Filter[T], additionalCacheKeys ...int64) ([]T, error) {
	var m T
	if isCacheModel(m) {
		if cached, err := efwcache.Get[T](additionalCacheKeys...); err == nil {
			return cached, nil
		} else {
			return nil, err
		}
	}

	if !isDatabaseModel(m) {
		return nil, nil
	}

	records, err := fetchFromDb(filter)
	if err != nil {
		return nil, err
	}

	if isCacheModel(m) {
		if err := efwcache.Set(records, additionalCacheKeys...); err != nil {
			return records, err
		}
	}

	return records, nil
}

func getModel[T Model, F Filter[T]](filter F, additionalCacheKeys ...int64) (*T, error) {
	records, err := getModels(filter, additionalCacheKeys...)
	if len(records) > 0 && err == nil {
		return records[0], nil
	}

	return nil, err
}

func getCount[T Model, F Filter[T]](filter F, additionalCacheKeys ...int64) (int64, error) {
	records, err := getModels(filter, additionalCacheKeys...)
	if err != nil {
		return 0, err
	}
	return int64(len(records)), nil
}
