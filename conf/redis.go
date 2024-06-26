package conf

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

var rdClient *redis.Client

type RedisClient struct {
}

func InitRedis() (*RedisClient, error) {
	rdClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("ur;"),
		Password: "",
		DB:       0,
	})

	_, err := rdClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{}, nil
}

func (rc *RedisClient) Set(key string, value any, duration time.Duration) error {
	return rdClient.Set(context.Background(), key, value, duration).Err()
}

func (rc *RedisClient) Get(key string) (any, error) {
	return rdClient.Get(context.Background(), key).Result()
}

func (rc *RedisClient) Delete(key ...string) error {
	return rdClient.Del(context.Background(), key...).Err()
}

func (rc *RedisClient) GetExpireDuration(key string) (time.Duration, error) {
	return rdClient.TTL(context.Background(), key).Result()
}
