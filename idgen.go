package genie

import (
	"fmt"

	nanoid "github.com/matoous/go-nanoid/v2"
)

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"
	length   = 12
)

// GeneratePublicID creates a URL-safe, collision-resistant public identifier
// using nanoid with a custom alphabet. Returns a 12-character string suitable
// for public-facing identifiers.
func GeneratePublicID() (string, error) {
	id, err := nanoid.Generate(alphabet, length)
	if err != nil {
		return "", fmt.Errorf("failed to generate public ID: %w", err)
	}
	return id, nil
}

// MustGeneratePublicID creates a URL-safe public identifier and panics on error.
// Use this when you need a public ID and cannot handle errors gracefully.
func MustGeneratePublicID() string {
	id, err := GeneratePublicID()
	if err != nil {
		panic(err)
	}
	return id
}
