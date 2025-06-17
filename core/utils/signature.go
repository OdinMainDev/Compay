package utils

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strings"
)

func ConstructSignatureBlackrabbit(endpoint string, dataMap map[string]string, secret string) string {
	keys := make([]string, 0, len(dataMap))
	for key := range dataMap {
		if key != "sp_sig" {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)

	var strBuilder strings.Builder
	for _, key := range keys {
		strBuilder.WriteString(dataMap[key])
		strBuilder.WriteString(";")
	}
	str := strBuilder.String()
	hash := md5.Sum([]byte(endpoint + ";" + str + secret))
	return hex.EncodeToString(hash[:])
}
