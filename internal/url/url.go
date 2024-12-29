package url

import (
	"crypto/sha256"
	"encoding/base64"
)

func ShortenUrl(url string) string {
	hasher := sha256.New()
	hasher.Write([]byte(url))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))[:8] // Use first 8 characters
}
