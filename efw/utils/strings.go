package utils

import (
	"crypto/rand"
	"fmt"
	"regexp"
	"strings"
)

func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func OnlyNumberString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func RandomString(length int) string {
	b := make([]byte, length/2+1)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", b)[:length]
}

func RemoveNonAlphaNumeric(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return str
	}

	return reg.ReplaceAllString(str, "")
}

func RemoveDoubleSlashesFromUrl(url string) string {
	parts := strings.Split(url, "/")
	cleanParts := []string{}
	for _, part := range parts {
		if part != "" {
			cleanParts = append(cleanParts, part)
		}
	}

	result := strings.Join(cleanParts, "/")
	if strings.HasPrefix(url, "http://") {
		result = "http://" + result
	} else if strings.HasPrefix(url, "https://") {
		result = "https://" + result
	}
	return result
}

func MaskUsername(s string) string {
	if s == "" {
		return "****"
	}
	runes := []rune(s)
	n := len(runes)
	keep := max(n-4, 1)
	return string(runes[:keep]) + "****"
}
