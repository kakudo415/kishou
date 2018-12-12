package api

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"

	"../kvs"
)

// Top API (List of now Info)
func Top(c *gin.Context) {
	keys := []string{}
	for _, key := range kvs.KEYS("KISHOW-XML:*") {
		keys = append(keys, strings.TrimPrefix(key, "KISHOW-XML:"))
	}
	b, err := json.MarshalIndent(keys, "", "  ")
	if err != nil {
		c.AbortWithStatus(500)
	}
	c.Data(200, "application/json; charset=utf-8", b)
}

// XML API
func XML(c *gin.Context) {
	data := kvs.GET("KISHOW-XML:" + c.Param("uuid"))
	if len(data) == 0 {
		c.AbortWithStatus(404)
		return
	}
	c.Header("Content-Type", "application/xml; charset=utf-8")
	c.String(200, data)
}

// JSON API
func JSON(c *gin.Context) {
	data := kvs.GET("KISHOW-JSON:" + c.Param("uuid"))
	if len(data) == 0 {
		c.AbortWithStatus(404)
		return
	}
	c.Header("Content-Type", "application/json; charset=utf-8")
	if c.Query("format") == "true" {
		var buf bytes.Buffer
		json.Indent(&buf, []byte(data), "", "  ")
		c.String(200, buf.String())
	} else {
		c.String(200, data)
	}
}
