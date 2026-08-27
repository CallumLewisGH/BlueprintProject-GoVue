package middleware

import (
	"net/http"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/database"
	user "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"
	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
)

// RequireAdminHuma must run after RequireAuthHuma, which populates UserIDKey
// in the request context. Role is looked up per-request rather than trusted
// from the JWT so a role change takes effect immediately, not after the
// token's 72h expiry.
func RequireAdminHuma(api huma.API) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		userIDStr, ok := ctx.Context().Value(UserIDKey).(string)
		if !ok {
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, "Missing user identification")
			return
		}

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid user identification")
			return
		}

		db := database.GetDatabase().GetGormDatabase()
		u, err := user.NewRepo(db).WithId(userID).First()
		if err != nil || u.Role != user.RoleAdmin {
			_ = huma.WriteErr(api, ctx, http.StatusForbidden, "Administrator access required")
			return
		}

		next(ctx)
	}
}
