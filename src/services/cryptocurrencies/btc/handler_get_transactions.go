package btc

import (
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
)

type TxOutput struct {
	TxId    string
	Vout    uint32
	Address string
	Amount  btcutil.Amount
	Block   int64
}

func GetTransactions(
	client *rpcclient.Client,
	lastHeight int64,
	network *chaincfg.Params,
) ([]TxOutput, int64, error) {

	currentHeight, err := client.GetBlockCount()
	if err != nil {
		return nil, lastHeight, err
	}

	var result []TxOutput

	for h := lastHeight + 1; h <= currentHeight; h++ {
		hash, err := client.GetBlockHash(h)
		if err != nil {
			return nil, lastHeight, err
		}

		block, err := client.GetBlockVerboseTx(hash)
		if err != nil {
			return nil, lastHeight, err
		}

		for _, tx := range block.Tx {
			for _, vout := range tx.Vout {
				if vout.ScriptPubKey.Addresses == nil {
					continue
				}

				for _, addr := range vout.ScriptPubKey.Addresses {
					amount, _ := btcutil.NewAmount(vout.Value)
					result = append(result, TxOutput{
						TxId:    tx.Txid,
						Vout:    uint32(vout.N),
						Address: addr,
						Amount:  amount,
						Block:   h,
					})
				}
			}
		}
	}

	return result, currentHeight, nil
}
