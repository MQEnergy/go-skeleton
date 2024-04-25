package bootstrap

import (
	"fmt"
	"log/slog"

	"github.com/MQEnergy/go-skeleton/internal/bootstrap/boots"
	"github.com/MQEnergy/go-skeleton/pkg/helper"
)

// Define service list
const (
	MysqlService = `Mysql`
	RedisService = `Redis`
)

type bootServiceMap map[string]func() error

var (
	BootedService []string
	serviceMap    = bootServiceMap{
		MysqlService: boots.InitMultiMysql,
		RedisService: boots.InitRedis,
	}
)

// BootService Load service
func BootService(services ...string) {
	// loading configuration
	if err := boots.InitConfig(); err != nil {
		panic("Failed to load config：" + err.Error())
	}
	if len(services) == 0 {
		services = serviceMap.keys()
	}
	BootedService = make([]string, 0)
	for k, val := range serviceMap {
		if helper.InAnySlice[string](services, k) {
			if err := val(); err != nil {
				panic(fmt.Sprintf("Failed to load service %s err: %s", k, err.Error()))
			}
			slog.Info("Loading " + k + " service successfully")
			BootedService = append(BootedService, k)
		}
	}
	boots.InitLogger()
	// load dao
	boots.InitDao()
}

// keys ...
func (m bootServiceMap) keys() []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
