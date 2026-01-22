package utils

import (
	"crypto/rand"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
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

	if strings.Contains(s, "@") {
		return maskEmail(s)
	}

	runes := []rune(s)
	n := len(runes)

	switch {
	case n > 8:
		return string(runes[:2]) +
			strings.Repeat("*", n-4) +
			string(runes[n-2:])
	case n > 6:
		return string(runes[0]) +
			strings.Repeat("*", n-2) +
			string(runes[n-1])
	default:
		return string(runes[0]) + "****"
	}
}

func maskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "****"
	}

	local := parts[0]
	domain := parts[1]

	localRunes := []rune(local)
	if len(localRunes) <= 2 {
		local = string(localRunes[0]) + "*"
	} else {
		local = string(localRunes[:2]) +
			strings.Repeat("*", len(localRunes)-2)
	}

	domainParts := strings.Split(domain, ".")
	if len(domainParts) < 2 {
		return local + "@****"
	}

	domainName := domainParts[0]
	tld := domainParts[len(domainParts)-1]

	if utf8.RuneCountInString(domainName) > 1 {
		domainName = string([]rune(domainName)[0]) + "***"
	} else {
		domainName = "*"
	}

	return local + "@" + domainName + "." + tld
}

func ToLower(s string) string {
	if len(s) == 0 {
		return ""
	}

	return strings.ToLower(s)
}
