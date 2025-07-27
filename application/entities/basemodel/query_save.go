package basemodel

import (
	"compay/core"
	"compay/core/escache"
	"compay/core/eslog"
	"time"

	"gorm.io/gorm/clause"
)

func (parentModel *BaseModel) saveToDb(value any, withMerchantId bool) {
	currentTime := time.Now().Unix()
	if parentModel.CreatedAt == 0 {
		parentModel.CreatedAt = currentTime
	}
	parentModel.UpdatedAt = currentTime
	tx := core.GetInstance().GetDB().Omit(clause.Associations).Save(value)
	if tx.Error != nil {
		// TODO: улучшить логирование ошибок
		eslog.LogError(tx.Error)
		return
	}

	if !escache.IsCachableModel(value) {
		return
	}

	if !withMerchantId {
		escache.Update(value)
		return
	}

	escache.Update(value, *value.(modelWithMerchantId).GetMerchantId())
}

func (parentModel *BaseModel) saveToCache(value any, withMerchantId bool) {
	currentTime := time.Now().Unix()
	if parentModel.GetId() == 0 {
		parentModel.CreatedAt = currentTime
	}
	parentModel.UpdatedAt = currentTime

	if !withMerchantId {
		escache.Update(value)
		return
	}

	escache.Update(value, *value.(modelWithMerchantId).GetMerchantId())
}
