package btc

import (
	"efw/components/entity"

	"github.com/btcsuite/btcd/btcutil"
)

func GetAddressBalance(address string) (entity.Amount, error) {
	addr, err := btcutil.DecodeAddress(address, networkConfig)
	if err != nil {
		return 0, err
	}

	// Получаем список неизрасходованных выходов для адреса
	// Используем ListUnspentMinMaxAddresses вместо ScanTxOutSet для быстродействия
	utxos, err := client.ListUnspentMinMaxAddresses(1, 9999999, []btcutil.Address{addr})
	if err != nil {
		return 0, err
	}

	// Суммируем баланс из всех UTXO
	var balance int64
	for _, utxo := range utxos {
		// utxo.Amount уже в BTC, конвертируем в satoshi
		satoshi := int64(utxo.Amount * 1e8)
		balance += satoshi
	}

	return entity.Amount(balance), nil
}
