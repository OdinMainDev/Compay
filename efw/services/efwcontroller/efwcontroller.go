package efwcontroller

type RequestWithId struct {
	Id int64 `form:"id" query:"id" uri:"id" json:"id" binding:"required,min=0"`
}
type RequestWithStringId struct {
	Id string `form:"id" query:"id" uri:"id" json:"id" binding:"required,min=0"`
}
