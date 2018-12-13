package api

import (
	"github.com/labstack/echo"
)

// Top API (Last 10min UUIDs)
func Top(c echo.Context) error {
	return c.JSONPretty(200, struct{ UUID []string }{UUID: []string{"適当", "適当", "適当", "適当"}}, "  ")
}

// JSON API
func JSON(c echo.Context) error {
	return c.JSONPretty(200, struct{ UUID string }{UUID: c.Param("uuid")}, "  ")
}
