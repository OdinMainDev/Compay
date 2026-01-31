package btc

func GetBlockCount() (int64, error) {
	return client.GetBlockCount()
}
