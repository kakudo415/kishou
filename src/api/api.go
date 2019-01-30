package api

import (
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo"

	"../db"
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
	if c.QueryParam("s") == "p" {
		return c.JSONPretty(200, TopJSON{UUID: ids}, "  ")
	}
	return c.JSON(200, TopJSON{UUID: ids})
}

// JSON API
func JSON(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	u, e := uuid.Parse(c.Param("uuid"))
	if e != nil {
		return c.NoContent(404)
	}
	d := db.Get(u)
	if d.UUID == uuid.Nil {
		return c.NoContent(404)
	}
	if c.QueryParam("s") == "p" {
		return c.String(200, d.JSONP)
	}
	return c.String(200, d.JSONM)
}
