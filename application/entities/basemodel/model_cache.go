package basemodel

type BaseCacheModel struct {
	BaseModel
}

type BaseCacheModelWithMerchantId struct {
	BaseModelWithMerchantId
}

// GetId функции

func (parentModel *BaseCacheModel) GetId() int64 {
	return parentModel.Id
}

func (parentModel *BaseCacheModelWithMerchantId) GetId() int64 {
	return parentModel.Id
}

// GetMerchantId функции

func (parentModel *BaseCacheModel) GetMerchantId() *int64 {
	return nil
}

func (parentModel *BaseCacheModelWithMerchantId) GetMerchantId() *int64 {
	return &parentModel.MerchantId
}

// GetCreatedAt функции

func (parentModel *BaseCacheModel) GetCreatedAt() int64 {
	return parentModel.CreatedAt
}

func (parentModel *BaseCacheModelWithMerchantId) GetCreatedAt() int64 {
	return parentModel.CreatedAt
}

// GetUpdatedAt функции

func (parentModel *BaseCacheModel) GetUpdatedAt() int64 {
	return parentModel.UpdatedAt
}

func (parentModel *BaseCacheModelWithMerchantId) GetUpdatedAt() int64 {
	return parentModel.UpdatedAt
}

// Save функции

func (parentModel *BaseCacheModel) Save(value any) {
	parentModel.saveToCache(value, false)
}

func (parentModel *BaseCacheModelWithMerchantId) Save(value any) {
	parentModel.saveToCache(value, true)
}

// Delete функции

func (parentModel *BaseCacheModel) Delete(value any) {
	parentModel.deleteFromCache(value, false)
}

func (parentModel *BaseCacheModelWithMerchantId) Delete(value any) {
	parentModel.deleteFromCache(value, true)
}
