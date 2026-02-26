package merchant

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/user"
	"compay/src/models/withdraw"
	"efw/components/entity"
	"efw/core/efwgin"
	"efw/services/efwcontroller"
	"efw/services/efwdt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type requestMerchantWithdrawsList struct {
	Id       string                   `json:"id" efwdt:"trx_id" binding:"omitempty,gt=0"`
	Currency entity.CurrencyCodeAlpha `json:"currency" efwdt:"currency" binding:"omitempty,currencyExist"`
	Address  types.PaymentAddress     `json:"address" efwdt:"address" binding:"omitempty,gt=0"`
	efwdt.WithCreatedAt
}

func MerchantWithdrawsList(context *gin.Context) {
	var input efwdt.Request[requestMerchantWithdrawsList]
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	userModel := context.MustGet(constants.GinUser).(*user.Model)

	efwcontroller.Datable[withdraw.Model, requestMerchantWithdrawsList](
		context,
		[]string{"id", "currency", "amount", "created_at"},
		[]efwdt.ChainFilter[*withdraw.Model]{
			func(chain gorm.ChainInterface[*withdraw.Model]) gorm.ChainInterface[*withdraw.Model] {
				return chain.Where("merchant_id = ?", userModel.MerchantId)
			},
		},
		nil,
		withdraw.DTOMerchant,
	)
}

func MerchantWithdrawsListFilters(context *gin.Context) {
	efwcontroller.DatatableFilters(context, []efwdt.Filter{
		{Name: "id", Type: types.EFWDT_FilterTypeText},
		{Name: "currency", Type: types.EFWDT_FilterTypeCurrency},
		{Name: "address", Type: types.EFWDT_FilterTypeText},
		{Name: "created_at", Type: types.EFWDT_FilterTypeDate},
	})
}
