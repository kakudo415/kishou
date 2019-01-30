package main

import (
	"os"

	"github.com/labstack/echo"

	"./api"
	"./websub"
)

func main() {
	app := echo.New()
	app.GET("/", api.Top)
	// JSON & XML API
	app.GET("/json/:uuid", api.JSON)
	app.GET("/xml/:uuid", api.XML)
	// Subscriber
	app.GET("/sub", websub.Sub)
	app.POST("/sub", websub.Sub)
	// Error = 404
	app.HTTPErrorHandler = func(e error, c echo.Context) {
		c.NoContent(404)
	}
	// Run on $PORT
	app.Start(":" + os.Getenv("PORT"))
}
