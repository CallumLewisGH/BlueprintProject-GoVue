package auth

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func NewAuth() {
	key := os.Getenv("SESSION_KEY")
	isProd := os.Getenv("ENVIRONMENT") == "prod"

	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	backendURL := os.Getenv("BACKEND_URL")

	store := sessions.NewCookieStore([]byte(key))

	sameSiteMode := http.SameSiteLaxMode
	if isProd {
		sameSiteMode = http.SameSiteNoneMode
	}

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 30, // 30 days
		HttpOnly: true,
		Secure:   isProd,
		SameSite: sameSiteMode,
	}

	gothic.Store = store

	providers := []goth.Provider{
		google.New(googleClientId, googleClientSecret, backendURL+"/authentication/google/callback"),
	}

	goth.UseProviders(providers...)
}
