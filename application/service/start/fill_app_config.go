package start

import (
	"compay/common/config"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
)

func fillAppConfig() {
	config.AppPort = os.Getenv("APP_PORT")
	config.AppSecret = os.Getenv("APP_SECRET")
	config.AppDebug = os.Getenv("APP_DEBUG") == "true"
	config.AppDomain = os.Getenv("APP_TG_API_DOMAIN")
	config.TelegramToken = os.Getenv("APP_TG_TOKEN")
	config.AdminDomain = os.Getenv("APP_ADMIN_DOMAIN")

	if os.Getenv("APP_SECRET_GENERATE") != "true" {
		fmt.Println("WARNING: APP_SECRET_GENERATE is not true")
		return
	}
	config.AppSecret = generateSecret(256)
}

func generateSecret(size int) string {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		panic(fmt.Sprintf("error while generating app secret: %s", err))
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
