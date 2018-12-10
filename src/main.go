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
	app.GET("/json", page.JSON)
	app.GET("/json/:ts", api.JSON)
	app.Run()
}
