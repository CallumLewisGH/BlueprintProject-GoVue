package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/golang-jwt/jwt/v5"
)

const UserIDKey string = "userId"

func RequireAuthHuma(api huma.API) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		authHeader := ctx.Header("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, "Missing or malformed token")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if userID, ok := claims[UserIDKey].(string); ok {
				next(huma.WithValue(ctx, UserIDKey, userID))
				return
			}
		}

		_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, "Token missing user identification")
	}
}
