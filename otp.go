package genie

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
)

const otpLength = 6

// GenerateOTP creates a 6-character one-time password (OTP) using base32 encoding,
// suitable for short-lived authentication codes.
func GenerateOTP() (string, error) {
	buffer := make([]byte, otpLength)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("failed to generate OTP: %w", err)
	}
	return base32.StdEncoding.EncodeToString(buffer)[:otpLength], nil
}
