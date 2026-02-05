package merchant

import (
	"compay/common/types"
	"compay/src/services/cconverter"
	"compay/src/services/processing"
	"compay/src/services/startup"
	"efw/config"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestAppInit struct {
	Cryptiv   string `json:"crypt_iv" binding:"required,len=32"`
	Cryptkey  string `json:"crypt_key" binding:"required,len=32"`
	CryptSalt string `json:"crypt_salt" binding:"omitempty,max=255"`
	Sign      string `json:"sign" binding:"required,len=128"`
}

func InitApp(context *gin.Context) {
	if config.Cryptkey != "" || config.Cryptiv != "" {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	var input requestAppInit
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	if !signVerify("e73afba1-7a42-4a38-b32c-5f1b9a6b6b98", signData{input.Cryptiv, types.CurrencyCodeAlpha(input.Cryptkey), input.Sign}) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	config.Cryptiv = input.Cryptiv
	config.Cryptkey = input.Cryptkey
	config.CrypSalt = input.CryptSalt
	startup.StartMerchant()
	cconverter.Start()
	processing.StartProcessing()

	efwginhelpers.ResponseSuccess(context, nil)
}
