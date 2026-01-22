package constants

import (
	"compay/common/types"
	"fmt"
)

// Actions
const (
	NotificationActionTest types.NotificationAction = "Test"
)

// Notifications
var NotificationTextTest = func(some string) types.NotificationText {
	return types.NotificationText(fmt.Sprintf("%s %s", "Test", some))
}
