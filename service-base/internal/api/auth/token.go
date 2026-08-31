package auth

import (
	"os"
	"time"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/middleware"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// AccessTokenLifetime is deliberately short - session longevity comes from
// the separate refresh token/cookie (see refresh_token.go and
// RotateRefreshTokenCommand), not from this.
const AccessTokenLifetime = 30 * time.Minute

// MintAccessToken is the single place access tokens get signed, used by
// both login and the refresh endpoint, so the claims shape and lifetime
// can't drift out of sync between them.
func MintAccessToken(userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		middleware.UserIDKey: userID,
		"exp":                time.Now().Add(AccessTokenLifetime).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
