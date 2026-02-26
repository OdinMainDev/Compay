package merchant

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/blockchainstate"
	"compay/src/models/user"
	"efw/components/entity"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/services/efwcontroller"
	"efw/services/efwdt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestAdminBlockchainStatesList struct {
	Currency entity.CurrencyCodeAlpha `json:"currency" efwdt:"currency" binding:"omitempty,currencyExist"`
}

type requestAdminBlockchainStatesProcess struct {
	Id           types.BlockchainStateId `json:"id" binding:"required,gt=0"`
	CurrentBlock int64                   `json:"current_block" binding:"required,gt=0"`
}

func AdminBlockchainStatesList(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	efwcontroller.Datable[blockchainstate.Model, requestAdminBlockchainStatesList](
		context,
		[]string{"id", "currency", "created_at"},
		nil,
		nil,
		nil,
	)
}

func AdminBlockchainStatesListFilters(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	efwcontroller.DatatableFilters(context, []efwdt.Filter{
		{Name: "currency", Type: types.EFWDT_FilterTypeCurrency},
	})
}

func AdminBlockchainStatesGet(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	efwcontroller.GetModel(
		context,
		blockchainstate.GetModel,
		nil,
		func(id int64) blockchainstate.QueryHelper {
			return blockchainstate.QueryHelper{Id: types.BlockchainStateId(id)}
		},
	)
}

func AdminBlockchainStatesProcess(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	if isMerchantStarted() {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	var input requestAdminBlockchainStatesProcess
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	blockchainStateModel := blockchainstate.GetModel(blockchainstate.QueryHelper{
		Id: input.Id,
	})
	if blockchainStateModel == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	blockchainStateModel.CurrentBlock = input.CurrentBlock
	blockchainstate.Save(blockchainStateModel, blockchainstate.QueryHelper{})

	createLogUser(context, adminUserModel, constants.LogUserAdminBlockchainStateProcess)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": blockchainStateModel})
}

func AdminBlockchainStatesDelete(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	if isMerchantStarted() {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	_, ok := efwcontroller.DeleteModel(
		context,
		blockchainstate.GetModel,
		blockchainstate.Delete,
		func(id int64) blockchainstate.QueryHelper {
			qH := blockchainstate.QueryHelper{Id: types.BlockchainStateId(id)}
			return qH
		},
	)
	if ok {
		createLogUser(context, adminUserModel, constants.LogUserAdminBlockchainStateDelete)
	}
}
