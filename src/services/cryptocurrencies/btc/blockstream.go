package btc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
)

// BlockstreamUTXO представляет UTXO из Blockstream API
type BlockstreamUTXO struct {
	TxID   string `json:"txid"`
	Vout   uint32 `json:"vout"`
	Status struct {
		Confirmed   bool  `json:"confirmed"`
		BlockHeight int64 `json:"block_height"`
	} `json:"status"`
	Value int64 `json:"value"` // в сатоши
}

var httpClient = &http.Client{
	Timeout: 30 * time.Second,
}

const maxRetries = 3

// getBlockstreamBaseURL возвращает базовый URL в зависимости от сети
func getBlockstreamBaseURL() string {
	if networkConfig == &chaincfg.MainNetParams {
		return "https://blockstream.info/api"
	}
	// testnet
	return "https://blockstream.info/testnet/api"
}

// GetUTXOsFromBlockstream получает список UTXO для адреса через Blockstream API
func GetUTXOsFromBlockstream(address string) ([]BlockstreamUTXO, error) {
	url := fmt.Sprintf("%s/address/%s/utxo", getBlockstreamBaseURL(), address)

	var lastErr error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		resp, err := httpClient.Get(url)
		if err != nil {
			lastErr = fmt.Errorf("failed to fetch UTXOs (attempt %d/%d): %w", attempt, maxRetries, err)
			time.Sleep(time.Duration(attempt) * time.Second)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			lastErr = fmt.Errorf("blockstream API error: %s, body: %s", resp.Status, string(body))
			time.Sleep(time.Duration(attempt) * time.Second)
			continue
		}

		var utxos []BlockstreamUTXO
		if err := json.NewDecoder(resp.Body).Decode(&utxos); err != nil {
			resp.Body.Close()
			return nil, fmt.Errorf("failed to decode UTXOs: %w", err)
		}
		resp.Body.Close()

		return utxos, nil
	}

	return nil, lastErr
}

// GetConfirmedUTXOsFromBlockstream получает только подтверждённые UTXO
func GetConfirmedUTXOsFromBlockstream(address string, minConfirmations int64) ([]BlockstreamUTXO, error) {
	utxos, err := GetUTXOsFromBlockstream(address)
	if err != nil {
		return nil, err
	}

	// Получаем текущую высоту блока для расчёта подтверждений
	blockCount, err := client.GetBlockCount()
	if err != nil {
		return nil, fmt.Errorf("failed to get block count: %w", err)
	}

	var confirmed []BlockstreamUTXO
	for _, utxo := range utxos {
		if !utxo.Status.Confirmed {
			continue
		}

		confirmations := blockCount - utxo.Status.BlockHeight + 1
		if confirmations >= minConfirmations {
			confirmed = append(confirmed, utxo)
		}
	}

	return confirmed, nil
}
