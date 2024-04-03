package config

import "time"

// CacheInterface cache interface
type CacheInterface interface {
	CommonInterface
	FuzzyDelete(keyPre string)
}

// Options 配置选项
type Options struct {
	FileName    string         // file name
	BasePath    string         // base path
	Ctype       string         // type of config file
	CachePrefix string         // cache prefix
	Cache       CacheInterface // Configure the register to customize the register, such as using redis You only need to implement the CacheInterface interface
}

// Interface cache interface
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
	GetStringMap(key string) map[string]any
}
