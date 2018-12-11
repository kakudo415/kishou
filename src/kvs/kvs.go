package kvs

import (
	"github.com/garyburd/redigo/redis"
)

// Conn redis
var Conn redis.Conn

func init() {
	Conn = Connect()
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
	Conn.Do("SET", key, value)
}

// EXPIRE command
func EXPIRE(key string, ttl int) {
	Conn.Do("EXPIRE", key, ttl)
}

// KEYS command
func KEYS(pattern string) []string {
	rep, err := redis.Strings(Conn.Do("KEYS", pattern))
	if err != nil {
		return []string{""}
	}
	return rep
}

// GET command
func GET(key string) string {
	rep, err := redis.String(Conn.Do("GET", key))
	if err != nil {
		return ""
	}
	return rep
}

// GETALL command
func GETALL(pattern string) (reps []string) {
	keys := KEYS(pattern)
	for _, key := range keys {
		rep, err := redis.String(Conn.Do("GET", key))
		if err != nil {
			continue
		}
		reps = append(reps, rep)
	}
	return reps
}
