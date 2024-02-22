package smap

import (
	"strings"
	"sync"
)

type SMap struct {
	cache sync.Map
}

func New() *SMap {
	return &SMap{
		cache: sync.Map{},
	}
}

func (c *SMap) Get(key string) any {
	value, ok := c.cache.Load(key)
	if ok {
		return value
	}
	return nil
}

func (c *SMap) Set(key string, value any) bool {
	var res bool
	if exist := c.Has(key); exist == false {
		c.cache.Store(key, value)
		res = true
	}
	return res
}

func (c *SMap) Has(key string) bool {
	_, ok := c.cache.Load(key)
	return ok
}

// FuzzyDelete 会删除所有以prefix为前缀的配置项
func (c *SMap) FuzzyDelete(keyPre string) {
	c.cache.Range(func(key, value interface{}) bool {
		if key, ok := key.(string); ok {
			if strings.HasPrefix(key, keyPre) {
				c.cache.Delete(key)
			}
		}
		return true
	})
}
