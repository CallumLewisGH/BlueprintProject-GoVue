package routes

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strings"
	"unicode"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/auth"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/contracts/requests"
	command "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/handlers/commands"
	query "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/handlers/queries"
	user "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

const refreshCookieName = "refresh_token"

// RegisterAuthRoutes godoc
// @Summary Authentication-related endpoints
// @Description Handles OAuth2 login, callbacks, and logout using Goth
// @Tags Authentication
func RegisterAuthRoutes(s *api.Server) {
	authGroup := s.Group("/authentication")
	{
		authGroup.GET("/:provider", beginAuth)
		authGroup.GET("/:provider/callback", authCallback)
		// Some OAuth providers post their callback as a form submission (form_post
		// response mode) rather than a GET redirect - Google doesn't, but keep this
		// route available for providers that do.
		authGroup.POST("/:provider/callback", authCallback)
		authGroup.POST("/refresh", refreshToken)
		authGroup.GET("/logout/:provider", logout)
	}
}

// Begin Auth godoc
// @Summary Initiate OAuth login
// @Description Redirects the user to the provider's login page
// @Tags Authentication
// @Param provider path string true "OAuth Provider (e.g., google)"
// @Success 307 "Temporary Redirect to Provider"
// @Router /authentication/{provider} [get]
func beginAuth(c *gin.Context) {
	q := c.Request.URL.Query()
	q.Add("provider", c.Param("provider"))
	c.Request.URL.RawQuery = q.Encode()

	if _, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		c.Redirect(http.StatusFound, os.Getenv("FRONTEND_URL"))
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}

// deriveUsername turns an OAuth email's local-part into a username satisfying the 8-16
// char alphanumeric constraint on CreateUserRequest. Google doesn't guarantee that range
// (e.g. "cal@gmail.com" is only 3 chars), so this pads or truncates rather than trusting
// the email to already fit.
func deriveUsername(email string) string {
	local := strings.SplitN(email, "@", 2)[0]

	var sb strings.Builder
	for _, r := range local {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(r)
		}
	}
	username := sb.String()

	for len(username) < 8 {
		username += randomDigit()
	}
	if len(username) > 16 {
		username = username[:16]
	}
	return username
}

func randomDigit() string {
	n, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		return "0"
	}
	return n.String()
}

// Auth Callback godoc
// @Summary OAuth callback URL
// @Description The endpoint the provider redirects to after successful login
// @Tags Authentication
// @Param provider path string true "OAuth Provider"
// @Success 200 Redirect to Frontend
// @Failure 400 {object} map[string]string "Authentication failed"
// @Router /authentication/{provider}/callback [get]
func authCallback(c *gin.Context) {
	q := c.Request.URL.Query()
	q.Add("provider", c.Param("provider")) /*  */
	c.Request.URL.RawQuery = q.Encode()

	authUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	loggedInUser, err := query.GetUserByAuthIdQuery(c, authUser.UserID)

	if err != nil {
		userReq := requests.CreateUserRequest{
			Email:    authUser.Email,
			AuthId:   authUser.UserID,
			Username: deriveUsername(authUser.Email),
		}

		loggedInUser, err = command.CreateUserCommand(c, userReq)

		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "The username based off of the email: " + userReq.Email + " is already in use. " + err.Error()})
			return
		}
	}

	accessToken, err := auth.MintAccessToken(loggedInUser.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// A failure here just means no silent refresh for this session (falls
	// back to a full re-login once the access token expires) - not worth
	// blocking login over.
	if rawRefreshToken, err := command.IssueRefreshTokenCommand(c, loggedInUser.ID); err == nil {
		setRefreshCookie(c, rawRefreshToken)
	}

	target := fmt.Sprintf("%s?token=%s", os.Getenv("FRONTEND_URL"), accessToken)
	c.Redirect(http.StatusFound, target)
}

// Refresh Token godoc
// @Summary Silently renew the access token
// @Description Reads the refresh cookie (works even if the access token has already expired), rotates it, and issues a fresh access token
// @Tags Authentication
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string "Missing, invalid, or expired refresh token"
// @Router /authentication/refresh [post]
func refreshToken(c *gin.Context) {
	rawToken, err := c.Cookie(refreshCookieName)
	if err != nil || rawToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No refresh token"})
		return
	}

	rotated, err := command.RotateRefreshTokenCommand(c, rawToken)
	if err != nil {
		clearRefreshCookie(c)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	accessToken, err := auth.MintAccessToken(rotated.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to issue token"})
		return
	}

	setRefreshCookie(c, rotated.RawToken)
	c.JSON(http.StatusOK, gin.H{"token": accessToken})
}

// Logout godoc
// @Summary Log out the user
// @Description Ends the refresh session and clears the OAuth session for the specified provider
// @Tags Authentication
// @Param provider path string true "OAuth Provider"
// @Success 307 "Redirect to Home"
// @Router /authentication/logout/{provider} [get]
func logout(c *gin.Context) {
	// Cookies attach to this navigation regardless of it not being a fetch
	// call, so the refresh cookie (if any) is present here without the
	// frontend needing to do anything extra for logout to work.
	if rawToken, err := c.Cookie(refreshCookieName); err == nil && rawToken != "" {
		_ = command.InvalidateRefreshTokenCommand(c, rawToken)
	}
	clearRefreshCookie(c)

	q := c.Request.URL.Query()
	q.Add("provider", c.Param("provider"))
	c.Request.URL.RawQuery = q.Encode()

	gothic.Logout(c.Writer, c.Request)
	target := os.Getenv("FRONTEND_URL")
	c.Redirect(http.StatusTemporaryRedirect, target)
}

// setRefreshCookie and clearRefreshCookie are scoped to /authentication -
// the only routes that ever need to read this cookie - rather than the
// whole site, keeping its exposure as narrow as possible.

func setRefreshCookie(c *gin.Context, rawToken string) {
	isProd := os.Getenv("ENVIRONMENT") != "dev" && os.Getenv("ENVIRONMENT") != "development"
	sameSite := http.SameSiteLaxMode
	if isProd {
		// If your frontend and backend end up on different domains in prod,
		// this needs SameSite=None (+ Secure, required alongside it) to be
		// sent on cross-origin fetches at all - same reasoning as the OAuth
		// session cookie in internal/api/auth/auth.go. Lax is fine if
		// they're on the same site.
		sameSite = http.SameSiteNoneMode
	}
	c.SetSameSite(sameSite)
	c.SetCookie(refreshCookieName, rawToken, int(user.RefreshTokenLifetime.Seconds()), "/authentication", "", isProd, true)
}

func clearRefreshCookie(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(refreshCookieName, "", -1, "/authentication", "", false, true)
}
