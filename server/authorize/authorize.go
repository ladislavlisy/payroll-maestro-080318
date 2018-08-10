package authorize

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func AuthMidleware(config *oauth2.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
