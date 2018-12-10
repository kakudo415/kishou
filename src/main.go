package main

import (
	"strings"

	"./api"
	"./page"
	"./websub"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.LoadHTMLGlob("doc/*.html")
	app.GET("/", page.Index)
	app.GET("/sub", websub.Subscriber)
	app.POST("/sub", websub.Receiver)
	app.GET("/xml", page.XML)
	app.GET("/xml/:time", api.XML)
	app.GET("/json", page.JSON)
	app.GET("/json/:time", api.JSON)
	// Redirect Trailing Slash
	app.GET("/sub/", func(c *gin.Context) {
		c.Redirect(301, strings.TrimSuffix(c.Request.URL.String(), "/"))
	})
	app.GET("/xml/", func(c *gin.Context) {
		c.Redirect(301, strings.TrimSuffix(c.Request.URL.String(), "/"))
	})
	app.GET("/json/", func(c *gin.Context) {
		c.Redirect(301, strings.TrimSuffix(c.Request.URL.String(), "/"))
	})
	app.Run()
}
