package merchantuser

import (
	"compay/core"
	"compay/core/eslog"
	"errors"

	"gorm.io/gorm"
)

type QueryHelper struct {
	Id         int64
	TelegramId int64
}

func GetModel(data QueryHelper) *Model {
	var model *Model
	query := prepareUserQuery(data)
	tx := query.First(&model)
	if tx.Error == nil {
		return model
	}

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return &Model{}
	}

	eslog.LogError(tx.Error)
	return &Model{}
}

func GetModels(data QueryHelper) []*Model {
	var model []*Model
	query := prepareUserQuery(data)

	tx := query.Find(&model)
	if tx.Error == nil {
		return model
	}

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return []*Model{}
	}

	eslog.LogError(tx.Error)
	return []*Model{}
}

func prepareUserQuery(data QueryHelper) *gorm.DB {
	query := core.GetInstance().GetDB()

	if data.Id != 0 {
		query = query.Where("id = ?", data.Id)
	}

	if data.TelegramId != 0 {
		query = query.Where("tg_id = ?", data.TelegramId)
	}

	query = query.Limit(50)
	return query
}

func GetCount() int64 {
	var result int64 = 0
	tx := core.GetInstance().GetDB().Model(&Model{}).Count(&result)

	if tx.Error != nil {
		eslog.LogError(tx.Error)
	}

	return result
}

func GetModelById(userId int64, data QueryHelper) *Model {
	data.Id = userId
	var model *Model
	query := prepareUserQuery(data)
	tx := query.First(&model)
	if tx.Error == nil {
		return model
	}

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return &Model{}
	}

	eslog.LogError(tx.Error)
	return &Model{}
}
