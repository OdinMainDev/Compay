package basemodel

func PaginateModels[T any](items []T, page int) []T {
	perPage := 10
	if page < 1 {
		page = 1
	}

	start := (page - 1) * perPage
	if start >= len(items) {
		return []T{}
	}
	end := start + perPage
	if end > len(items) {
		end = len(items)
	}
	return items[start:end]
}
