package redis_client

import (
	"context"

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
}
