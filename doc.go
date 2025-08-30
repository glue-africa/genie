// Package genie provides utilities for generating various types of secure,
// cryptographically random identifiers and tokens including:
//
// - Public IDs using nanoid with custom alphabet for URL-safe, collision-resistant identifiers
// - One-time passwords (OTPs) using base32 encoding for authentication codes
// - General-purpose tokens using base64 URL encoding
// - Session tokens with 256 bits of entropy using base32 encoding
package genie
