package start

import (
	"compay/core/config"
	"crypto/rand"
	"fmt"
	"os"
	"time"
)

func fillCryptoConfig() {
	config.CryptUseRandSalt = os.Getenv("CRYPT_USE_RAND_SALT") == "true"
	config.CryptGenerateKeys = os.Getenv("CRYPT_GENERATE_KEYS") == "true"
	if !config.CryptUseRandSalt {
		fmt.Println("WARNING: CRYPT_USE_RAND_SALT is not true")
	}
	buf := make([]byte, 8)
	_, err := rand.Read(buf)
	if err != nil {
		panic(fmt.Sprintf("error while generating random string: %s", err))
	}

	if config.CryptUseRandSalt {
		config.CryptRandSaltString = fmt.Sprint(time.Now().Unix(), fmt.Sprintf("%x", buf))
	}
	if !config.CryptGenerateKeys {
		fmt.Println("WARNING: CRYPT_GENERATE_KEYS is not true")
		return
	}
	key := make([]byte, 16)
	_, err = rand.Read(key)
	if err != nil {
		panic(fmt.Sprintf("error while generating random key: %s", err))
	}
	iv := make([]byte, 16)
	_, err = rand.Read(iv)
	if err != nil {
		panic(fmt.Sprintf("error while generating random iv: %s", err))
	}
}
