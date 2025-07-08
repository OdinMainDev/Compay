package cryptowallet

import (
	"compay/application/entities/basemodel"
	"compay/application/entities/wallet"
	"compay/common/cperrors"
	"fmt"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func CreateWallet(userId int64) (*wallet.Model, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return nil, cperrors.ErrorWalletCreateError
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, cperrors.ErrorWalletCreateError2
	}

	seed := bip39.NewSeed(mnemonic, "")

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, cperrors.ErrorWalletCreateMasterKey
	}

	masterPrivateKey := fmt.Sprintf("%x", masterKey.Key)

	walletModel := &wallet.Model{
		UserId:     userId,
		Mnemonic:   basemodel.EncryptedString(mnemonic),
		PrivateKey: basemodel.EncryptedString(masterPrivateKey),
	}

	walletModel.Save(walletModel)
	return walletModel, nil
}
