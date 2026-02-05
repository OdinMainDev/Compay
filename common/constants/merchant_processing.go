package constants

import "compay/common/types"

// Payments
const (
	PaymentStatusCreated           types.PaymentStatus = "Created"
	PaymentStatusFundsDetected     types.PaymentStatus = "FundsDetected"
	PaymentStatusFundsReceived     types.PaymentStatus = "FundsReceived"
	PaymentStatusCompleted         types.PaymentStatus = "Completed"
	PaymentStatusInsufficientFunds types.PaymentStatus = "InsufficientFunds" // ошибка (недостаточно средств отправили, и по истечению времени заявка умерла)
	PaymentStatusExpired           types.PaymentStatus = "Expired"
)

var PaymentStatuses = []types.PaymentStatus{
	PaymentStatusCreated,
	PaymentStatusFundsDetected,
	PaymentStatusFundsReceived,
	PaymentStatusCompleted,
	PaymentStatusInsufficientFunds,
	PaymentStatusExpired,
}
