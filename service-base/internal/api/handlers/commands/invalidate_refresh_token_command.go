package command

import (
	"context"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/task"
	user "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"
	"gorm.io/gorm"
)

// InvalidateRefreshTokenCommand ends the refresh session rawToken belongs
// to (logout). Already-issued access tokens keep working until their own
// short expiry - there's no cheap way to kill an already-signed JWT
// instantly without a lookup on every single request, and that's exactly
// the per-request cost the short access-token lifetime is meant to avoid.
// If rawToken doesn't match anything (already logged out, expired, or
// simply absent), this is a no-op rather than an error - logging out an
// already-logged-out session isn't a failure.
func InvalidateRefreshTokenCommand(ctx context.Context, rawToken string) error {
	commandFunc := func(db *gorm.DB, ctx context.Context) (struct{}, error) {
		hash := user.HashRefreshToken(rawToken)
		repo := user.NewRepo(db)

		existing, err := repo.WithRefreshTokenHash(hash).First()
		if err != nil {
			return struct{}{}, nil
		}

		_, err = repo.WithId(existing.ID).UpdateOne(map[string]any{
			"refresh_token_hash":       nil,
			"refresh_token_expires_at": nil,
		})
		return struct{}{}, err
	}
	_, err := task.NewCommandTask(ctx, commandFunc).Await()
	return err
}
