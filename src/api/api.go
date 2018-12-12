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
	c.Header("Content-Type", "application/json; charset=utf-8")
	data := `{"body":{` + strings.Join(infos, `,`) + `}}`
	if c.Query("format") == "true" {
		var buf bytes.Buffer
		json.Indent(&buf, []byte(data), "", "  ")
		c.String(200, buf.String())
	} else {
		c.String(200, data)
	}
}
