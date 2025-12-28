package efwdt

import "github.com/gin-gonic/gin"

type Response[T any] struct {
	Data            T     `json:"data"`
	RecordsTotal    int64 `json:"recordsTotal"`
	RecordsFiltered int64 `json:"recordsFiltered"`
	RecordsCount    int64 `json:"recordsCount"`
}

func responseEmpty(recordsTotal int64) gin.H {
	return gin.H{
		"data": Response[[]int64]{
			Data:         []int64{},
			RecordsTotal: recordsTotal,
		},
	}
}

func response[T any](data T, recordsTotal, recordsFiltered, recordsCount int64) gin.H {
	return gin.H{
		"data": Response[T]{
			Data:            data,
			RecordsTotal:    recordsTotal,
			RecordsFiltered: recordsFiltered,
			RecordsCount:    recordsCount,
		},
	}
}
