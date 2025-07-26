package merchantuser

import (
	"compay/core"
	"compay/core/eslog"
	"errors"

	"gorm.io/gorm"
)

type QueryUserHelper struct {
	Id         int64
	TelegramId int64
}

func GetUserModel(data QueryUserHelper) *UserModel {
	var model *UserModel
	query := prepareUserQuery(data)
	tx := query.First(&model)
	if tx.Error == nil {
		return model
	}

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return &UserModel{}
	}

	eslog.LogError(tx.Error)
	return &UserModel{}
}

func GetUserModels(data QueryUserHelper) []*UserModel {
	var model []*UserModel
	query := prepareUserQuery(data)

	tx := query.Find(&model)
	if tx.Error == nil {
		return model
	}

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return []*UserModel{}
	}

	eslog.LogError(tx.Error)
	return []*UserModel{}
}

func prepareUserQuery(data QueryUserHelper) *gorm.DB {
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
	tx := core.GetInstance().GetDB().Model(&UserModel{}).Count(&result)

	if tx.Error != nil {
		eslog.LogError(tx.Error)
	}

	return result
}

func GetUserModelById(userId int64, data QueryUserHelper) *UserModel {
	data.Id = userId
	var model *UserModel
	query := prepareUserQuery(data)
	tx := query.First(&model)
	if tx.Error == nil {
		return model
	}

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return &UserModel{}
	}

	eslog.LogError(tx.Error)
	return &UserModel{}
}
