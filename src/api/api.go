package api

import (
	"../kvs"

	"github.com/gin-gonic/gin"
)

// XML API (RAW)
func XML(c *gin.Context) {
	rep := kvs.GET(c.Param("time") + "xml")
	if len(rep) == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "no data or illegal time"})
	}
	c.Header("Content-Type", "application/xml")
	c.String(200, rep)
}

// JSON API
func JSON(c *gin.Context) {
	rep := kvs.GET(c.Param("time") + "json")
	if len(rep) == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "no data or illegal time"})
	}
	c.Header("Content-Type", "application/json")
	c.String(200, rep)
}
