package efwcontroller

type RequestGetById struct {
	Id int64 `uri:"id" json:"id" binding:"required,min=0"`
}
