package websub

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"

	"github.com/labstack/echo"
)

// Sub - scriber
func Sub(c echo.Context) error {
	method := c.Request().Method

	// Subscribe / Unsubscribe
	if method == "GET" {
		if mode := c.QueryParam("hub.mode"); mode != "subscribe" && mode != "unsubscribe" {
			return c.String(404, "HUB MODE ERROR")
		}
		if c.QueryParam("hub.verify_token") != os.Getenv("JMA_VERIFY_TOKEN") {
			return c.String(404, "VERIFY TOKEN ERROR")
		}
		return c.String(200, c.QueryParam("hub.challenge"))
	}

	// Data receiver
	if method == "POST" {
		b := new(bytes.Buffer)
		b.ReadFrom(c.Request().Body)
		var src Tag
		if err := xml.Unmarshal(b.Bytes(), &src); err != nil {
			return c.String(404, "UNMARSHAL XML ERROR")
		}
		return c.JSONPretty(200, src, "  ")
	}

	return nil
}

// Tag struct for XML JSON
type Tag struct {
	Name     string
	Value    string
	Children []*Tag
}

// UnmarshalXML func
func (t *Tag) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	t.Name = start.Name.Local
	for {
		token, err := d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			var data *Tag
			if err := d.DecodeElement(&data, &tok); err != nil {
				return err
			}
			t.Children = append(t.Children, data)
		case xml.CharData:
			cd := string(token.(xml.CharData).Copy())
			if cd != "\n" {
				t.Value = cd
			}
		}
	}
}

// MarshalJSON func
func (t *Tag) MarshalJSON() ([]byte, error) {
	var result string
	dk := dupKeys(t.Children)
	for _, c := range t.Children {
		j, e := t.marshalJSON(c.Name, dk[c.Name])
		if e != nil {
			return []byte{}, e
		}
		result += j
	}
	return []byte(`[` + `]`), nil
}

func (t *Tag) marshalJSON(name string, dup bool) (string, error) {
	var result = `"` + name + `":`
	if dup {
		result += `[`
		for _, c := range t.Children {
			if c.Name == name {
				result += `{`
			}
		}
	} else {
		result += `{`
	}
	return result, nil
}

func dupKeys(t []*Tag) (m map[string]bool) {
	var d map[string]bool
	for _, v := range t {
		if d[v.Name] {
			m[v.Name] = true
		} else {
			d[v.Name] = true
		}
	}
	return m
}
