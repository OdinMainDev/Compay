package monitoring

import "compay/common/types"

type baseResponse struct {
	Status  int                 `json:"status"`
	Success bool                `json:"success"`
	Error   string              `json:"error"`
	Errors  map[string][]string `json:"errors"`
}

type getBlockCountBtcResponse struct {
	baseResponse
	Data int64 `json:"data"`
}

type getTransactionsBtcRequest struct {
	LastHeight int64 `json:"last_height"`
}

type getTransactionsBtcResponse struct {
	baseResponse
	Data getTransactionsBtcResponseData `json:"data"`
}

type getTransactionsBtcResponseData struct {
	Transactions []types.TxData `json:"transactions"`
	NewHeight    int64          `json:"new_height"`
}
