package redis

import (
	"context"
	"time"

	"server/core/infra/repository"
	"server/infrastructure/env"

	"github.com/redis/go-redis/v9"
)

var _ repository.IMemoryRepository = &MemoryRepository{}

type MemoryRepository struct {
	rdb *redis.Client
}

func NewMemoryRepository() (*MemoryRepository, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     env.GetEnv(env.RedisHost) + ":" + env.GetEnv(env.RedisPort),
		Password: env.GetEnv(env.RedisPass),
		DB:       0, // use default DB
	})
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return &MemoryRepository{
		rdb: rdb,
	}, nil
}

func (mr *MemoryRepository) Get(key string) (*[]byte, error) {
	value, err := mr.rdb.Get(context.Background(), key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &value, err
}

func (mr *MemoryRepository) Set(key string, value []byte, expire time.Duration) error {
	err := mr.rdb.Set(context.Background(), key, value, expire).Err()
	return err
}

func (mr *MemoryRepository) Delete(key string) error {
	err := mr.rdb.Del(context.Background(), key).Err()
	return err
}
