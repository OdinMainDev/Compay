package interfaces

type IdInt64 interface{ ~int64 }

type SaveType int // Тип сохранения
const (           //
	SAVE_TYPE_DB_CACHE SaveType = iota
	SAVE_TYPE_DB
	SAVE_TYPE_CACHE
	SAVE_TYPE_LOCAL_CACHE
)

// абстракция определяющая наличие модели
type HasModel[ID IdInt64] interface {
	GetId() ID
	GetIdInt64() int64
	GetCreatedAt() int64
	GetUpdatedAt() int64
	GetSaveType() SaveType

	// GetTable

	Lock()
	Unlock()
}
