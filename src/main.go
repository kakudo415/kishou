package main

import (
	"github.com/labstack/echo"

	"./api"
	"./websub"
)

func main() {
	app := echo.New()
	app.GET("/", api.Top)
	app.GET("/:uuid", api.JSON)
	app.GET("/sub", websub.Sub)
	app.POST("/sub", websub.Sub)
	app.Start(":10200")
}
