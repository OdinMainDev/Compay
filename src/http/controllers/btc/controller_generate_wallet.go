package btc

import (
	"compay/common/cperrors"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/gin-gonic/gin"
)

func GenerateWallet(context *gin.Context) {
	// Генерируем приватный ключ
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorBtcFailedGeneratePrivateKey(err.Error())))
		return
	}

	// Конвертируем privateKey в WIF формат
	wif, err := btcutil.NewWIF(privateKey, networkConfig, true)
	if err != nil {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorBtcFailedCreateWif(err.Error())))
		return
	}

	// Получаем публичный ключ
	publicKey := privateKey.PubKey()

	// Создаем адрес
	pubKeyHash := btcutil.Hash160(publicKey.SerializeCompressed())
	address, err := btcutil.NewAddressPubKeyHash(pubKeyHash, networkConfig)
	if err != nil {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorBtcFailedCreateAddress(err.Error())))
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": gin.H{
		"address":         address.EncodeAddress(),
		"private_key":     fmt.Sprintf("%x", privateKey.Serialize()),
		"private_key_wif": wif.String(),
		"public_key":      fmt.Sprintf("%x", publicKey.SerializeCompressed()),
		"public_key_full": fmt.Sprintf("%x", publicKey.SerializeUncompressed()),
	}})
}
