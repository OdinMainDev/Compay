package notification

import (
	"compay/common/types"
	"efw/components/model"

	"gorm.io/gorm"
)

type QueryHelper struct {
	model.QueryHelper
	Id       types.NotificationId
	IsRead   bool
	IsUnread bool
}

func (helper QueryHelper) CacheAdditionalKeys() []int64         { return nil }
func (helper QueryHelper) ApplyCacheFilters(entity *Model) bool { return true }

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	if helper.IsRead {
		chain = chain.Where("is_read = ?", true)
	}
	if helper.IsUnread {
		chain = chain.Where("is_read = ?", false)
	}
	return chain.Order("id DESC")
}
