package btc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// BTCNodeClient представляет клиент для взаимодействия с Bitcoin нодой
type BTCNodeClient struct {
	nodeURL  string
	username string
	password string
	client   *http.Client
}

// NewBTCNodeClient создает новый экземпляр клиента для Bitcoin ноды
func NewBTCNodeClient(nodeURL, username, password string) *BTCNodeClient {
	return &BTCNodeClient{
		nodeURL:  nodeURL,
		username: username,
		password: password,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// JSONRPCRequest структура для JSON-RPC запроса
type JSONRPCRequest struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
}

// JSONRPCResponse структура для JSON-RPC ответа
type JSONRPCResponse struct {
	Result json.RawMessage `json:"result"`
	Error  *RPCError       `json:"error"`
	ID     string          `json:"id"`
}

// RPCError структура для ошибок RPC
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// callRPC выполняет RPC вызов к Bitcoin ноде
func (c *BTCNodeClient) callRPC(method string, params []any) (json.RawMessage, error) {
	requestBody := JSONRPCRequest{
		JSONRPC: "1.0",
		ID:      "compay",
		Method:  method,
		Params:  params,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", c.nodeURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var rpcResp JSONRPCResponse
	if err := json.Unmarshal(body, &rpcResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if rpcResp.Error != nil {
		return nil, fmt.Errorf("RPC error %d: %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}

	return rpcResp.Result, nil
}

// Transaction представляет транзакцию Bitcoin
type Transaction struct {
	TxID          string  `json:"txid"`
	Amount        float64 `json:"amount"`
	Confirmations int     `json:"confirmations"`
	Address       string  `json:"address"`
	Time          int64   `json:"time"`
	BlockTime     int64   `json:"blocktime"`
}

// ReceivedByAddress представляет информацию о полученных средствах на адрес
type ReceivedByAddress struct {
	Address       string   `json:"address"`
	Amount        float64  `json:"amount"`
	Confirmations int      `json:"confirmations"`
	TxIDs         []string `json:"txids"`
}

// GetReceivedByAddress получает информацию о полученных средствах на указанный адрес
func (c *BTCNodeClient) GetReceivedByAddress(address string, minConfirmations int) (float64, error) {
	result, err := c.callRPC("getreceivedbyaddress", []any{address, minConfirmations})
	if err != nil {
		return 0, err
	}

	var amount float64
	if err := json.Unmarshal(result, &amount); err != nil {
		return 0, fmt.Errorf("failed to unmarshal amount: %w", err)
	}

	return amount, nil
}

// ListReceivedByAddress получает список полученных средств по адресам
func (c *BTCNodeClient) ListReceivedByAddress(minConfirmations int, includeEmpty bool, includeWatchOnly bool) ([]ReceivedByAddress, error) {
	result, err := c.callRPC("listreceivedbyaddress", []any{minConfirmations, includeEmpty, includeWatchOnly})
	if err != nil {
		return nil, err
	}

	var received []ReceivedByAddress
	if err := json.Unmarshal(result, &received); err != nil {
		return nil, fmt.Errorf("failed to unmarshal received list: %w", err)
	}

	return received, nil
}

// ListTransactions получает список транзакций
func (c *BTCNodeClient) ListTransactions(account string, count int, skip int) ([]Transaction, error) {
	result, err := c.callRPC("listtransactions", []any{account, count, skip})
	if err != nil {
		return nil, err
	}

	var transactions []Transaction
	if err := json.Unmarshal(result, &transactions); err != nil {
		return nil, fmt.Errorf("failed to unmarshal transactions: %w", err)
	}

	return transactions, nil
}

// GetTransaction получает информацию о конкретной транзакции
func (c *BTCNodeClient) GetTransaction(txid string) (*TransactionDetail, error) {
	result, err := c.callRPC("gettransaction", []any{txid})
	if err != nil {
		return nil, err
	}

	var txDetail TransactionDetail
	if err := json.Unmarshal(result, &txDetail); err != nil {
		return nil, fmt.Errorf("failed to unmarshal transaction detail: %w", err)
	}

	return &txDetail, nil
}

// TransactionDetail детальная информация о транзакции
type TransactionDetail struct {
	Amount        float64             `json:"amount"`
	Fee           float64             `json:"fee"`
	Confirmations int                 `json:"confirmations"`
	BlockHash     string              `json:"blockhash"`
	BlockIndex    int                 `json:"blockindex"`
	BlockTime     int64               `json:"blocktime"`
	TxID          string              `json:"txid"`
	Time          int64               `json:"time"`
	TimeReceived  int64               `json:"timereceived"`
	Details       []TransactionOutput `json:"details"`
	Hex           string              `json:"hex"`
}

// TransactionOutput представляет выход транзакции
type TransactionOutput struct {
	Address  string  `json:"address"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
	Vout     int     `json:"vout"`
}

// ImportAddress импортирует адрес в кошелек для отслеживания
func (c *BTCNodeClient) ImportAddress(address string, label string, rescan bool) error {
	_, err := c.callRPC("importaddress", []any{address, label, rescan})
	return err
}

// GetBlockCount получает текущую высоту блокчейна
func (c *BTCNodeClient) GetBlockCount() (int, error) {
	result, err := c.callRPC("getblockcount", []any{})
	if err != nil {
		return 0, err
	}

	var blockCount int
	if err := json.Unmarshal(result, &blockCount); err != nil {
		return 0, fmt.Errorf("failed to unmarshal block count: %w", err)
	}

	return blockCount, nil
}

// GetBalance получает баланс кошелька
func (c *BTCNodeClient) GetBalance() (float64, error) {
	result, err := c.callRPC("getbalance", []any{})
	if err != nil {
		return 0, err
	}

	var balance float64
	if err := json.Unmarshal(result, &balance); err != nil {
		return 0, fmt.Errorf("failed to unmarshal balance: %w", err)
	}

	return balance, nil
}
