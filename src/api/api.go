package api

import (
	"strings"

	"github.com/labstack/echo"

	"../kvs"
)

// TopJSON Format
type TopJSON struct {
	UUID []string `json:"UUID"`
}

// Top API (Last 10min UUIDs)
func Top(c echo.Context) error {
	ids := kvs.KEYS("KISHOW:*")
	for i := 0; i < len(ids); i++ {
		ids[i] = strings.TrimPrefix(ids[i], "KISHOW:")
	}
	return c.JSONPretty(200, TopJSON{UUID: ids}, "  ")
}

// JSON API
func JSON(c echo.Context) error {
	d := kvs.GET("KISHOW:" + c.Param("uuid"))
	if len(d) == 0 {
		return c.String(200, "NOT FOUND")
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return c.String(200, d)
}
