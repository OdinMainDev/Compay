package types

// Merchant Types
type MerchantId int64
type MerchantCurrencyId int64
type MerchantWalletId int64

// Notification Types
type NotificationId int64
type NotificationAction string
type NotificationText string

// Currency Types
type CurrencyId int64
type CurrencyNetwork string
type CurrencyIsCoin bool

// Blockchain
type BlockchainStateId int64

// Payment Types
type PaymentId int64
type PaymentStatus string
type PaymentAddress string
type PaymentTxId string
type PaymentTxHash string

// Withdraw Types
type WithdrawId int64
type WithdrawAddress string
type WithdrawTxId string
type WithdrawTxHash string

// Log Types
type LogUserId int64
type LogUserAction string
type LogAuthId int64
type LogAuthAction string
type LogPaymentId int64
type LogMerchantWalletId int64
