package merchant

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/common/types"
	"compay/src/models/user"
	"efw/components/entity"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/services/efwcontroller"
	"efw/services/efwdt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type requestAdminUsersList struct {
	Id         int64  `json:"id" efwdt:"id" binding:"omitempty,gt=0"`
	MerchantId int64  `json:"merchant_id" efwdt:"merchant_id" binding:"omitempty,gt=0"`
	Username   string `json:"username" efwdt:"username" binding:"omitempty,max=255"`
	IsBanned   *bool  `json:"is_banned" efwdt:"is_banned" binding:"omitempty"`
	efwdt.WithCreatedAt
}

type requestAdminUsersCreate struct {
	Username   string `json:"username" binding:"required,min=6,max=255"`
	Password   string `json:"password" binding:"required,min=6,max=255"`
	IsAdmin    bool   `json:"is_admin"`
	MerchantId int64  `json:"merchant_id" binding:"required_if=IsAdmin false,merchantExist"`
}

type requestAdminUsersProcess struct {
	Id       int64  `json:"id" binding:"required,gt=0"`
	Username string `json:"username" binding:"omitempty,min=6,max=255"`
	Password string `json:"password" binding:"omitempty,min=6,max=255"`
	IsBanned bool   `json:"is_banned"`
}

type requestAdminUsersAction struct {
	Id int64 `json:"id" binding:"required,gt=0"`
}

func AdminUsersList(context *gin.Context) {
	efwcontroller.Datable[user.Model, requestAdminUsersList](
		context,
		[]string{"id", "username", "is_banned", "merchant_id", "created_at"},
		[]efwdt.ChainFilter[user.Model]{
			func(chain gorm.ChainInterface[user.Model]) gorm.ChainInterface[user.Model] {
				return chain.Where("id != ?", 1)
			},
		},
		nil,
		// efwdt.SpecialFilterMap{
		// 	"merchant_total_profit": func(fieldName string, value any) (efwdt.SpecialFilters, bool) {
		// 		return efwdt.SpecialFilters{
		// 			Query: fieldName + " > ?",
		// 			Args:  []any{value},
		// 		}, true
		// 	},
		// },
	)
}

func AdminUsersListFilters(context *gin.Context) {
	efwcontroller.DatatableFilters(context, []efwdt.Filter{
		{Name: "id", Type: types.EFWDT_FilterTypeInteger},
		{Name: "merchant_id", Type: types.EFWDT_FilterTypeMerchantId},
		{Name: "username", Type: types.EFWDT_FilterTypeText},
		{Name: "is_banned", Type: types.EFWDT_FilterTypeBoolean},
		{Name: "created_at", Type: types.EFWDT_FilterTypeDate},
	})
}

func AdminUsersDelete(context *gin.Context) {
	_, ok := efwcontroller.DeleteModel(
		context,
		user.GetModel,
		user.Delete,
		func(id int64) user.QueryHelper {
			return user.QueryHelper{Id: entity.UserId(id), DontCheckMerchantId: true}
		},
	)
	if ok {
		userModel := context.MustGet(constants.GinUser).(*user.Model)
		createLogUser(context, userModel, constants.LogUserAdminUserDelete)
	}
}

func AdminUsersGet(context *gin.Context) {
	efwcontroller.GetModel(
		context,
		user.GetModel,
		nil,
		func(id int64) user.QueryHelper {
			if isUserRootAdmin(entity.UserId(id)) {
				id = 0
			}
			return user.QueryHelper{Id: entity.UserId(id), DontCheckMerchantId: true}
		},
	)
}

func AdminUsersCreate(context *gin.Context) {
	var input requestAdminUsersCreate
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	if input.IsAdmin && input.MerchantId != 0 {
		efwginhelpers.ResponseError(context, cperrors.ErrorUserCreateAdminWithMerchantIdIsNotAllowed)
		return
	}
	if !input.IsAdmin && input.MerchantId == 0 {
		efwginhelpers.ResponseError(context, cperrors.ErrorUserCreateUserWithoutMerchantIdIsNotAllowed)
		return
	}

	checkUserExist := user.GetModel(user.QueryHelper{Username: input.Username, DontCheckMerchantId: true})
	if checkUserExist != nil {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorUserUsernameIsTaken))
		return
	}

	userModel := &user.Model{}
	userModel.Username = input.Username
	userModel.MerchantId = types.MerchantId(input.MerchantId)
	userModel.NewPassword = input.Password
	user.Save(userModel, user.QueryHelper{})

	createLogUser(context, userModel, constants.LogUserAdminUserCreate)
	efwginhelpers.ResponseSuccess(context, nil)
}

func AdminUsersProcess(context *gin.Context) {
	var input requestAdminUsersProcess
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	userModel := user.GetModel(user.QueryHelper{Id: entity.UserId(input.Id), DontCheckMerchantId: true})
	if userModel == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	if isUserRootAdmin(userModel.GetId()) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	if userModel.Username != input.Username {
		checkUserExist := user.GetModel(user.QueryHelper{Username: input.Username, DontCheckMerchantId: true})
		if checkUserExist != nil {
			efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorUserUsernameIsTaken))
			return
		}
	}

	userModel.Username = input.Username
	userModel.NewPassword = input.Password
	user.Save(userModel, user.QueryHelper{})

	createLogUser(context, userModel, constants.LogUserAdminUserProcess)
	efwginhelpers.ResponseSuccess(context, nil)
}

func AdminUsersActionDisableTotp(context *gin.Context) {
	var input requestAdminUsersAction
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	userModel := user.GetModel(user.QueryHelper{Id: entity.UserId(input.Id), DontCheckMerchantId: true})
	if userModel == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	if isUserRootAdmin(userModel.GetId()) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	if userModel.IsTwofaEnabled {
		userModel.TotpToken = ""
		userModel.IsTwofaEnabled = false
		user.Save(userModel, user.QueryHelper{})
		createLogUser(context, userModel, constants.LogUserAdminUserDisableTotp)
	}

	efwginhelpers.ResponseSuccess(context, nil)
}

func AdminUsersSearch(context *gin.Context) {
	var input requestSearch
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": user.GetModels(user.QueryHelper{
		Search: input.Search,
	})})
}
