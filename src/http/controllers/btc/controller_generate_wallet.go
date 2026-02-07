package btc

import (
	"compay/src/services/cryptocurrencies/btc"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
)

func GenerateWallet(context *gin.Context) {
	address, privateKey, privateKeyWif, publicKey, publicKeyFull, err := btc.GenerateWallet()
	if err != nil {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(err))
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": gin.H{
		"address":         address,
		"private_key":     privateKey,
		"private_key_wif": privateKeyWif,
		"public_key":      publicKey,
		"public_key_full": publicKeyFull,
	}})
}
