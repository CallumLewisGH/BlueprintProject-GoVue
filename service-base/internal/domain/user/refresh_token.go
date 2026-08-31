package user

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// RefreshTokenLifetime is a sliding window - every successful refresh
// rotates the token and resets this from the moment of use, not from
// original login. See RotateRefreshTokenCommand.
const RefreshTokenLifetime = 30 * 24 * time.Hour

// GenerateRefreshToken returns a fresh opaque secret (raw, to hand to the
// client as a cookie value) and its hash (to store - never the raw value,
// same principle as password storage).
func GenerateRefreshToken() (raw string, hash string, err error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", "", err
	}
	raw = hex.EncodeToString(buf)
	return raw, HashRefreshToken(raw), nil
}

func HashRefreshToken(raw string) string {
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}
