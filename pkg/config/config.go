package config

import (
	"sync"
	"time"

	"go-skeleton/pkg/cache/smap"

	"github.com/spf13/cast"
)

var ConfEnv string

// CommonInterface ...
type CommonInterface interface {
	Get(key string) any
	Set(key string, value any) bool
	Has(key string) bool
}

type Config struct {
	viper CommonInterface
	cache CacheInterface
	mu    *sync.Mutex
}

var _ Interface = (*Config)(nil)

// ============================= Config Cache ====================================

func New(config CommonInterface, option Options) (provider *Config, err error) {
	if option.Cache == nil {
		option.Cache = smap.New()
	}
	if option.CachePrefix == "" {
		option.CachePrefix = "config"
	}
	if option.Ctype == "" {
		option.Ctype = "yaml"
	}
	if d, ok := config.(interface {
		Apply(Options) error
	}); ok {
		if err = d.Apply(option); err != nil {
			return
		}
	}
	return &Config{
		viper: config,
		cache: option.Cache,
		mu:    &sync.Mutex{},
	}, nil
}

func (c *Config) _cache(key string, value any) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cache.Has(key) {
		return true
	}
	return c.cache.Set(key, value)
}

func (c *Config) Get(key string) any {
	if c.cache.Has(key) {
		return c.cache.Get(key)
	}
	val := c.viper.Get(key)
	c._cache(key, val)
	return val
}

func (c *Config) GetString(key string) string {
	return cast.ToString(c.Get(key))
}

func (c *Config) GetBool(key string) bool {
	return cast.ToBool(c.Get(key))
}

func (c *Config) GetInt(key string) int {
	return cast.ToInt(c.Get(key))
}

func (c *Config) GetInt32(key string) int32 {
	return cast.ToInt32(c.Get(key))
}

func (c *Config) GetInt64(key string) int64 {
	return cast.ToInt64(c.Get(key))
}

func (c *Config) GetFloat64(key string) float64 {
	return cast.ToFloat64(c.Get(key))
}

func (c *Config) GetDuration(key string) time.Duration {
	return cast.ToDuration(c.Get(key))
}

func (c *Config) GetStringSlice(key string) []string {
	return cast.ToStringSlice(c.Get(key))
}

// ============================= Config Cache ====================================
