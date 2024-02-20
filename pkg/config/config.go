package config

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go-skeleton/pkg/cache/smap"
	"strings"
	"sync"
	"time"
)

var (
	ConfEnv        string = "dev"
	lastChangeTime time.Time
)

func init() {
	lastChangeTime = time.Now()
}

// ContainerInterface 容器相关方法
type ContainerInterface interface {
	Get(key string) any
	Set(key string, value any) bool
	Has(key string) bool
}

type ViperInterface interface {
	ContainerInterface
}

// CacheInterface config cache
type CacheInterface interface {
	ContainerInterface
	FuzzyDelete(key string)
}

// Options 配置选项
type Options struct {
	// 文件名称
	Filename string

	// 工作目录，项目根目录
	BasePath string

	// 配置文件类型
	Ctype string

	// 配置缓存前缀
	CachePrefix string

	// 配置缓存器，可自定义缓存器，比如使用redis
	// 只需要实现CacheInterface接口即可
	Cache CacheInterface
}

// Interface 缓存接口
// 应用到redis、cache等cache db
type Interface interface {
	Get(key string) any
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetFloat64(key string) float64
	GetDuration(key string) time.Duration
	GetStringSlice(key string) []string
}

var _ Interface = (*Config)(nil)
var _ ViperInterface = (*ViperConfig)(nil)

// ============================= Viper Config ====================================

type ViperConfig struct {
	viper  *viper.Viper
	option Options
}

func NewViper() *ViperConfig {
	return &ViperConfig{}
}

func (v *ViperConfig) Get(key string) any {
	return v.viper.Get(key)
}

func (v *ViperConfig) Set(key string, value any) bool {
	v.viper.Set(key, value)
	return true
}

func (v *ViperConfig) Has(key string) bool {
	return v.viper.IsSet(key)
}

// Apply 创建实例
func (v *ViperConfig) Apply(option Options) error {
	v.option = option
	viperConfig := viper.New()
	viperConfig.AddConfigPath(option.BasePath + "/configs")
	if strings.TrimSpace(option.Filename) == "" {
		viperConfig.SetConfigName("config")
	} else {
		viperConfig.SetConfigName(option.Filename)
	}
	viperConfig.SetConfigType(option.Ctype)
	if err := viperConfig.ReadInConfig(); err != nil {
		return err
	}
	v.viper = viperConfig
	return nil
}

// ============================= Viper Config ====================================

// ============================= Config Cache ====================================

type Config struct {
	Viper ViperInterface
	cache CacheInterface
	mu    *sync.Mutex
}

func New(config ViperInterface, option Options) (provider *Config, err error) {
	if option.Cache == nil {
		option.Cache = smap.New()
	}
	if option.CachePrefix == "" {
		option.CachePrefix = "config"
	}
	if option.Ctype == "" {
		option.Ctype = "yaml"
	}
	if d, ok := config.(interface{ Apply(Options) error }); ok {
		if err = d.Apply(option); err != nil {
			return
		}
	}
	return &Config{
		Viper: config,
		cache: option.Cache,
		mu:    &sync.Mutex{},
	}, nil
}

func (c *Config) Cache(key string, value any) bool {
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
	val := c.Viper.Get(key)
	c.Cache(key, val)
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
