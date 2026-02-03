package btc

import (
	"compay/common/types"

	"github.com/btcsuite/btcd/txscript"
)

func GetTransactions(lastHeight int64) ([]types.TxData, int64, error) {
	currentHeight, err := GetBlockCount()
	if err != nil {
		return nil, lastHeight, err
	}
	var result []types.TxData

	hash, err := client.GetBlockHash(lastHeight)
	if err != nil {
		return nil, lastHeight, err
	}

	block, err := client.GetBlock(hash)
	if err != nil {
		return nil, lastHeight, err
	}

	confirmations := int64(0)
	if currentHeight >= lastHeight {
		confirmations = currentHeight - lastHeight + 1
	}

	timeBlock := block.Header.Timestamp.Unix()

	for _, tx := range block.Transactions {
		txid := tx.TxID()

		for voutIndex, txOut := range tx.TxOut {
			_, addrs, _, err := txscript.ExtractPkScriptAddrs(
				txOut.PkScript,
				networkConfig,
			)
			if err != nil || len(addrs) == 0 {
				continue
			}

			for _, addr := range addrs {
				address := addr.EncodeAddress()

				result = append(result, types.TxData{
					TxId:          types.PaymentTxId(txid),
					TxHash:        types.PaymentTxHash(txid),
					Address:       types.PaymentAddress(address),
					BlockHash:     hash.String(),
					BlockHeight:   lastHeight,
					Time:          timeBlock,
					Amount:        types.Amount(txOut.Value),
					Confirmations: types.Confirmation(confirmations),
					VoutPosition:  voutIndex,
					AssetId:       "BTC", // заполняем просто так
					SentTo:        []string{address},
				})
			}
		}
	}

	return result, currentHeight, nil
}
