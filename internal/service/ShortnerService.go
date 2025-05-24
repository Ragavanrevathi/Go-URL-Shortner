package service

import (
	"math/rand"
	"shorten-url/pkg/config"
	"strings"
)

var m = make(map[string]string)

func ShortURL(URL string) string {
	key := generateShortKey()
	if !strings.HasPrefix(URL, "http://") && !strings.HasPrefix(URL, "https://") {
		URL = "http://" + URL
	}
	m[key] = URL
	shortenURL := config.Env.Domain + "/" + key
	return shortenURL
}

func GetUserURL(shortKey string) (string, bool) {
	originalURL, exists := m[shortKey]
	return originalURL, exists
}

// Generate Unique key with a length of six
func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
