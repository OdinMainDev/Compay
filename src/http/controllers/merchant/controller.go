package merchant

type requestWithId struct {
	Id int64 `uri:"id" json:"id" binding:"required,min=0"`
}
