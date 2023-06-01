package redis_client

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

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

func SetIntoRedis(validToken string, email string) error {
	err := Client.Set(Ctx, validToken, email, time.Second*60).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetFromRedis(key string) (string, error) {
	val, err := Client.Get(Ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("key not found")
	} else if err != nil {
		return "", err
	}

	return val, nil
}
