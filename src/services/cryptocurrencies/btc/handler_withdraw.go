package btc

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
)

func Withdraw(
	client *rpcclient.Client,
	wifStr string,
	toAddress string,
	amount btcutil.Amount,
	network *chaincfg.Params,
) (string, error) {

	wif, err := btcutil.DecodeWIF(wifStr)
	if err != nil {
		return "", err
	}

	pubKey := wif.PrivKey.PubKey()
	pubKeyHash := btcutil.Hash160(pubKey.SerializeCompressed())
	fromAddr, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, network)
	if err != nil {
		return "", err
	}

	utxos, err := client.ListUnspentMinMaxAddresses(1, 9999999, []btcutil.Address{fromAddr})
	if err != nil {
		return "", err
	}

	var inputs []btcjson.TransactionInput
	var total btcutil.Amount

	for _, u := range utxos {
		h, _ := chainhash.NewHashFromStr(u.TxID)
		inputs = append(inputs, btcjson.TransactionInput{
			Txid: h.String(),
			Vout: u.Vout,
		})
		total += btcutil.Amount(u.Amount * 1e8)
		if total >= amount {
			break
		}
	}

	if total < amount {
		return "", fmt.Errorf("insufficient funds")
	}

	destAddr, err := btcutil.DecodeAddress(toAddress, network)
	if err != nil {
		return "", err
	}

	outputs := map[btcutil.Address]btcutil.Amount{
		destAddr: amount,
	}

	rawTx, err := client.CreateRawTransaction(inputs, outputs, nil)
	if err != nil {
		return "", err
	}

	signedTx, err := txscript.SignRawTransaction(
		rawTx,
		txscript.NewTxSigHashes(rawTx),
		txscript.KeyClosure(func(addr btcutil.Address) (*btcec.PrivateKey, bool, error) {
			if addr.EncodeAddress() == fromAddr.EncodeAddress() {
				return wif.PrivKey, true, nil
			}
			return nil, false, nil
		}),
		nil,
	)
	if err != nil {
		return "", err
	}

	txHash, err := client.SendRawTransaction(signedTx, false)
	if err != nil {
		return "", err
	}

	return txHash.String(), nil
}
