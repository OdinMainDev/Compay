package merchant

import "efw/components/entity"

type requestWithId struct {
	Id int64 `uri:"id" json:"id" binding:"required,min=0"`
}

type requestSearch struct {
	Search string `form:"search" query:"search" uri:"search" json:"search" binding:"omitempty,max=255"`
}

func isUserRootAdmin(userId entity.UserId) bool {
	return userId == 1
}
