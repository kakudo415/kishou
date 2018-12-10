package page

import (
	"github.com/gin-gonic/gin"
)

// Index page
func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

// JSON page
func JSON(c *gin.Context) {
	c.HTML(200, "json.html", gin.H{})
}
