package utils

import (
	"crypto/sha1"
	"encoding/base64"
)

// GenerateDeviceID creates new device ID out of its name
func GenerateDeviceID(name string) string {
	hash := sha1.Sum([]byte(name))
	hash64 := base64.StdEncoding.EncodeToString(hash[:])
	return hash64
}
