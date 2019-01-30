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
	app.GET("/json/:uuid", api.JSON)
	app.GET("/xml/:uuid", api.XML)
	app.GET("/sub", websub.Sub)
	app.POST("/sub", websub.Sub)
	app.Start(":" + os.Getenv("PORT"))
}
