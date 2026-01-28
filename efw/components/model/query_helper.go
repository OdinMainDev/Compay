package model

type QueryHelper struct {
	HasPagination
}

func (helper QueryHelper) GetPagination() *HasPagination { return &helper.HasPagination }
