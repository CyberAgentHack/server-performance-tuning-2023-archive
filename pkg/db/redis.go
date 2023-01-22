package db

import (
	"time"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/go-redis/redis"
)

type RedisClient interface {
	Set(key, value string, ttl time.Duration) error
}

type redisClient struct {
	client *redis.Client
}

func NewRedisClient(endpoint string) (RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr: endpoint,
	})

	return redisClient{
		client: client,
	}, nil
}

func (r redisClient) Set(key, value string, ttl time.Duration) error {
	err := r.client.Set(key, value, ttl).Err()
	return errcode.New(err)
}
