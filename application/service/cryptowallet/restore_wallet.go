package cryptowallet

import (
	"compay/common/cperrors"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func RestoreWallet(mnemonic string, userId int64) error {
	seed := bip39.NewSeed(mnemonic, "")

	_, err := bip32.NewMasterKey(seed)
	if err != nil {
		return cperrors.ErrorWalletCreateMasterKeyOnRestore
	}

	// masterPrivateKey := fmt.Sprintf("%x", masterKey.Key)

	return nil
}

// func RestoreWallet(mnemonic string, userId int64) (*wallet.Model, error) {
// 	seed := bip39.NewSeed(mnemonic, "")

// 	masterKey, err := bip32.NewMasterKey(seed)
// 	if err != nil {
// 		return nil, cperrors.ErrorWalletCreateMasterKeyOnRestore
// 	}

// 	masterPrivateKey := fmt.Sprintf("%x", masterKey.Key)

// 	walletModel := &wallet.Model{
// 		UserId:     userId,
// 		Mnemonic:   basemodel.EncryptedString(mnemonic),
// 		PrivateKey: basemodel.EncryptedString(masterPrivateKey),
// 	}

// 	walletModel.Save(walletModel)

// 	return walletModel, nil
// }
