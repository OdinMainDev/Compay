package constants

import "compay/common/types"

// Payments
const (
	PaymentStatusCreated           types.PaymentStatus = "Created"
	PaymentStatusInsufficientFunds types.PaymentStatus = "InsufficientFunds"
	PaymentStatusFundsReceived     types.PaymentStatus = "FundsReceived"
	PaymentStatusCompleted         types.PaymentStatus = "Completed"
	PaymentStatusExpired           types.PaymentStatus = "Expired"
)
