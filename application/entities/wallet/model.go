package wallet

import (
	"compay/application/entities/basemodel"
	"compay/core/escache"
)

type QueryHelper struct {
	Id int64
}

func getAndSort(data QueryHelper) []*Model {
	rows := escache.Get[*Model]()
	if rows == nil {
		rows = basemodel.GetAllFromDbQuery(rows, basemodel.QueryHelper{}).([]*Model)
		escache.Set(rows)
	}
	ret := []*Model{}
	for _, item := range rows {

		if data.Id != 0 && item.Id != data.Id {
			continue
		}

		ret = append(ret, item)
	}

	return ret
}

func GetModel(data QueryHelper) *Model {
	models := getAndSort(data)
	if len(models) > 0 {
		return models[0]
	}

	return &Model{}
}

func GetModels(data QueryHelper) []*Model {
	return getAndSort(data)
}

func GetCount() int64 {
	return int64(len(getAndSort(QueryHelper{})))
}

func GetModelById(id int64, data QueryHelper) *Model {
	data.Id = id
	return GetModel(data)
}
