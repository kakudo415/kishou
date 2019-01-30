package websub

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/mmcdole/gofeed"

	"../db"
	"../kvs"
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
		fp := gofeed.NewParser()
		feed, err := fp.Parse(c.Request().Body)
		if err != nil {
			return nil
		}
		for _, item := range feed.Items {
			if !strings.HasPrefix(item.Link, "http://xml.kishou.go.jp/") {
				fmt.Fprintf(os.Stderr, "[ERROR] 気象庁以外のリンク %s\n", item.Link)
				continue
			}
			res, _ := http.Get(item.Link)
			// XML => JSON
			var src Tag
			b := bytes.NewBuffer([]byte{})
			b.ReadFrom(res.Body)
			err := xml.Unmarshal(b.Bytes(), &src)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[ERROR] %s\n", err.Error())
				continue
			}
			// Save to MySQL
			id, err := uuid.Parse(strings.TrimPrefix(item.GUID, "urn:uuid:"))
			if err != nil {
				continue
			}
			j, err := json.Marshal(&src)
			if err != nil {
				continue
			}
			db.Add(id, time.Now(), string(j), b.String())
			// Save to Redis (Latest keys)
			k := "KISHOW:" + id.String()
			kvs.SET(k, "RECEIVED")
			kvs.EXPIRE(k, (time.Minute * 10))
		}
		return c.String(200, "THANK YOU")
	}
	return nil
}

// Tag struct for XML JSON
type Tag struct {
	Name     string
	Value    string
	Children []*Tag
}

// MarshalJSON func
func (t *Tag) MarshalJSON() ([]byte, error) {
	result, err := innerJSON(t.Name, []*Tag{t})
	if err != nil {
		return []byte{}, err
	}
	result = `{` + result + `}`
	return []byte(result), nil
}

func innerJSON(name string, elms []*Tag) (string, error) {
	result := `"` + name + `":`
	if len(elms) >= 2 {
		result += `[`
	}
	for i, elm := range elms {
		if i >= 1 {
			result += `,`
		}
		if len(elm.Children) == 0 {
			result += `"` + elm.Value + `"`
		} else {
			result += `{`
			dupKeys := map[string]bool{}
			for i, ec := range elm.Children {
				if dupKeys[ec.Name] {
					continue
				}
				inner, err := innerJSON(ec.Name, sameKeys(ec.Name, elm.Children))
				dupKeys[ec.Name] = true
				if err != nil {
					return result, err
				}
				if i >= 1 {
					result += `,`
				}
				result += inner
			}
			result += `}`
		}
	}
	if len(elms) >= 2 {
		result += `]`
	}
	return result, nil
}

func sameKeys(n string, t []*Tag) []*Tag {
	l := []*Tag{}
	for _, v := range t {
		if v.Name == n {
			l = append(l, v)
		}
	}
	return l
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
			t.Value = strings.Replace(cd, "\n", "", -1)
		}
	}
}
