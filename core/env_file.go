package core

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func InitAndCheckEnv(envCheckKeys []string, envCallbacks []func()) {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Environment loading error %s", err.Error()))
	}

	for _, value := range envCheckKeys {
		if os.Getenv(value) == "" {
			panic(fmt.Sprintf(".env parameter %s is missing", value))
		}
	}

	for _, callback := range envCallbacks {
		callback()
	}
}
