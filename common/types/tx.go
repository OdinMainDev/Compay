package types

type TxData struct {
	TxId          PaymentTxId    `json:"tx_id"`         // tx id ( у бтц совпадает с хеш )
	TxHash        PaymentTxHash  `json:"tx_hash"`       // tx hash
	Address       PaymentAddress `json:"address"`       // адрес кошелька
	BlockHash     string         `json:"block_hash"`    // block hash
	BlockHeight   int64          `json:"block_height"`  // block height
	Time          int64          `json:"time"`          // время ( обычно время создания блока )
	Amount        Amount         `json:"amount"`        // сумма транзакции ( получение/отрпавка)
	Confirmations Confirmation   `json:"confirmations"` // кол-во конфирмов
	VoutPosition  int            `json:"vout_position"` // vout позиция ( # адреса в транзакции)
	Nonce         string         `json:"nonce"`         // порядковый номер операции аккаунта
	Memo          string         `json:"memo"`          // комментарий/полезная нагрузка
	AssetId       string         `json:"asset_id"`      // идентификатор актива
	SentTo        []string       `json:"sent_to"`       // кому пришли

	// ReceivedFrom []string `json:"received_from"` // от кого пришли (tx/address) DEPRECATED
	// Fee          float64  `json:"fee"`           // комса транзы DEPRECATED
}
