package efwdt

import "github.com/gin-gonic/gin"

type Response[T any] struct {
	RecordsTotal    int64 `json:"recordsTotal"`
	RecordsFiltered int64 `json:"recordsFiltered"`
	RecordsCount    int64 `json:"recordsCount"`
	Data            []T   `json:"data"`
}

func responseEmpty[T any](totalRecords int64) gin.H {
	return gin.H{
		"data": Response[T]{
			Data:         []T{},
			RecordsTotal: totalRecords,
		},
	}
}
