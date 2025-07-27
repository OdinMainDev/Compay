package start

import (
	"compay/common/config"
	configCore "compay/core/config"
	"os"
)

func fillAppConfig() {
	config.AppPort = os.Getenv("APP_PORT")
	config.AppDebug = os.Getenv("APP_DEBUG") == "true"
	config.AppDomain = os.Getenv("APP_DOMAIN")
	config.AppTelegramToken = os.Getenv("APP_TELEGRAM_TOKEN")
	config.AppAdminDomain = os.Getenv("APP_ADMIN_DOMAIN")

	config.AppSecret = generateSecret(256)

	configCore.ClamavSkipValidation = true
	configCore.CdnUploaderDisabled = true
}

// TODO: потом вернуть генерацию секретки
func generateSecret(size int) string {
	return "ONLY FOR TESTS"
	// bytes := make([]byte, size)
	// if _, err := rand.Read(bytes); err != nil {
	// 	panic(fmt.Sprintf("error while generating app secret: %s", err))
	// }
	// return base64.URLEncoding.EncodeToString(bytes)
}
