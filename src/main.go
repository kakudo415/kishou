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
	app.GET("/:uuid", api.JSONMinify)
	app.GET("/json/:uuid", api.JSONMinify)
	app.GET("/json/minify/:uuid", api.JSONMinify)
	app.GET("/json/pretty/:uuid", api.JSONPretty)
	app.GET("/xml/:uuid", api.XML)
	app.GET("/sub", websub.Sub)
	app.POST("/sub", websub.Sub)
	app.Start(":" + os.Getenv("PORT"))
}
