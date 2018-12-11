package websub

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"

	"../kvs"
)

// Feeds XML
type Feeds []struct {
	Entrys []struct {
		Link struct {
			Href string `xml:"href,attr"`
		} `xml:"link"`
	} `xml:"entry"`
}

// Subscriber (subscribe / unsubscribe)
func Subscriber(c *gin.Context) {
	if c.Query("hub.mode") != "subscribe" && c.Query("hub.mode") != "unsubscribe" {
		fmt.Fprintln(os.Stderr, "[hub.mode error] "+c.Query("hub.mode"))
		c.AbortWithStatusJSON(404, gin.H{"error": "illegal hub.mode"})
		return
	}

	if c.Query("hub.verify_token") != os.Getenv("JMA_VERIFY_TOKEN") {
		fmt.Fprintln(os.Stderr, "[verify_token error] "+c.Query("hub.verify_token"))
		c.AbortWithStatusJSON(404, gin.H{"error": "illegal hub.verify_token"})
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
		c.AbortWithStatusJSON(404, gin.H{"error": "illegal atom feed"})
		return
	}
	// Get more information
	for _, item := range atom.Items {
		resp, err := http.Get(item.Link)
		if err != nil {
			continue
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}
		var info interface{}
		err = xml.Unmarshal(data, &info)
		if err != nil {
			continue
		}
		// Save to KVS
		id := UUID()
		save("KISHOW-XML:"+id, `"`+id+`":"`+removeFromString(strings.Replace(string(data), `"`, `\"`, -1), "\n")+`"`) // Escape '"' and Remove '\n'
		save("KISHOW-JSON:"+id, `{"`+id+`":""}`)
	}
}

// UUID gen
func UUID() string {
	return uuid.New().String()
}

func save(key string, value string) {
	kvs.SET(key, value)
	kvs.EXPIRE(key, 600)
}

func removeFromString(s string, r ...string) string {
	for _, o := range r {
		s = strings.Replace(s, o, "", -1)
	}
	return s
}
