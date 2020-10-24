package problem055

import (
	"crypto/sha256"
	"fmt"
)

var resolver map[string]string

func Shorten(url string) string {
	if resolver == nil {
		resolver = map[string]string{}
	}
	shortenedUrl := fmt.Sprintf("%x", sha256.Sum256([]byte(url)))[:6]
	resolver[shortenedUrl] = url
	return shortenedUrl
}

func Resolve(shortened string) string {
	return resolver[shortened]
}
