package boots

import (
	"fmt"
	"time"

	"github.com/MQEnergy/go-skeleton/internal/vars"
	"github.com/MQEnergy/go-skeleton/pkg/cache/redis"
)

// InitRedis ...
func InitRedis() error {
	var err error
	if vars.Redis != nil {
		return nil
	}
	if vars.Config.GetBool("redis.enabled") == false {
		return nil
	}
	vars.Redis, err = redis.New(redis.Config{
		Addr:        fmt.Sprintf("%s:%s", vars.Config.GetString("redis.host"), vars.Config.GetString("redis.port")),
		Password:    vars.Config.GetString("redis.password"),
		DbNum:       vars.Config.GetInt("redis.dbname"),
		PoolSize:    vars.Config.GetInt("redis.poolSize"),
		MinIdleConn: vars.Config.GetInt("redis.minIdleConn"),
		MaxIdleConn: vars.Config.GetInt("redis.maxIdleConn"),
		MaxLifetime: vars.Config.GetDuration("redis.maxLifeTime") * time.Second,
		MaxIdleTime: vars.Config.GetDuration("redis.maxIdleTime") * time.Minute,
	})
	return err
}
