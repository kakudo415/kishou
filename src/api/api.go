package api

import (
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo"

	"../kvs"
)

// TopJSON Format
type TopJSON struct {
	UUID []string `json:"UUID"`
}

// Top API (Last 10min UUIDs)
func Top(c echo.Context) error {
	ids := kvs.KEYS("KISHOW:JSON:*")
	for i := 0; i < len(ids); i++ {
		ids[i] = strings.TrimPrefix(ids[i], "KISHOW:JSON:")
	}
	return c.JSON(200, TopJSON{UUID: ids})
}

// JSON API
func JSON(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	u, e := uuid.Parse(c.Param("uuid"))
	if e != nil {
		return c.NoContent(404)
	}
	d := kvs.GET("KISHOW:JSON:" + u.String())
	if len(d) == 0 {
		return c.NoContent(404)
	}
	return c.String(200, d)
}

// XML API
func XML(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)
	u, e := uuid.Parse(c.Param("uuid"))
	if e != nil {
		return c.NoContent(404)
	}
	d := kvs.GET("KISHOW:XML:" + u.String())
	if len(d) == 0 {
		return c.NoContent(404)
	}
	return c.String(200, d)
}
