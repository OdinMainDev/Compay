package model

func GetModels[T Model, F Filter[T]](filter F, additionalCacheKeys ...int64) ([]T, error) {
	return getModels(filter, additionalCacheKeys...)
}

func Save[T Model](model T, additionalCacheKeys ...int64) error {
	return save(model, additionalCacheKeys...)
}

func Delete[T Model](model T, additionalCacheKeys ...int64) error {
	return delete(model, additionalCacheKeys...)
}

func GetModel[T Model, F Filter[T]](filter F, additionalCacheKeys ...int64) (T, error) {
	return getModel(filter, additionalCacheKeys...)
}

func GetCount[T Model, F Filter[T]](filter F, additionalCacheKeys ...int64) (int64, error) {
	return getCount(filter, additionalCacheKeys...)
}
