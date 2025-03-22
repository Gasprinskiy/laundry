package redisclient

import (
	"context"
	"laundry/config"
	"laundry/internal/entity/global"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisClient struct {
	db       *redis.Client
	duration time.Duration
}

func NewRedisClient(
	rdb *redis.Client,
	conf *config.Config,
) *RedisClient {
	return &RedisClient{
		db:       rdb,
		duration: time.Minute * time.Duration(conf.RedisTtl),
	}
}

func (rc *RedisClient) Set(key string, value []byte) error {
	cmd := rc.db.Set(ctx, key, value, rc.duration)

	if cmd.Err() != nil {
		return global.ErrInternalError
	}

	return nil
}

func (rc *RedisClient) Get(key string) (string, error) {
	val, err := rc.db.Get(ctx, key).Result()

	if err != nil {
		if err == redis.Nil {
			err = global.ErrNoData
		} else {
			err = global.ErrInternalError
		}
		return val, err
	}

	return val, nil
}
