package websub

import (
	"os"

	"github.com/gin-gonic/gin"
)

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

}
