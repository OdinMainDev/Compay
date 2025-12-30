package merchant

import (
	"compay/common/types"
	"compay/src/models/user"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/services/efwdt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestAdminUsersList struct {
	Id         int64  `json:"id"  binding:"omitempty,gt=0"`
	MerchantId string `json:"merchant_id" binding:"omitempty,gt=0"`
	Username   string `json:"username" binding:"omitempty,max=255"`
	IsBanned   *bool  `json:"is_banned"`
}

func AdminUsersList(context *gin.Context) {
	efwdt.Datable[user.Model, requestAdminUsersList](
		context,
		[]string{"id", "username", "is_banned", "created_at"},
		func() int64 { return user.GetCount(user.QueryHelper{}) },
		nil,
	)
}

func AdminUsersListFilters(context *gin.Context) {
	efwdt.Filters(context, []efwdt.Filter{
		{Name: "id", Type: types.EFWDT_FilterTypeInteger},
		{Name: "merchant_id", Type: types.EFWDT_FilterTypeMerchantId},
		{Name: "username", Type: types.EFWDT_FilterTypeText},
		{Name: "is_banned", Type: types.EFWDT_FilterTypeBoolean},
		{Name: "example_list", Type: types.EFWDT_FilterTypeList, Values: []string{"val1", "val2"}},
		{Name: "example_list_number", Type: types.EFWDT_FilterTypeList, Values: []int64{1, 2}},
		{Name: "created_at", Type: types.EFWDT_FilterTypeDate},
	})
}

func AdminUsersDelete(context *gin.Context) {
	var input requestWithId
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}
	userModel := user.GetModel(user.QueryHelper{
		Id:                  types.UserId(input.Id),
		DontCheckMerchantId: true,
	})
	if userModel == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}
	user.Delete(userModel)
	efwginhelpers.ResponseSuccess(context, nil)
}
