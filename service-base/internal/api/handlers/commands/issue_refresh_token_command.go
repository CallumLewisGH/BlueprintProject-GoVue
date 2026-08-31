package command

import (
	"context"
	"time"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/task"
	user "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// IssueRefreshTokenCommand starts a fresh refresh session for userID
// (login), overwriting any previous one - only one active session per user
// is supported for now, so logging in elsewhere ends the old session's
// ability to silently refresh.
func IssueRefreshTokenCommand(ctx context.Context, userID uuid.UUID) (string, error) {
	commandFunc := func(db *gorm.DB, ctx context.Context) (string, error) {
		raw, hash, err := user.GenerateRefreshToken()
		if err != nil {
			return "", err
		}

		expiresAt := time.Now().Add(user.RefreshTokenLifetime)
		repo := user.NewRepo(db)
		if _, err := repo.WithId(userID).UpdateOne(map[string]any{
			"refresh_token_hash":       hash,
			"refresh_token_expires_at": expiresAt,
		}); err != nil {
			return "", err
		}

		return raw, nil
	}
	return task.NewCommandTask(ctx, commandFunc).Await()
}
