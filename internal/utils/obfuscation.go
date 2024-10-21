package utils

import (
	"encoding/base64"
)

// Obfuscate encodes the input string to Base64.
func Obfuscate(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// Deobfuscate decodes the Base64 encoded string back to the original string.
func Deobfuscate(data string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}
