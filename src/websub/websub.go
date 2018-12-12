package websub

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"

	"../kvs"
)

var escapeNL *regexp.Regexp

// Tag for Unmarshal
type Tag struct {
	Name string `json:"name"`
	// Attr     []Attr        `json:"attr"`
	Value    string        `json:"value"`
	Children []interface{} `json:"children"`
}

// Attr for Unmarshal
type Attr struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func init() {
	escapeNL = regexp.MustCompile(`(\n|\r|\r\n)`)
}

// Subscriber (subscribe / unsubscribe)
func Subscriber(c *gin.Context) {
	if c.Query("hub.mode") != "subscribe" && c.Query("hub.mode") != "unsubscribe" {
		fmt.Fprintln(os.Stderr, "[hub.mode error] "+c.Query("hub.mode"))
		c.String(404, "NOT FOUND")
		return
	}

	if c.Query("hub.verify_token") != os.Getenv("JMA_VERIFY_TOKEN") {
		fmt.Fprintln(os.Stderr, "[verify_token error] "+c.Query("hub.verify_token"))
		c.String(404, "NOT FOUND")
		return
	}

	c.String(200, c.Query("hub.challenge"))
}

// Receiver func
func Receiver(c *gin.Context) {
	fp := gofeed.NewParser()
	atom, err := fp.Parse(c.Request.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		c.String(404, "NOT FOUND")
		return
	}

	// Get more information
	for _, item := range atom.Items {
		if !strings.HasPrefix(item.Link, `http://xml.kishou.go.jp/`) {
			continue
		}
		resp, err := http.Get(item.Link)
		if err != nil {
			continue
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}
		var info *Tag
		err = xml.NewDecoder(bytes.NewReader(data)).Decode(&info)
		if err != nil {
			continue
		}

		// Save to KVS
		id := UUID()
		escapedXML := strings.Replace(escapeNL.ReplaceAllString(string(data), ``), `"`, `\"`, -1)
		save("KISHOW-XML:"+id, `"`+id+`":"`+escapedXML+`"`)
		escapedJSON, err := json.Marshal(info)
		if err != nil {
			continue
		}
		save("KISHOW-JSON:"+id, `"`+id+`":`+string(escapedJSON))
	}
}

// UUID gen
func UUID() string {
	return uuid.New().String()
}

func save(key string, value string) {
	kvs.SET(key, value)
	kvs.EXPIRE(key, 180)
}

// UnmarshalXML func
func (t *Tag) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	t.Name = start.Name.Local
	// for _, attr := range start.Attr {
	// 	t.Attr = append(t.Attr, Attr{attr.Name.Space + attr.Name.Local, attr.Value})
	// }
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
				t.Value = escapeNL.ReplaceAllString(cd, "")
			}
		}
	}
}
