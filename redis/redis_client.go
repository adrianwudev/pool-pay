package redis_client

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6000",
		Password: "",
		DB:       0,
	})

	checkRedisConnection(Client, Ctx)
}

func checkRedisConnection(client *redis.Client, ctx context.Context) {
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Printf("failed to ping Redis: %s\n", err)
	}

	fmt.Printf("Redis ping response: %s\n", pong)
}
