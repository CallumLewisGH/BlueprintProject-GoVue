package routes

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/contracts/requests"
	command "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/handlers/commands"
	query "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/handlers/queries"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/markbates/goth/gothic"
)

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

	user, err := query.GetUserByAuthIdQuery(c, authUser.UserID)

	if err != nil {
		userReq := requests.CreateUserRequest{
			Email:    authUser.Email,
			AuthId:   authUser.UserID,
			Username: deriveUsername(authUser.Email),
		}

		user, err = command.CreateUserCommand(c, userReq)

		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "The username based off of the email: " + userReq.Email + " is already in use. " + err.Error()})
			return
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userAuthId": authUser.UserID,
		"userId":     user.ID,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	target := fmt.Sprintf("%s?token=%s", os.Getenv("FRONTEND_URL"), tokenString)
	c.Redirect(http.StatusFound, target)
}

// Logout godoc
// @Summary Log out the user
// @Description Clears the session for the specified provider
// @Tags Authentication
// @Param provider path string true "OAuth Provider"
// @Success 307 "Redirect to Home"
// @Router /authentication/logout/{provider} [get]
func logout(c *gin.Context) {
	q := c.Request.URL.Query()
	q.Add("provider", c.Param("provider"))
	c.Request.URL.RawQuery = q.Encode()

	gothic.Logout(c.Writer, c.Request)
	target := os.Getenv("FRONTEND_URL")
	c.Redirect(http.StatusTemporaryRedirect, target)
}
