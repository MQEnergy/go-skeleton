package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Addr        string
	Password    string
	DbNum       int
	PoolSize    int
	MaxIdleConn int
	MinIdleConn int
	MaxIdleTime time.Duration
	MaxLifetime time.Duration
}

// New ...
func New(config Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:            config.Addr,
		Password:        config.Password,
		DB:              config.DbNum,
		PoolSize:        config.PoolSize,
		MaxIdleConns:    config.MaxIdleConn,
		MinIdleConns:    config.MinIdleConn,
		ConnMaxIdleTime: config.MaxIdleTime,
		ConnMaxLifetime: config.MaxLifetime,
	})
	_, err := client.Ping(context.Background()).Result()
	return client, err
}
