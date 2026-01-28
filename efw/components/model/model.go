package model

import (
	"efw/components/entity"

	"gorm.io/gorm"
)

// Model это абстрактный класс для определения является ли вложение моделью или нет
type Model interface {
	GetIdInt64() int64
	GetCreatedAt() int64
	GetUpdatedAt() int64
	GetSaveType() entity.SaveType

	Lock()
	Unlock()
	UpdateTimestamps()
}

// Filter это абстрактный класс который наследуется от Model. Служит для фильтрации в базе/кеш данных для получения моделей
type Filter[T Model] interface {
	CacheAdditionalKeys(T) []int64
	ApplyDBFilters(gorm.ChainInterface[T]) gorm.ChainInterface[T]
	ApplyCacheDBFilters(gorm.ChainInterface[T]) gorm.ChainInterface[T]
	ApplyCacheFilters(T) bool
	GetPagination() *HasPagination
}
