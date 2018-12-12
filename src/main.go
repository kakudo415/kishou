package main

import (
	"github.com/gin-gonic/gin"

	"./api"
	"./websub"
)

func main() {
	app := gin.New()

	// Subscriber
	app.GET("/subscriber", websub.Subscriber)
	app.POST("/subscriber", websub.Receiver)
	// API
	app.GET("/", api.Top)
	app.GET("/xml/:uuid", api.XML)
	app.GET("/json/:uuid", api.JSON)

	// Run on $PORT
	app.Run()
}
