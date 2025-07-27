package basemodel

import (
	"compay/core"
	"compay/core/escache"
	"compay/core/eslog"
)

func (parentModel *BaseModel) deleteFromDb(value any, withMerchantId bool) {
	tx := core.GetInstance().GetDB().Delete(value)
	if tx.Error != nil {
		// TODO: улучшить логирование ошибок
		eslog.LogError(tx.Error)
		return
	}
	if !escache.IsCachableModel(value) {
		return
	}

	if !withMerchantId {
		escache.Delete(value)
		return
	}

	escache.Delete(value, *value.(modelWithMerchantId).GetMerchantId())
}

func (parentModel *BaseModel) deleteFromCache(value any, withMerchantId bool) {
	if !withMerchantId {
		escache.Delete(value)
		return
	}

	escache.Delete(value, *value.(modelWithMerchantId).GetMerchantId())
}
