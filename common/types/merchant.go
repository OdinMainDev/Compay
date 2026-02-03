package types

// Merchant Types
type MerchantId int64
type MerchantCurrencyId int64

// Notification Types
type NotificationId int64
type NotificationAction string
type NotificationText string

// Currency Types
type CurrencyId int64
type CurrencyCodeAlpha string
type CurrencyNetwork string
type CurrencyIsCoin bool
type CurrencyMinorUnits int

// Blockchain
type BlockchainStateId int64

// Payment Types
type PaymentId int64
type PaymentStatus string
type PaymentAddress string
type PaymentTxId string
type PaymentTxHash string

// Log Types
type LogUserId int64
type LogUserAction string
type LogAuthId int64
type LogAuthAction string
type LogPaymentId int64
