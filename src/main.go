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
	// Router
	app.GET("/", page.Index)
	app.GET("/subscriber", websub.Subscriber)
	app.POST("/subscriber", websub.Receiver)
	app.GET("/json", api.JSON)
	// Redirect Trailing Slash
	app.GET("/subscriber/", func(c *gin.Context) {
		c.Redirect(301, strings.TrimSuffix(c.Request.URL.String(), "/"))
	})
	app.GET("/json/", func(c *gin.Context) {
		c.Redirect(301, strings.TrimSuffix(c.Request.URL.String(), "/"))
	})
	// Run on $PORT
	app.Run()
}
