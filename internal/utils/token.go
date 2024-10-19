package utils

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

func GenerateToken(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func GenerateExpireDate(days int) time.Time {
	if days < 0 {
		return time.Now().UTC()
	}
	expire := time.Now().UTC().Add(time.Duration(days) * 24 * time.Hour)
	return expire
}
