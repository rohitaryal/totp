// Package totp contains helper functions for secret key generation
// and totp generation using secret key and timestamp
package totp

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
)

// GenerateSecret - generates a random base32 secret key
func GenerateSecret() (string, error) {
	buffer := make([]byte, 20)

	if _, err := rand.Read(buffer); err != nil {
		return "", err
	}

	// Convert byte to base32 and remove '='s
	return strings.TrimRight(base32.StdEncoding.EncodeToString(buffer), "="), nil
}

// GenerateTotp - Generates TOTP for corresponding secret key
func GenerateTotp(secretKey string, timestamp int64) (string, error) {
	// From the HOTP, the counter c(t) is defined as:
	//
	// c(t) = (t - t_0) / X
	// Where,
	//		t_0 = starting time (0 for unix timestamp)
	//		t   = current time
	//		X   = Period for code rotation (30 seconds here)
	counter := timestamp / 30

	// Convert the counter to byte slice
	counterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(counterBytes, uint64(counter))

	// Convert string to hex literal
	decodedKey, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secretKey)
	if err != nil {
		return "", err
	}

	// Hmac with (in their byte/byte slice form):
	// 		key = secretKey
	// 		msg = counter
	mac := hmac.New(sha1.New, decodedKey)
	mac.Write(counterBytes)
	hash := mac.Sum(nil)

	offset := hash[len(hash)-1] & 0x0f

	binaryCode := (int(hash[offset]&0x7f) << 24) |
		(int(hash[offset+1]&0xff) << 16) |
		(int(hash[offset+2]&0xff) << 8) |
		(int(hash[offset+3] & 0xff))

	// Generate 6 digit TOTP
	return fmt.Sprintf("%06d", binaryCode%1_000_000), nil
}
