package authorize

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func AuthMidleware(config *oauth2.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		domain := os.Getenv("AUTH0_DOMAIN")

		urlState := ctx.Request.URL.Query().Get("state")

		sessions := sessions.Default(ctx)

		state := sessions.Get("state")

		if state != urlState {
			ctx.AbortWithError(http.StatusInternalServerError, errors.New("Invalid state parameter"))
			return
		}

		code := ctx.Request.URL.Query().Get("code")

		token, err := config.Exchange(context.TODO(), code)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// Getting now the userInfo
		client := config.Client(context.TODO(), token)
		resp, err := client.Get("https://" + domain + "/userinfo")
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		defer resp.Body.Close()

		var profile map[string]interface{}
		if err = json.NewDecoder(resp.Body).Decode(&profile); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		sessions.Set("id_token", token.Extra("id_token"))
		sessions.Set("access_token", token.AccessToken)
		sessions.Set("profile", profile)
		err = sessions.Save()

		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// Redirect to logged in page
		ctx.Redirect(http.StatusSeeOther, "/user")
	}
}
