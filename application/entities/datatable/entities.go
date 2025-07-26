package datatable

type Response struct {
	RecordsTotal    int64       `json:"recordsTotal"`
	RecordsFiltered int         `json:"recordsFiltered"`
	Data            interface{} `json:"data"`
}

type Data struct {
	Title   string `json:"Title"`
	Country string `json:"Country"`
	Date    string `json:"Date"`
	Author  string `json:"Author"`
	Id      string `json:"identifier"`
}
