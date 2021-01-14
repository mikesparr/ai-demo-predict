package cache

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Client contains the redis pool for injection
type Client struct {
	Pool *redis.Pool
}

var redisPool *redis.Pool

// Initialize establishes the cache connection
func Initialize(redisAddr string, maxConnections int) (Client, error) {
	fmt.Printf("Connecting to Redis at %s with max conn %d\n", redisAddr, maxConnections)
	client := Client{}

	redisPool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisAddr)
		},
	}

	client.Pool = redisPool

	fmt.Println("Redis connection pool initialized")
	return client, nil
}
