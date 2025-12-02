package efw

import (
	"crypto/rand"
	"fmt"
)

func GenerateCryptKeys() (string, string, error) {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		return "", "", fmt.Errorf("%s %s", "[EFRAME_FUNCTIONS] ErrorWhileGeneratingRandom[Key]:", err.Error())
	}
	iv := make([]byte, 16)
	_, err = rand.Read(iv)
	if err != nil {
		return "", "", fmt.Errorf("%s %s", "[EFRAME_FUNCTIONS] ErrorWhileGeneratingRandom[Iv]:", err.Error())
	}
	return fmt.Sprintf("%x", key), fmt.Sprintf("%x", iv), nil
}
