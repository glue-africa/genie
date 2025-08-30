package genie

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"fmt"
)

// GenerateToken creates a URL-safe random token of the specified length using
// base64 URL encoding, suitable for general-purpose secure tokens.
//
// Note: This function truncates the encoded result to match the requested length,
// which may result in slightly reduced entropy compared to the raw bytes.
// For session tokens, use GenerateSessionToken instead.
func GenerateToken(length int) (string, error) {
	byteLength := (length * 3) / 4
	if (length*3)%4 != 0 {
		byteLength++
	}

	randomBytes := make([]byte, byteLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	token := base64.URLEncoding.EncodeToString(randomBytes)
	return token[:length], nil
}

// GenerateSessionToken creates a cryptographically secure session token with
// 256 bits of entropy using base32 encoding. The resulting token is 52 characters
// long and is case-insensitive, making it suitable for session management.
//
// Returns a base32-encoded string representing 32 bytes of random data.
func GenerateSessionToken() (string, error) {
	token := make([]byte, 32) // 256 bits of entropy
	if _, err := rand.Read(token); err != nil {
		return "", fmt.Errorf("failed to generate session token: %w", err)
	}
	return base32.StdEncoding.EncodeToString(token), nil
}
