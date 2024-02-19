package smap

import (
	"strings"
	"sync"
)

var cache sync.Map

type SMap struct {
}

func New() *SMap {
	return &SMap{}
}

func (c *SMap) Get(key string) any {
	value, ok := cache.Load(key)
	if ok {
		return value
	}
	return nil
}

func (c *SMap) Set(key string, value any) bool {
	var res bool
	if exist := c.Has(key); exist == false {
		cache.Store(key, value)
		res = true
	}
	return res
}

func (c *SMap) Has(key string) bool {
	_, ok := cache.Load(key)
	return ok
}

// FuzzyDelete 会删除所有以prefix为前缀的配置项
func (c *SMap) FuzzyDelete(keyPre string) {
	cache.Range(func(key, value interface{}) bool {
		if key, ok := key.(string); ok {
			if strings.HasPrefix(key, keyPre) {
				cache.Delete(key)
			}
		}
		return true
	})
}
