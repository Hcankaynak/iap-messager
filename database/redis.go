package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type RedisConnection struct {
	host string
	port string
}

func connectRedis(connection RedisConnection) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", connection.host, connection.port),
	})

	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}
}
