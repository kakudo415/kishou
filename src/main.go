package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"./api"
	"./websub"
)

func main() {
	app := gin.New()

	// Index page
	app.StaticFile("/", "doc")
	app.StaticFS("/static/", http.Dir("doc/static"))
	// Subscriber
	app.GET("/subscriber", websub.Subscriber)
	app.POST("/subscriber", websub.Receiver)
	// API
	app.GET("/raw", api.XML)
	app.GET("/xml", api.XML)
	app.GET("/json", api.JSON)

	// Redirect Trailing Slash
	app.GET("/subscriber/", func(c *gin.Context) {
		c.Redirect(301, strings.TrimSuffix(c.Request.URL.String(), "/"))
	})
	app.GET("/raw/", func(c *gin.Context) {
		c.Redirect(301, strings.TrimSuffix(c.Request.URL.String(), "/"))
	})
	app.GET("/xml/", func(c *gin.Context) {
		c.Redirect(301, strings.TrimSuffix(c.Request.URL.String(), "/"))
	})
	app.GET("/json/", func(c *gin.Context) {
		c.Redirect(301, strings.TrimSuffix(c.Request.URL.String(), "/"))
	})

	// Run on $PORT
	app.Run()
}
