package page

import (
	"github.com/gin-gonic/gin"
)

// Index page
func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
