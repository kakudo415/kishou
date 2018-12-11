package kvs

import (
	"github.com/garyburd/redigo/redis"
)

var conn redis.Conn

func init() {
	conn = Connect()
}

// Connect to redis-server
func Connect() redis.Conn {
	c, e := redis.Dial("tcp", "127.0.0.1:6379")
	if e != nil {
		panic(e)
	}
	return c
}

// SET command
func SET(key string, value string) {
	conn.Do("SET", key, value)
}

// EXPIRE command
func EXPIRE(key string, ttl int) {
	conn.Do("EXPIRE", key, ttl)
}

// GET command
func GET(key string) string {
	rep, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return ""
	}
	return rep
}
