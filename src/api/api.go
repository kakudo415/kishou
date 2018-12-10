package api

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// JSON API
func JSON(c *gin.Context) {
	ts, err := strconv.ParseInt(c.Param("ts"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "illegal time"})
	}

	t := time.Unix(ts, 0)

	println(t.String())
}
