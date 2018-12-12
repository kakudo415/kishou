package api

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"

	"../kvs"
)

// XML API (RAW)
func XML(c *gin.Context) {
	infos := kvs.GETALL("KISHOW-XML:*")
	serve(c, infos)
}

// JSON API
func JSON(c *gin.Context) {
	infos := kvs.GETALL("KISHOW-JSON:*")
	serve(c, infos)
}

func serve(c *gin.Context, infos []string) {
	if len(infos) == 0 {
		c.String(404, "NOT FOUND")
		return
	}
	c.Header("Content-Type", "application/json; charset=utf-8")
	data := `{"body":{` + strings.Join(infos, `,`) + `}}`
	if c.Query("format") == "true" {
		var buf bytes.Buffer
		if json.Indent(&buf, []byte(data), "", "  ") != nil {
			c.String(404, "NOT FOUND")
			return
		}
		c.String(200, buf.String())
	} else {
		c.String(200, data)
	}
}
