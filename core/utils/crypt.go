package utils

import (
	"bytes"
	"compay/core/config"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/bits"
	"strings"
	"time"

	"github.com/google/uuid"
)

func init() {
	// cryptKey
	e1 := []byte{185, 247, 69, 234, 76, 138, 218, 53, 64, 70, 184, 159, 41, 223, 37, 184, 46, 60, 122, 155, 126, 139, 56, 15, 255, 221, 124, 42, 150, 235, 128, 247}
	// cryptIv
	e2 := []byte{153, 69, 245, 163, 21, 94, 243, 53, 19, 240, 144, 159, 127, 170, 61, 191, 173, 97, 102, 206, 94, 219, 0, 3, 85, 221, 84, 42, 22, 154, 176, 219}
	// КЛЮЧ для дешифра/шифра
	key := []byte{
		245, 202, 102, 217, 118, 27, 209, 177,
		112, 64, 69, 207, 76, 153, 197, 143,
		142, 89, 11, 171, 193, 237, 101, 182,
		207, 71, 233, 102, 27, 159, 106, 197,
		85, 130, 93, 5, 77, 234, 157, 95,
		109, 151, 37, 76, 159, 179, 127, 215,
		96, 172, 229, 97, 114, 176, 84, 22,
		123, 11, 84, 152, 149, 120, 248, 133,
	}

	out := make([]byte, len(e1))
	keyLen := len(key)
	for i, b := range e1 {
		k := key[i%keyLen]
		r := int(k^byte(i)) & 7
		x := bits.RotateLeft8(b, 8-r)
		out[i] = x ^ k
	}

	out2 := make([]byte, len(e2))
	for i, b := range e2 {
		k := key[i%keyLen]
		r := int(k^byte(i)) & 7
		x := bits.RotateLeft8(b, 8-r)
		out2[i] = x ^ k
	}

	config.Cryptkey = string(out)
	config.Cryptiv = string(out2)
}

// Crypt Errors
var ErrorCryptBadPadding error = fmt.Errorf("%s", "BadPadding")

func EncryptString(plainText string) (string, error) {
	plainText += "-" + config.CryptRandSaltString
	key, err := hex.DecodeString(config.Cryptkey)
	if err != nil {
		return "", err
	}

	iv, err := hex.DecodeString(config.Cryptiv)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	paddedText := pkcs7Pad([]byte(plainText), aes.BlockSize)
	mode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(paddedText))
	mode.CryptBlocks(cipherText, paddedText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func DecryptString(cipherText string) (string, error) {
	key, err := hex.DecodeString(config.Cryptkey)
	if err != nil {
		return "", err
	}

	iv, err := hex.DecodeString(config.Cryptiv)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherBytes))
	mode.CryptBlocks(plainText, cipherBytes)

	unpaddedText, err := pkcs7Unpad(plainText)
	if err != nil {
		return "", err
	}

	return strings.Replace(string(unpaddedText), "-"+config.CryptRandSaltString, "", -1), nil
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	padding := int(data[length-1])
	if padding > length || padding == 0 {
		return nil, ErrorCryptBadPadding
	}
	return data[:length-padding], nil
}

func GetHashString(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GenerateRandomSalt() string {
	return GetHashString("esvEdroSaltRand" + uuid.NewString() + RandomString(20) + fmt.Sprint(time.Now().Unix()))
}
