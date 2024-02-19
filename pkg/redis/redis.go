package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Addr     string
	Password string
	DbNum    int
}

// New ...
func New(config Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DbNum,
	})
	err := client.Ping(context.Background()).Err()
	return client, err
}
