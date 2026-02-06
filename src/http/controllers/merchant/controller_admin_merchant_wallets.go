package merchant

import (
	"compay/common/types"
	"compay/src/models/merchantwallet"
	"efw/services/efwcontroller"
	"efw/services/efwdt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type requestAdminMerchantWalletsList struct {
	Id types.PaymentId `json:"id" efwdt:"id" binding:"omitempty,gt=0"`
	efwdt.WithCreatedAt
}

func AdminMerchantWalletsList(context *gin.Context) {
	efwcontroller.Datable[merchantwallet.Model, requestAdminMerchantWalletsList](
		context,
		[]string{"id", "created_at"},
		[]efwdt.ChainFilter[*merchantwallet.Model]{
			func(chain gorm.ChainInterface[*merchantwallet.Model]) gorm.ChainInterface[*merchantwallet.Model] {
				return chain.Preload("Payment", func(db gorm.PreloadBuilder) error { return nil })
			},
		},
		nil,
		merchantwallet.DTO,
	)
}

func AdminMerchantWalletsListFilters(context *gin.Context) {
	efwcontroller.DatatableFilters(context, []efwdt.Filter{
		{Name: "id", Type: types.EFWDT_FilterTypeInteger},
		{Name: "created_at", Type: types.EFWDT_FilterTypeDate},
	})
}

func AdminMerchantWalletsGet(context *gin.Context) {
	// adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	efwcontroller.GetModel(
		context,
		merchantwallet.GetModel,
		merchantwallet.DTO,
		func(id int64) merchantwallet.QueryHelper {
			qH := merchantwallet.QueryHelper{Id: types.MerchantWalletId(id)}
			return qH
		},
	)
}
