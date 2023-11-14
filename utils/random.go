package utils

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomNumber generates random integer between min and max
func RandomNumber(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates random string for a given length
func RandomString(n int) string {
	var sb strings.Builder

	for i := 1; i <= n; i++ {
		c := alphabet[rand.Intn(i)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomEmail generates random email address for a given domain
func RandomEmail(params ...string) string {
	prefix := RandomString(8)

	var domain string
	if len(params) <= 0 {
		domain = "example.com"
	} else {
		domain = params[0]
	}

	return fmt.Sprintf("%s@%s", prefix, domain)
}
