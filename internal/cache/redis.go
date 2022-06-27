package cache

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

type Config struct {
	Addr     string
	Password string
	DB       int
}

func InitRedis(cfg Config) error {
	tmp := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if tmp == nil {
		return errors.New("redis init fail")
	}
	rdb = tmp
	return nil
}

func SetKeyValue(ctx context.Context, key, value string) error {
	timeout := rand.Intn(100) + 200
	return rdb.Set(ctx, key, value, time.Duration(timeout)*time.Second).Err()
}

func GetKey(ctx context.Context, key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}
