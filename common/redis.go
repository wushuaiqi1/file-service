package common

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var RedisInstance *redis.Client

func InitRedis() error {
	ctx := context.Background()
	RedisInstance = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	res, err := RedisInstance.Get(ctx, "health").Result()
	if err != nil {
		return err
	}
	log.Println("Redis Service Health:", res)
	return nil
}
