package btc

import (
	"bytes"
	"compay/common/cperrors"
	"efw/components/entity"
	"efw/core/efwamount"
	"encoding/hex"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

// WithdrawResponse содержит результат вывода средств
type WithdrawResponse struct {
	TxHash       string        `json:"tx_hash"`
	TotalInput   entity.Amount `json:"total_input"`
	TotalOutput  entity.Amount `json:"total_output"`
	Fee          entity.Amount `json:"fee"`
	ChangeAmount entity.Amount `json:"change_amount"`
}

// StoredUTXO представляет UTXO из базы данных
type StoredUTXO struct {
	TxID   string `json:"tx_id"`
	Vout   uint32 `json:"vout_position"`
	Amount int64  `json:"amount"` // в сатоши
}

// Withdraw выводит средства с адреса
// storedUTXOs - список UTXO из базы данных (неизрасходованные входящие транзакции)
func Withdraw(wifStr string, toAddress string, estimatedFeeBlockCount int64, amount btcutil.Amount, storedUTXOs []StoredUTXO) (*WithdrawResponse, error) {
	// Декодируем WIF приватный ключ
	wif, err := btcutil.DecodeWIF(wifStr)
	if err != nil {
		return nil, cperrors.ErrorBtcWithdrawFailedDecodePrivateWif(err.Error())
	}

	// Получаем публичный ключ и адрес отправителя
	publicKey := wif.PrivKey.PubKey()
	pubKeyHash := btcutil.Hash160(publicKey.SerializeCompressed())
	fromAddr, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, networkConfig)
	if err != nil {
		return nil, cperrors.ErrorBtcWithdrawFailedCreateNewAddress(err.Error())
	}

	// Проверяем наличие UTXO
	if len(storedUTXOs) == 0 {
		return nil, cperrors.ErrorBtcWithdrawFailedBadResponseListUnspent(fromAddr.EncodeAddress())
	}

	// Генерируем scriptPubKey из адреса отправителя
	pkScript, err := txscript.PayToAddrScript(fromAddr)
	if err != nil {
		return nil, cperrors.ErrorBtcWithdrawFailedToCreateDestinationScript(err.Error())
	}
	pkScriptHex := hex.EncodeToString(pkScript)

	// Конвертируем StoredUTXO в формат btcjson.ListUnspentResult
	utxos := make([]btcjson.ListUnspentResult, 0, len(storedUTXOs))
	for _, u := range storedUTXOs {
		utxos = append(utxos, btcjson.ListUnspentResult{
			TxID:         u.TxID,
			Vout:         u.Vout,
			Amount:       float64(u.Amount) / 1e8, // сатоши -> BTC
			ScriptPubKey: pkScriptHex,
		})
	}

	// Получаем оптимальную ставку комиссии из Bitcoin Core
	feeResult, err := client.EstimateSmartFee(estimatedFeeBlockCount, &btcjson.EstimateModeConservative)
	if err != nil {
		return nil, cperrors.ErrorBtcWithdrawFailedGetEstimatedFee(err.Error())
	}

	// Если нет данных для оценки, используем fallback значение
	var feeRatePerKB btcutil.Amount
	if feeResult.FeeRate != nil && *feeResult.FeeRate > 0 {
		// FeeRate в BTC/KB, конвертируем в satoshi/KB
		feeRatePerKB = btcutil.Amount(efwamount.ToAmount(*feeResult.FeeRate, 8))
	} else {
		// Fallback: 10 satoshi/byte = 10000 satoshi/KB (средняя оптимальная комиссия)
		feeRatePerKB = 10000
	}

	// Собираем входы для транзакции
	var selectedUTXOs []btcjson.ListUnspentResult
	var totalInput btcutil.Amount
	var estimatedFee btcutil.Amount

	for _, utxo := range utxos {
		selectedUTXOs = append(selectedUTXOs, utxo)
		// utxo.Amount уже в BTC, конвертируем в satoshi
		totalInput += btcutil.Amount(efwamount.ToAmount(utxo.Amount, 8))

		// Предварительная оценка размера транзакции
		// SegWit вход: ~68 vbytes, выход: ~31 vbytes, overhead: ~10 vbytes
		estimatedSize := len(selectedUTXOs)*68 + 2*31 + 10 // 2 выхода (получатель + сдача)
		tempFee := btcutil.Amount(estimatedSize) * feeRatePerKB / 1000

		// Проверяем, достаточно ли средств с учетом комиссии
		if totalInput >= amount+tempFee {
			break
		}
	}

	// Финальный расчет комиссии с учетом выбранных входов
	txSize := len(selectedUTXOs)*68 + 2*31 + 10
	estimatedFee = btcutil.Amount(txSize) * feeRatePerKB / 1000

	// Проверяем достаточность средств
	if totalInput < amount {
		return nil, cperrors.ErrorBtcWithdrawInsufficientFunds(int64(totalInput), int64(amount))
	}

	// Проверяем что после комиссии хватит средств
	if totalInput < amount+estimatedFee {
		return nil, cperrors.ErrorBtcWithdrawInsufficientFundsForFee(int64(totalInput), int64(amount), int64(estimatedFee))
	}

	// Создаем новую транзакцию
	tx := wire.NewMsgTx(wire.TxVersion)

	// Добавляем входы
	for _, utxo := range selectedUTXOs {
		txHash, err := chainhash.NewHashFromStr(utxo.TxID)
		if err != nil {
			return nil, cperrors.ErrorBtcWithdrawFailedToParseTxId(err.Error())
		}

		outPoint := wire.NewOutPoint(txHash, utxo.Vout)
		// избегаем возможности отменить выплату
		txIn := wire.NewTxIn(outPoint, nil, nil)
		tx.AddTxIn(txIn)
	}

	// Декодируем адрес получателя
	destAddr, err := btcutil.DecodeAddress(toAddress, networkConfig)
	if err != nil {
		return nil, cperrors.ErrorBtcWithdrawFailedToDecodeDesitnationAddress(err.Error())
	}

	// Создаем скрипт для получателя
	destPkScript, err := txscript.PayToAddrScript(destAddr)
	if err != nil {
		return nil, cperrors.ErrorBtcWithdrawFailedToCreateDestinationScript(err.Error())
	}

	// Добавляем выход для получателя
	tx.AddTxOut(wire.NewTxOut(int64(amount), destPkScript))

	// Рассчитываем сдачу
	change := totalInput - amount - estimatedFee

	// Финальная комиссия (может измениться если нет сдачи)
	finalFee := estimatedFee

	// Если есть сдача, добавляем выход для сдачи
	// Минимальная сдача должна быть больше dust limit (~546 satoshi)
	if change > 546 {
		changePkScript, err := txscript.PayToAddrScript(fromAddr)
		if err != nil {
			return nil, cperrors.ErrorBtcWithdrawFailedToCreateChangeScript(err.Error())
		}
		tx.AddTxOut(wire.NewTxOut(int64(change), changePkScript))
	} else if change > 0 {
		// Если сдача меньше dust limit, добавляем её к комиссии
		// Пересчитываем размер транзакции (теперь только 1 выход вместо 2)
		txSizeWithoutChange := len(selectedUTXOs)*68 + 1*31 + 10
		finalFee = btcutil.Amount(txSizeWithoutChange) * feeRatePerKB / 1000
		// Остаток идет в комиссию
		finalFee += change
		change = 0
	}

	// Подписываем каждый вход
	for i, utxo := range selectedUTXOs {
		// Получаем скрипт предыдущего выхода
		prevPkScript, err := hex.DecodeString(utxo.ScriptPubKey)
		if err != nil {
			return nil, cperrors.ErrorBtcWithdrawFailedToDecodeScriptPublicKey(err.Error())
		}

		// Создаем witness для SegWit транзакции
		witnessScript, err := txscript.WitnessSignature(
			tx,
			txscript.NewTxSigHashes(tx, txscript.NewCannedPrevOutputFetcher(prevPkScript, int64(utxo.Amount*1e8))),
			i,
			int64(utxo.Amount*1e8),
			prevPkScript,
			txscript.SigHashAll,
			wif.PrivKey,
			true,
		)
		if err != nil {
			return nil, cperrors.ErrorBtcWithdrawFailedToCreateWitnessSignature(err.Error())
		}

		tx.TxIn[i].Witness = witnessScript
	}

	// Сериализуем транзакцию
	var buf bytes.Buffer
	if err := tx.Serialize(&buf); err != nil {
		return nil, cperrors.ErrorBtcWithdrawFailedToSerializeTransaction(err.Error())
	}

	// Отправляем транзакцию
	txHash, err := client.SendRawTransaction(tx, false)
	if err != nil {
		return nil, cperrors.ErrorBtcWithdrawFailedToSendTransaction(err.Error())
	}

	return &WithdrawResponse{
		TxHash:       txHash.String(),
		TotalInput:   entity.Amount(totalInput),
		TotalOutput:  entity.Amount(amount + change),
		Fee:          entity.Amount(finalFee),
		ChangeAmount: entity.Amount(change),
	}, nil
}
