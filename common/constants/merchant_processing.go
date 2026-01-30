package constants

import "compay/common/types"

// Payments
const (
	PaymentStatusCreated           types.PaymentStatus = "Created"
	PaymentStatusInsufficientFunds types.PaymentStatus = "InsufficientFunds"
	PaymentStatusCompleted         types.PaymentStatus = "Completed"
	PaymentStatusExpired           types.PaymentStatus = "Expired"
)
