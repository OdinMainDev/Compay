package startup

import (
	"compay/common/config"
	"compay/src/helpers"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
)

func FillConfig() {
	config.AppPort = os.Getenv("APP_PORT")

	// Notifications
	if value := os.Getenv("NOTIFY_TELEGRAM_CHAT_DEBUG_ID"); value != "" {
		if parsed, err := strconv.ParseInt(value, 10, 64); err == nil {
			config.NotificationChatDebug = parsed
		}
	}
	if value := os.Getenv("NOTIFY_TELEGRAM_CHAT_LOGS_ID"); value != "" {
		if parsed, err := strconv.ParseInt(value, 10, 64); err == nil {
			config.NotificationChatLogs = parsed
		}
	}
	if value := os.Getenv("NOTIFY_TELEGRAM_CHAT_ERRORS_ID"); value != "" {
		if parsed, err := strconv.ParseInt(value, 10, 64); err == nil {
			config.NotificationChatError = parsed
		}
	}
	if value := os.Getenv("NOTIFY_TELEGRAM_CHAT_WITHDRAWS_ID"); value != "" {
		if parsed, err := strconv.ParseInt(value, 10, 64); err == nil {
			config.NotificationChatWithdraws = parsed
		}
	}

	// Monitoring
	config.MonitoringLogs = os.Getenv("MONITORING_LOGS") == "true"

	// Bitcoin RPC Configuration
	config.BtcApiUrl = os.Getenv("BTC_API_URL")
	config.BtcNodeHost = os.Getenv("BTC_RPC_HOST")
	config.BtcNodeUsername = os.Getenv("BTC_RPC_USERNAME")
	config.BtcNodePassword = os.Getenv("BTC_RPC_PASSWORD")
	config.BtcNetworkMainnet = os.Getenv("BTC_NETWORK_MAINNET") == "true"
	config.BtcDisableTLS = os.Getenv("BTC_DISABLE_TLS") == "true"
}

func FillMerchantConfig() {
	config.AppInitKey = uuid.NewString()
	fmt.Println("TEST SUCC")

	helpers.NotificationDebug("<code>" + config.AppInitKey + "</code>")
}
