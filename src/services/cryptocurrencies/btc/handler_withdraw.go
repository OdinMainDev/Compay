package btc

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

// WithdrawResponse содержит результат вывода средств
type WithdrawResponse struct {
	TxHash       string         `json:"tx_hash"`
	TotalInput   btcutil.Amount `json:"total_input"`
	TotalOutput  btcutil.Amount `json:"total_output"`
	Fee          btcutil.Amount `json:"fee"`
	ChangeAmount btcutil.Amount `json:"change_amount"`
}

// Withdraw выводит средства с адреса
func Withdraw(
	client *rpcclient.Client,
	wifStr string,
	toAddress string,
	amount btcutil.Amount,
	network *chaincfg.Params,
) (*WithdrawResponse, error) {
	// Декодируем WIF приватный ключ
	wif, err := btcutil.DecodeWIF(wifStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode WIF: %w", err)
	}

	// Получаем публичный ключ и адрес отправителя
	pubKey := wif.PrivKey.PubKey()
	pubKeyHash := btcutil.Hash160(pubKey.SerializeCompressed())
	fromAddr, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, network)
	if err != nil {
		return nil, fmt.Errorf("failed to create address: %w", err)
	}

	// Получаем UTXO для адреса
	utxos, err := client.ListUnspentMinMaxAddresses(1, 9999999, []btcutil.Address{fromAddr})
	if err != nil {
		return nil, fmt.Errorf("failed to list unspent: %w", err)
	}

	if len(utxos) == 0 {
		return nil, fmt.Errorf("no UTXOs found for address %s", fromAddr.EncodeAddress())
	}

	// Собираем входы для транзакции
	var selectedUTXOs []btcjson.ListUnspentResult
	var totalInput btcutil.Amount

	for _, utxo := range utxos {
		selectedUTXOs = append(selectedUTXOs, utxo)
		totalInput += btcutil.Amount(utxo.Amount * 1e8)
		if totalInput >= amount {
			break
		}
	}

	// Проверяем достаточность средств
	if totalInput < amount {
		return nil, fmt.Errorf("insufficient funds: have %d satoshi, need %d satoshi", totalInput, amount)
	}

	// Оцениваем комиссию (примерно 1000 satoshi за вход + 1000 за выход)
	estimatedFee := btcutil.Amount(len(selectedUTXOs)*1000 + 2000)

	// Проверяем что после комиссии хватит средств
	if totalInput < amount+estimatedFee {
		return nil, fmt.Errorf("insufficient funds for fee: have %d satoshi, need %d satoshi (amount) + %d satoshi (fee)",
			totalInput, amount, estimatedFee)
	}

	// Создаем новую транзакцию
	tx := wire.NewMsgTx(wire.TxVersion)

	// Добавляем входы
	for _, utxo := range selectedUTXOs {
		txHash, err := chainhash.NewHashFromStr(utxo.TxID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse txid: %w", err)
		}

		outPoint := wire.NewOutPoint(txHash, utxo.Vout)
		txIn := wire.NewTxIn(outPoint, nil, nil)
		tx.AddTxIn(txIn)
	}

	// Декодируем адрес получателя
	destAddr, err := btcutil.DecodeAddress(toAddress, network)
	if err != nil {
		return nil, fmt.Errorf("failed to decode destination address: %w", err)
	}

	// Создаем скрипт для получателя
	destPkScript, err := txscript.PayToAddrScript(destAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to create destination script: %w", err)
	}

	// Добавляем выход для получателя
	tx.AddTxOut(wire.NewTxOut(int64(amount), destPkScript))

	// Рассчитываем сдачу
	change := totalInput - amount - estimatedFee

	// Если есть сдача, добавляем выход для сдачи
	if change > 0 {
		changePkScript, err := txscript.PayToAddrScript(fromAddr)
		if err != nil {
			return nil, fmt.Errorf("failed to create change script: %w", err)
		}
		tx.AddTxOut(wire.NewTxOut(int64(change), changePkScript))
	}

	// Подписываем каждый вход
	for i, utxo := range selectedUTXOs {
		// Получаем скрипт предыдущего выхода
		prevPkScript, err := hex.DecodeString(utxo.ScriptPubKey)
		if err != nil {
			return nil, fmt.Errorf("failed to decode scriptPubKey: %w", err)
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
			return nil, fmt.Errorf("failed to create witness signature: %w", err)
		}

		tx.TxIn[i].Witness = witnessScript
	}

	// Сериализуем транзакцию
	var buf bytes.Buffer
	if err := tx.Serialize(&buf); err != nil {
		return nil, fmt.Errorf("failed to serialize transaction: %w", err)
	}

	// Отправляем транзакцию
	txHash, err := client.SendRawTransaction(tx, false)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	return &WithdrawResponse{
		TxHash:       txHash.String(),
		TotalInput:   totalInput,
		TotalOutput:  amount + change,
		Fee:          estimatedFee,
		ChangeAmount: change,
	}, nil
}
