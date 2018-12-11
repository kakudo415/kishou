package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"

	"./api"
	"./websub"
)

func main() {
	app := gin.New()
	app.Use(gzip.Gzip(gzip.DefaultCompression))

	// Index page
	app.StaticFile("/", "doc")
	app.StaticFS("/static/", http.Dir("doc/static"))
	// Subscriber
	app.GET("/subscriber", websub.Subscriber)
	app.POST("/subscriber", websub.Receiver)
	// API
	app.GET("/3min.xml", api.XML)
	app.GET("/3min.json", api.JSON)

	// Redirect Trailing Slash
	app.GET("/subscriber/", func(c *gin.Context) {
		c.Redirect(301, strings.TrimSuffix(c.Request.URL.String(), "/"))
	})
	app.GET("/3min.xml/", func(c *gin.Context) {
		c.Redirect(301, strings.TrimSuffix(c.Request.URL.String(), "/"))
	})
	app.GET("/3min.json/", func(c *gin.Context) {
		c.Redirect(301, strings.TrimSuffix(c.Request.URL.String(), "/"))
	})

	// Run on $PORT
	app.Run()
}
