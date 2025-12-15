package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"efw/config"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

// Crypt Errors
var (
	ErrorCryptBadPadding         = fmt.Errorf("%s", "BadPadding")
	ErrorCryptHexDecodeKeyError  = func(err error) error { return fmt.Errorf("%s: %s", "HexDecodeKeyError", err.Error()) }
	ErrorCryptHexDecodeIvError   = func(err error) error { return fmt.Errorf("%s: %s", "HexDecodeIvError", err.Error()) }
	ErrorCryptAesNewCipher       = func(err error) error { return fmt.Errorf("%s: %s", "AesNewCipher", err.Error()) }
	ErrorCryptBase64DecodeString = func(err error) error { return fmt.Errorf("%s: %s", "Base64DecodeString", err.Error()) }
	ErrorCryptPkcs7Unpad         = func(err error) error { return fmt.Errorf("%s: %s", "Pkcs7Unpad", err.Error()) }
)

func EncryptString(plainText string) (string, error) {
	plainText += "-" + config.CrypSalt
	key, err := hex.DecodeString(config.Cryptkey)
	if err != nil {
		return "", ErrorCryptHexDecodeKeyError(err)
	}

	iv, err := hex.DecodeString(config.Cryptiv)
	if err != nil {
		return "", ErrorCryptHexDecodeIvError(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", ErrorCryptAesNewCipher(err)
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
		return "", ErrorCryptHexDecodeKeyError(err)
	}

	iv, err := hex.DecodeString(config.Cryptiv)
	if err != nil {
		return "", ErrorCryptHexDecodeIvError(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", ErrorCryptAesNewCipher(err)
	}

	cipherBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", ErrorCryptBase64DecodeString(err)
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherBytes))
	mode.CryptBlocks(plainText, cipherBytes)

	unpaddedText, err := pkcs7Unpad(plainText)
	if err != nil {
		return "", ErrorCryptPkcs7Unpad(err)
	}

	return strings.ReplaceAll(string(unpaddedText), "-"+config.CrypSalt, ""), nil
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
