package btc

import (
	"bytes"
	"compay/common/cperrors"
	"compay/common/types"
	"efw/components/entity"
	"efw/core/efwamount"
	"encoding/hex"
	"time"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

// StoredUTXO представляет UTXO из базы данных
type StoredUTXO struct {
	TxID   string `json:"tx_id"`
	Vout   uint32 `json:"vout_position"`
	Amount int64  `json:"amount"` // в сатоши
}

// Withdraw выводит средства с адреса
// storedUTXOs - список UTXO из базы данных (неизрасходованные входящие транзакции)
func Withdraw(wifStr string, toAddress string, estimatedFeeBlockCount int64, amount btcutil.Amount, storedUTXOs []StoredUTXO) (*types.TxWithdrawData, error) {
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

	// Минимальный relay fee в Bitcoin Core: 1 sat/vbyte = 1000 sat/KB
	minRelayFeePerKB := btcutil.Amount(1000)
	if feeRatePerKB < minRelayFeePerKB {
		feeRatePerKB = minRelayFeePerKB
	}

	for _, utxo := range utxos {
		selectedUTXOs = append(selectedUTXOs, utxo)
		// utxo.Amount уже в BTC, конвертируем в satoshi
		totalInput += btcutil.Amount(efwamount.ToAmount(utxo.Amount, 8))

		// Проверяем, достаточно ли средств
		if totalInput >= amount {
			break
		}
	}

	// Финальный расчет комиссии с учетом выбранных входов
	// Более точная оценка размера SegWit P2WPKH транзакции:
	// - Overhead: 10.5 vbytes (version 4 + locktime 4 + marker 0.25 + flag 0.25 + input count 1 + output count 1)
	// - Каждый вход: 68 vbytes (outpoint 36 + scriptSig 1 + sequence 4 + witness ~27)
	// - Каждый выход: 31 vbytes (value 8 + scriptPubKey 23)
	txSize := len(selectedUTXOs)*68 + 2*31 + 11
	estimatedFee = btcutil.Amount(txSize) * feeRatePerKB / 1000

	// Минимальная комиссия = размер транзакции * 1 sat/vbyte (min relay fee)
	// Добавляем небольшой запас для округления
	minFee := btcutil.Amount(txSize + 5)
	if estimatedFee < minFee {
		estimatedFee = minFee
	}

	// Проверяем достаточность средств
	if totalInput < amount {
		return nil, cperrors.ErrorBtcWithdrawInsufficientFunds(int64(totalInput), int64(amount))
	}

	// Проверяем что сумма больше комиссии (иначе получателю нечего отправлять)
	if amount <= estimatedFee {
		return nil, cperrors.ErrorBtcWithdrawAmountTooSmall(int64(amount), int64(estimatedFee))
	}

	// Сумма, которую получит получатель (amount минус комиссия)
	sendAmount := amount - estimatedFee

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

	// Рассчитываем сдачу (totalInput - amount, т.к. amount уже включает комиссию)
	change := totalInput - amount

	// Финальная комиссия (может измениться если нет сдачи)
	finalFee := estimatedFee
	finalSendAmount := sendAmount

	// Если сдача меньше dust limit, добавляем её к комиссии
	if change > 0 && change <= 546 {
		// Пересчитываем размер транзакции (теперь только 1 выход вместо 2)
		txSizeWithoutChange := len(selectedUTXOs)*68 + 1*31 + 11
		finalFee = btcutil.Amount(txSizeWithoutChange) * feeRatePerKB / 1000
		// Минимальная комиссия для транзакции без сдачи
		minFeeWithoutChange := btcutil.Amount(txSizeWithoutChange + 5)
		if finalFee < minFeeWithoutChange {
			finalFee = minFeeWithoutChange
		}
		// Сдача идет в комиссию, но получатель получает больше (экономия на размере транзакции)
		// finalSendAmount = totalInput - finalFee (без сдачи)
		finalSendAmount = totalInput - finalFee
		change = 0
	}

	// Добавляем выход для получателя
	tx.AddTxOut(wire.NewTxOut(int64(finalSendAmount), destPkScript))

	// Если есть сдача больше dust limit, добавляем выход для сдачи
	if change > 546 {
		changePkScript, err := txscript.PayToAddrScript(fromAddr)
		if err != nil {
			return nil, cperrors.ErrorBtcWithdrawFailedToCreateChangeScript(err.Error())
		}
		tx.AddTxOut(wire.NewTxOut(int64(change), changePkScript))
	}

	// Подписываем каждый вход
	for i, utxo := range selectedUTXOs {
		// Получаем скрипт предыдущего выхода
		prevPkScript, err := hex.DecodeString(utxo.ScriptPubKey)
		if err != nil {
			return nil, cperrors.ErrorBtcWithdrawFailedToDecodeScriptPublicKey(err.Error())
		}

		// Создаем witness для SegWit транзакции
		// Используем efwamount.ToAmount для точной конвертации BTC -> satoshi (избегаем потерю точности float64)
		utxoAmountSat := int64(efwamount.ToAmount(utxo.Amount, 8))
		witnessScript, err := txscript.WitnessSignature(
			tx,
			txscript.NewTxSigHashes(tx, txscript.NewCannedPrevOutputFetcher(prevPkScript, utxoAmountSat)),
			i,
			utxoAmountSat,
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

	// Собираем список использованных UTXO
	usedUTXOs := make([]types.UsedUTXO, 0, len(selectedUTXOs))
	for _, utxo := range selectedUTXOs {
		usedUTXOs = append(usedUTXOs, types.UsedUTXO{
			TxID: utxo.TxID,
			Vout: utxo.Vout,
		})
	}

	return &types.TxWithdrawData{
		TxData: types.TxData{
			TxId:          types.PaymentTxId(txHash.String()),
			TxHash:        types.PaymentTxHash(txHash.String()),
			Address:       types.PaymentAddress(fromAddr.EncodeAddress()),
			Time:          time.Now().Unix(),
			Amount:        entity.Amount(finalSendAmount), // сумма, которую получит получатель
			Confirmations: 0,
			VoutPosition:  0,
			AssetId:       "BTC",
			SentTo:        []string{toAddress},
		},
		TotalInput:   entity.Amount(totalInput),
		TotalOutput:  entity.Amount(finalSendAmount + change), // получатель + сдача
		Fee:          entity.Amount(finalFee),
		ChangeAmount: entity.Amount(change),
		UsedUTXOs:    usedUTXOs,
	}, nil
}
