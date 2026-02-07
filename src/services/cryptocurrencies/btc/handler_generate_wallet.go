package btc

import (
	"compay/common/cperrors"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
)

func GenerateWallet() (string, string, string, string, string, error) {
	// Генерируем приватный ключ
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		return "", "", "", "", "", cperrors.ErrorBtcFailedGeneratePrivateKey(err.Error())
	}

	// Конвертируем privateKey в WIF формат
	wif, err := btcutil.NewWIF(privateKey, networkConfig, true)
	if err != nil {
		return "", "", "", "", "", cperrors.ErrorBtcFailedCreateWif(err.Error())
	}

	// Получаем публичный ключ
	publicKey := privateKey.PubKey()

	// Создаем адрес
	pubKeyHash := btcutil.Hash160(publicKey.SerializeCompressed())
	address, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, networkConfig)
	if err != nil {
		return "", "", "", "", "", cperrors.ErrorBtcFailedCreateAddress(err.Error())
	}

	return address.EncodeAddress(),
		fmt.Sprintf("%x", privateKey.Serialize()),
		wif.String(),
		fmt.Sprintf("%x", publicKey.SerializeCompressed()),
		fmt.Sprintf("%x", publicKey.SerializeUncompressed()),
		nil
}
