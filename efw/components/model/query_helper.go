package model

type QueryHelper struct {
	HasPagination
}

func (helper QueryHelper) GetPagination() *HasPagination  { return &helper.HasPagination }
func (helper QueryHelper) CacheAdditionalKeys() []int64   { return nil }
