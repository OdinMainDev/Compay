package utils

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func GetHashString(text string) string {
	hash := sha512.New()
	hash.Write([]byte(text))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GenerateRandomHash() string {
	return GetHashString("efwSaltForHashWithRandom" + uuid.NewString() + RandomString(20) + fmt.Sprint(time.Now().UnixNano()))
}
