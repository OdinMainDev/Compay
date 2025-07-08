package basemodel

type BaseModel struct {
	Id        int64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`
}

type BaseCacheModel struct {
	BaseModel
}

// Общие функции

func (parentModel *BaseModel) GetId() int64 {
	return parentModel.Id
}

func (parentModel *BaseModel) GetCreatedAt() int64 {
	return parentModel.CreatedAt
}

func (parentModel *BaseModel) GetUpdatedAt() int64 {
	return parentModel.UpdatedAt
}
