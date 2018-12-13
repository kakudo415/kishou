package websub

import (
	"github.com/labstack/echo"
)

// Sub - scriber
func Sub(c echo.Context) error {
	return c.String(200, "SUBSCRIBER")
}
