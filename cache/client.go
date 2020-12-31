package cache

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Client struct {
	Pool *redis.Pool
}

var redisPool *redis.Pool

func Initialize(redisAddr string, maxConnections int) (Client, error) {
	fmt.Printf("Connecting to Redis at %s with max conn %d\n", redisAddr, maxConnections)
	client := Client{}

	redisPool = redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", redisAddr)
	}, maxConnections)

	client.Pool = redisPool

	fmt.Println("Redis connection pool initialized")
	return client, nil
}
