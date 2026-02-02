package btc

// DEPRECATED

// func GetTransactionConfirmations(txHash string) (*types.TxData, error) {
// 	hash, err := chainhash.NewHashFromStr(txHash)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Получаем raw транзакцию
// 	tx, err := client.GetRawTransactionVerbose(hash)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Если транзакция не в блоке — confirmations = 0
// 	var confirmations int64 = 0
// 	var blockHash string
// 	var blockHeight int64 = 0
// 	var blockTime int64 = 0

// 	if tx.BlockHash != "" {
// 		blockHash = tx.BlockHash

// 		blockHashParsed, err := chainhash.NewHashFromStr(tx.BlockHash)
// 		if err != nil {
// 			return nil, err
// 		}

// 		block, err := client.GetBlockVerbose(blockHashParsed)
// 		if err != nil {
// 			return nil, err
// 		}

// 		currentHeight, err := GetBlockCount()
// 		if err != nil {
// 			return nil, err
// 		}

// 		blockHeight = block.Height
// 		blockTime = block.Time
// 		confirmations = currentHeight - int64(block.Height) + 1
// 	}

// 	return &types.TxData{
// 		TxId:          tx.Txid,
// 		TxHash:        tx.Hash,
// 		BlockHash:     blockHash,
// 		BlockHeight:   blockHeight,
// 		Time:          blockTime,
// 		Confirmations: types.Confirmation(confirmations),
// 		AssetId:       "BTC",
// 	}, nil
// }
