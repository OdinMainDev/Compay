package basemodel

type BaseModel struct {
	Id        int64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`
}

type BaseModelWithMerchantId struct {
	BaseModel
	MerchantId int64 `json:"merchant_id" gorm:"column:merchant_id"`
}

type modelWithMerchantId interface {
	GetMerchantId() *int64
}

// GetId функции

func (parentModel *BaseModel) GetId() int64 {
	return parentModel.Id
}

func (parentModel *BaseModelWithMerchantId) GetId() int64 {
	return parentModel.Id
}

// GetMerchantId функции

func (parentModel *BaseModel) GetMerchantId() *int64 {
	return nil
}

func (parentModel *BaseModelWithMerchantId) GetMerchantId() *int64 {
	return &parentModel.MerchantId
}

// GetCreatedAt функции

func (parentModel *BaseModel) GetCreatedAt() int64 {
	return parentModel.CreatedAt
}

func (parentModel *BaseModelWithMerchantId) GetCreatedAt() int64 {
	return parentModel.CreatedAt
}

// GetUpdatedAt функции

func (parentModel *BaseModel) GetUpdatedAt() int64 {
	return parentModel.UpdatedAt
}

func (parentModel *BaseModelWithMerchantId) GetUpdatedAt() int64 {
	return parentModel.UpdatedAt
}

// Save функции

func (parentModel *BaseModel) Save(value any) {
	parentModel.saveToDb(value, false)
}

func (parentModel *BaseModelWithMerchantId) Save(value any) {
	parentModel.saveToDb(value, true)
}

// Delete функции

func (parentModel *BaseModel) Delete(value any) {
	parentModel.deleteFromDb(value, false)
}

func (parentModel *BaseModelWithMerchantId) Delete(value any) {
	parentModel.deleteFromDb(value, true)
}
