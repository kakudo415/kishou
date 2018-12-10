package main

import (
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
	app.Run()
}
