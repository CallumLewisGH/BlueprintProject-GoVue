package command

import (
	"context"
	"errors"
	"time"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/task"
	user "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrInvalidRefreshToken = errors.New("invalid or expired refresh token")

type RotatedRefreshToken struct {
	UserID   uuid.UUID
	RawToken string
}

// RotateRefreshTokenCommand validates rawToken against the stored hash and,
// if it's still within its sliding window, issues a brand new one for the
// same user - the old one stops matching anything the instant this runs, so
// a stale (already-rotated) cookie is naturally rejected if reused, and an
// active session's expiry keeps resetting for as long as it keeps being used.
func RotateRefreshTokenCommand(ctx context.Context, rawToken string) (*RotatedRefreshToken, error) {
	commandFunc := func(db *gorm.DB, ctx context.Context) (*RotatedRefreshToken, error) {
		hash := user.HashRefreshToken(rawToken)
		repo := user.NewRepo(db)

		existing, err := repo.WithRefreshTokenHash(hash).First()
		if err != nil {
			return nil, ErrInvalidRefreshToken
		}
		if existing.RefreshTokenExpiresAt == nil || existing.RefreshTokenExpiresAt.Before(time.Now()) {
			return nil, ErrInvalidRefreshToken
		}

		newRaw, newHash, err := user.GenerateRefreshToken()
		if err != nil {
			return nil, err
		}

		// Scoped by ID, not by the (now-stale) old hash - the old hash is
		// exactly the column this update is about to change.
		if _, err := repo.WithId(existing.ID).UpdateOne(map[string]any{
			"refresh_token_hash":       newHash,
			"refresh_token_expires_at": time.Now().Add(user.RefreshTokenLifetime),
		}); err != nil {
			return nil, err
		}

		return &RotatedRefreshToken{UserID: existing.ID, RawToken: newRaw}, nil
	}
	return task.NewCommandTask(ctx, commandFunc).Await()
}
