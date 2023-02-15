package db

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
)

var tracer = otel.Tracer("github.com/CyberAgentHack/server-performance-tuning-2023/pkg/db")

type RedisClient interface {
	Get(ctx context.Context, key string, dst any) (bool, error)
	Set(ctx context.Context, key string, v any, ttl time.Duration) error
}

type redisClient struct {
	client *redis.Client
}

func NewRedisClient(endpoint string) (RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr: endpoint,
	})

	return &redisClient{
		client: client,
	}, nil
}

func (c *redisClient) Get(ctx context.Context, key string, dst any) (bool, error) {
	ctx, span := tracer.Start(ctx, "db.redisClient#Get")
	defer span.End()

	b, err := c.client.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	if err != nil {
		return false, errcode.New(err)
	}
	err = gob.NewDecoder(bytes.NewBuffer(b)).Decode(dst)
	if err != nil {
		return false, errcode.New(err)
	}
	return true, nil
}

func (c *redisClient) Set(ctx context.Context, key string, v any, ttl time.Duration) error {
	ctx, span := tracer.Start(ctx, "db.redisClient#Set")
	defer span.End()

	buf := new(bytes.Buffer)
	err := gob.NewEncoder(buf).Encode(v)
	if err != nil {
		return errcode.New(err)
	}
	_, err = c.client.Set(ctx, key, buf.Bytes(), ttl).Result()
	return errcode.New(err)
}
