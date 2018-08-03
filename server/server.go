package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	router := gin.Default()

	initRoutes(router)

	return router
}

func initRoutes(r *gin.Engine) {
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
}
