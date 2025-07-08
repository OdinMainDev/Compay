package start

import (
	"compay/common/config"
	"os"
)

func fillAppConfig() {
	config.AppPort = os.Getenv("APP_PORT")
	config.AppDebug = os.Getenv("APP_DEBUG") == "true"

	config.AppSecret = generateSecret(256)
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
