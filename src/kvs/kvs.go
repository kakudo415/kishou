package kvs

import (
	"time"

	"github.com/go-redis/redis"
)

var conn *redis.Client

// Connect to Redis
func Connect() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
}

// SET COMMAND
func SET(k string, v string) {
	conn.Set(k, v, 0)
}

// EXPIRE COMMAND
func EXPIRE(k string, t time.Duration) {
	conn.Expire(k, t)
}

// GET COMMAND
func GET(k string) string {
	r, _ := conn.Get(k).Result()
	return r
}

// KEYS COMMAND
func KEYS(kr string) (v []string) {
	r, _ := conn.Keys(kr).Result()
	return r
}

func init() {
	conn = Connect()
}
