package basemodel

import (
	"compay/core"
	"compay/core/escache"
	"compay/core/eslog"
)

func (parentModel *BaseModel) Delete(value any) {
	parentModel.deleteFromDb(value)
}

func (parentModel *BaseCacheModel) Delete(value any) {
	parentModel.deleteFromCache(value)
}

func (parentModel *BaseModel) deleteFromDb(value any) {
	tx := core.GetInstance().GetDB().Delete(value)
	if tx.Error != nil {
		// TODO: улучшить логирование ошибок
		eslog.LogError(tx.Error)
		return
	}
	if !escache.IsCachableModel(value) {
		return
	}

	escache.Delete(value)
}

func (parentModel *BaseModel) deleteFromCache(value any) {
	escache.Delete(value)
}
