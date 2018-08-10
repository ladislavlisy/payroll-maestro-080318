package authorize

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// Midleware - Handler to Callback
func Midleware(config *oauth2.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Auth request (begin) >>>")

		domain := os.Getenv("AUTH0_DOMAIN")

		urlState := ctx.Request.URL.Query().Get("state")
		log.Println("Auth request (state): " + urlState)

		sessions := sessions.Default(ctx)

		state := sessions.Get("state")
		log.Printf("Auth request (state): %s\n", state)

		if state != urlState {
			err := errors.New("Invalid state parameter")
			log.Println("Auth request (state): " + err.Error())
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		code := ctx.Request.URL.Query().Get("code")
		log.Println("Auth request (code): " + code)

		token, err := config.Exchange(oauth2.NoContext, code)
		if err != nil {
			log.Println("Auth request (token): " + err.Error())
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		log.Println("Auth request (token): " + token.AccessToken)

		// Getting now the userInfo
		client := config.Client(context.TODO(), token)
		resp, err := client.Get("https://" + domain + "/userinfo")
		if err != nil {
			log.Println("Auth request (client): " + err.Error())
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		defer resp.Body.Close()

		var profile map[string]interface{}
		if err = json.NewDecoder(resp.Body).Decode(&profile); err != nil {
			log.Println("Auth request (profile): " + err.Error())
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		sessions.Set("id_token", token.Extra("id_token"))
		sessions.Set("access_token", token.AccessToken)
		sessions.Set("profile", profile)
		err = sessions.Save()

		if err != nil {
			log.Println("Auth request (sessions): " + err.Error())
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// Redirect to logged in page
		ctx.Redirect(http.StatusSeeOther, "/user")
	}
}
