package websub

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"../kvs"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
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
	atom, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		c.AbortWithStatusJSON(404, gin.H{"error": "illegal atom feed"})
		return
	}

	var feeds Feeds
	if err := xml.Unmarshal(atom, &feeds); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		c.AbortWithStatusJSON(404, gin.H{"error": "illegal atom feed"})
		return
	}

	// Get more information
	for _, feed := range feeds {
		for _, entry := range feed.Entrys {
			res, err := http.Get(entry.Link.Href)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			defer res.Body.Close()

			// Save to KVS
			id := "KISHOW" + uuid.New().String()
			kvs.SET(id+"JSON", "")
			kvs.EXPIRE(id+"JSON", 600)
			fmt.Println(id + "JSON")
		}
	}
}
