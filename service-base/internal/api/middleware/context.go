package middleware

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
)

type AuthenticatedInput struct {
	UserID uuid.UUID `json:"-"`
}

func (a *AuthenticatedInput) Resolve(ctx huma.Context) []error {
	if val := ctx.Context().Value(UserIDKey); val != nil {
		userId, err := uuid.Parse(val.(string))

		if err != nil {
			return []error{&huma.ErrorDetail{
				Message: "Invalid user ID in context",
			}}
		}

		a.UserID = userId
		return nil
	}

	return []error{&huma.ErrorDetail{
		Message: "Unauthorized: User ID not found in context",
	}}
}
