package config

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Address  string
	Username string
	Password string
}

var DefaultTTL = time.Duration(24 * time.Hour)

func InitRedis(params *RedisConfig) (client *redis.Client, err error) {
	for i := 0; i < 10; i++ {
		client = redis.NewClient(&redis.Options{
			Addr:     params.Address,
			Username: params.Username,
			Password: params.Password,
		})

		_, err = client.Ping(context.Background()).Result()
		if err == nil {
			log.Println("[InitRedis] Init redis succefull")
			break
		}

		log.Printf("[InitRedis] error init redis: %+v, retrying in 1 second\n", err)
		time.Sleep(time.Duration(time.Second))
	}

	return
}
