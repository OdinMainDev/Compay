package merchant

import (
	"compay/application/entities/datatable"
	"compay/application/entities/merchantuser"
	"compay/application/helpers"
	"compay/application/http/requests"
	"compay/application/service/esvalidator"
	"compay/common/constants"
	"compay/core/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestUserProcess struct {
	Id         int64 `json:"id" binding:"min=0"`
	IsApproved bool  `json:"is_approved"`
	MerchantId int64 `json:"merchant_id" binding:"min=0"`
}

var usersColumnsOrder = []string{"id", "name", "token", "is_enabled", "created_at"}

func UsersList(context *gin.Context) {
	requests.DatatableResponse[merchantuser.Model](context, usersColumnsOrder, func(_ int64) int64 {
		return merchantuser.GetCount()
	}, nil)
}

func UsersListFilters(context *gin.Context) {
	helpers.GinSuccessResponse(context, gin.H{"data": []datatable.ResponseFilter{
		{
			Name: "id",
			Type: datatable.FilterTypeNumeric,
		},
		{
			Name: "name",
			Type: datatable.FilterTypeText,
		},
		{
			Name: "token",
			Type: datatable.FilterTypeText,
		},
		{
			Name: "is_enabled",
			Type: datatable.FilterTypeBool,
		},
		{
			Name: "created_at",
			Type: datatable.FilterTypeDate,
		},
	}})
}

func UserGet(context *gin.Context) {
	requests.GetRecord(context, merchantuser.GetModelById, merchantuser.QueryHelper{})
}

func UserProcess(context *gin.Context) {
	var input requestUserProcess

	if err := context.ShouldBindJSON(&input); err != nil {
		helpers.GinErrorResponse(context, esvalidator.ValidatorResponse(err))
		return
	}

	userModel := merchantuser.GetModel(merchantuser.QueryHelper{
		Id: input.Id,
	})
	if userModel == nil || userModel.GetId() == 0 {
		helpers.GinErrorResponse(context, nil, http.StatusNotFound)
		return
	}
	if utils.InArray(userModel.TgId, constants.AdminTelegramIds) {
		helpers.GinErrorResponse(context, nil, http.StatusNotFound)
		return
	}

	userModel.IsApproved = input.IsApproved

	adminUserModel := context.MustGet(constants.GinUser).(*merchantuser.Model)
	if utils.InArray(adminUserModel.TgId, constants.AdminTelegramIds) {
		userModel.MerchantId = input.MerchantId
	}

	userModel.Save(userModel)

	helpers.GinSuccessResponse(context, gin.H{"data": userModel})
}
