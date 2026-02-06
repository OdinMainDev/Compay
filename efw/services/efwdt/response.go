package efwdt

import "github.com/gin-gonic/gin"

type response[T any] struct {
	Data            T     `json:"data"`
	RecordsTotal    int64 `json:"recordsTotal"`
	RecordsFiltered int64 `json:"recordsFiltered"`
	RecordsCount    int64 `json:"recordsCount"`
}

func Response[T any](models []*T, recordsTotal, recordsFiltered, recordsCount int64, dtoFunc func(model *T) gin.H) gin.H {
	if dtoFunc == nil {
		return gin.H{
			"data": gin.H{
				"data":            models,
				"recordsTotal":    recordsTotal,
				"recordsFiltered": recordsFiltered,
				"recordsCount":    recordsCount,
			},
		}
	}

	var resp []any
	for _, model := range models {
		resp = append(resp, dtoFunc(model))
	}
	return gin.H{
		"data": gin.H{
			"data":            resp,
			"recordsTotal":    recordsTotal,
			"recordsFiltered": recordsFiltered,
			"recordsCount":    recordsCount,
		},
	}
}
