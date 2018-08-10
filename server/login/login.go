package login

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func createRandomUID() string {
	// Generate random state
	byteuid := make([]byte, 32)
	rand.Read(byteuid)
	cstruid := base64.StdEncoding.EncodeToString(byteuid)
	return cstruid
}

// Midleware - Handler to Login
func Midleware(config *oauth2.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		domain := os.Getenv("AUTH0_DOMAIN")
		aud := os.Getenv("AUTH0_AUDIENCE")

		if aud == "" {
			aud = "https://" + domain + "/userinfo"
		}

		sessions := sessions.Default(ctx)

		state := createRandomUID()
		sessions.Set("state", state)
		sessions.Save()

		audience := oauth2.SetAuthURLParam("audience", aud)
		url := config.AuthCodeURL(state, audience)

		ctx.Redirect(http.StatusTemporaryRedirect, url)
	}
}
