package efwcache

func Get[T CacheableModel](args ...int64) ([]T, error) {
	instanceMu.RLock()
	defer instanceMu.RUnlock()

	if instance == nil {
		return nil, ErrorCacheNotInitialized
	}

	key := constructKey[T](args...)
	items, err := instance.get(key)
	if err != nil {
		return nil, err
	}

	result := make([]T, 0, len(items))
	for _, item := range items {
		if model, ok := item.Data.(T); ok {
			result = append(result, model)
		}
	}

	return result, nil
}

func Set[T CacheableModel](models []T, args ...int64) error {
	instanceMu.RLock()
	defer instanceMu.RUnlock()

	if instance == nil {
		return ErrorCacheNotInitialized
	}

	key := constructKey[T](args...)
	items := make([]CacheItem, len(models))

	for i, model := range models {
		items[i] = CacheItem{
			Id:   model.GetIdInt64(),
			Data: model,
		}
	}

	sortDataIfNeeded(items)

	return instance.set(key, items)
}

func Update[T CacheableModel](model T, args ...int64) error {
	instanceMu.RLock()
	defer instanceMu.RUnlock()

	if instance == nil {
		return ErrorCacheNotInitialized
	}

	key := constructKeyFromModel(model, args...)
	item := CacheItem{
		Id:   model.GetIdInt64(),
		Data: model,
	}

	return instance.update(key, item)
}

func Delete[T CacheableModel](model T, args ...int64) error {
	instanceMu.RLock()
	defer instanceMu.RUnlock()

	if instance == nil {
		return ErrorCacheNotInitialized
	}

	key := constructKeyFromModel(model, args...)
	return instance.delete(key, model.GetIdInt64())
}

func ClearCache() error {
	instanceMu.RLock()
	defer instanceMu.RUnlock()

	if instance == nil {
		return ErrorCacheNotInitialized
	}

	instance.clear()
	return nil
}
