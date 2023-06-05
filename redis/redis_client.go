package redis_client

import (
	"context"
	"log"
	"pool-pay/internal/constants"
	"pool-pay/internal/util"
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

	log.Printf("Redis ping response: %s\n", pong)
}

func SetIntoRedis(validToken string, email string) error {
	err := Client.Set(Ctx, validToken, email, time.Minute*60).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetFromRedis(key string) (string, error) {
	val, err := Client.Get(Ctx, key).Result()
	if err != nil {
		log.Println(err)

		if err == redis.Nil {
			return "", util.SetApiError(constants.ERRORCODE_KEYNOTFOUND)
		}
		return "", util.SetDefaultApiError(err)
	}

	return val, nil
}

func RefreshExpiredTime(key string) error {
	err := Client.Expire(Ctx, key, time.Minute*60).Err()
	if err != nil {
		return err
	}
	return nil
}
