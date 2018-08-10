package main

import (
	"crypto/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ladislavlisy/payroll-maestro-080318/server/authorize"
	"github.com/ladislavlisy/payroll-maestro-080318/server/login"
	"golang.org/x/oauth2"
)

const (
	PAYROLL_MAESTRO_SESSION = "payroll-maestro-session"
)

func NewServer() *gin.Engine {
	authClientID := os.Getenv("AUTH0_CLIENT_ID")
	authSecret := os.Getenv("AUTH0_API_CLIENT_SECRET")
	authDomain := os.Getenv("AUTH0_DOMAIN")
	authCallback := os.Getenv("AUTH0_CALLBACK_URL")
	authAudience := os.Getenv("AUTH0_API_AUDIENCE")

	config := &oauth2.Config{
		ClientID:     authClientID,
		ClientSecret: authSecret,
		RedirectURL:  authCallback,
		Scopes:       []string{"openid", "profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://" + authDomain + "/auth",
			TokenURL: "https://" + authDomain + "/auth-token",
		},
	}

	if authAudience == "" {
		authAudience = "https://" + authDomain + "/userinfo"
	}

	router := gin.Default()

	initRoutes(router, config)

	return router
}

func createSessionStore() sessions.CookieStore {
	// Generate random state
	state := make([]byte, 32)
	rand.Read(state)
	store := sessions.NewCookieStore(state)
	return store
}

func initRoutes(r *gin.Engine, config *oauth2.Config) {
	store := createSessionStore()

	r.Use(sessions.Sessions(PAYROLL_MAESTRO_SESSION, store))

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/js/", "templates/js")

	r.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Payroll MAESTRO",
			},
		)
	})
	r.GET("/login", login.LoginMidleware(config))
	r.GET("/auth-callback", authorize.AuthMidleware(config))
}
