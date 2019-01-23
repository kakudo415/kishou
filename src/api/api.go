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
	if c.Request().Header.Get("X-JSON-STYLE") == "PRETTY" {
		return c.JSONPretty(200, TopJSON{UUID: ids}, "  ")
	}
	return c.JSON(200, TopJSON{UUID: ids})
}

// JSON API
func JSON(c echo.Context) error {
	var info string
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	if c.Request().Header.Get("X-JSON-STYLE") == "PRETTY" {
		info = kvs.GET("KISHOW-PRETTY:" + c.Param("uuid"))
	} else {
		info = kvs.GET("KISHOW:" + c.Param("uuid"))
	}
	return c.String(200, info)
}
