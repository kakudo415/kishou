package websub

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"../kvs"

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
		c.AbortWithStatusJSON(404, gin.H{"error": "hub.mode error"})
	}

	if c.Query("hub.verify_token") != os.Getenv("JMA_VERIFY_TOKEN") {
		c.AbortWithStatusJSON(404, gin.H{"error": "hub.verify_token error"})
	}

	c.String(200, c.Query("hub.challenge"))
}

// Receiver func
func Receiver(c *gin.Context) {
	atom, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		c.AbortWithStatusJSON(404, gin.H{"error": "illegal atom feed"})
	}

	var feeds Feeds
	if err := xml.Unmarshal(atom, &feeds); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		c.AbortWithStatusJSON(404, gin.H{"error": "illegal atom feed"})
	}

	// Get more information
	for _, feed := range feeds {
		for _, entry := range feed.Entrys {
			res, err := http.Get(entry.Link.Href)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			defer res.Body.Close()
			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}

			// SET XML (RAW)
			ts := strconv.FormatInt(time.Now().Unix(), 10)
			kvs.SET(ts+"xml", string(data))
			println(ts)
		}
	}
}
