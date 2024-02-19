package bootstrap

import (
	"fmt"
	"go-skeleton/internal/variable"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/mysql"
	"go-skeleton/pkg/redis"
	"go-skeleton/pkg/utils"
	"log/slog"
)

// 定义服务列表
const (
	MysqlService = `Mysql`
	RedisService = `Redis`
)

type bootServiceMap map[string]func() error

// BootedService 已经加载的服务
var (
	BootedService []string
	err           error
	// serviceMap 程序启动时需要自动加载的服务
	serviceMap = bootServiceMap{
		MysqlService: bootMysql,
		RedisService: bootRedis,
	}
)

// BootService 加载服务
func BootService(services ...string) {
	if err = bootConfig(); err != nil {
		panic("初始化config配置失败：" + err.Error())
	}
	if err = bootLogger(); err != nil {
		panic("初始化log日志失败：" + err.Error())
	}
	if len(services) == 0 {
		services = serviceMap.keys()
	}
	BootedService = make([]string, 0)
	for k, val := range serviceMap {
		if utils.InAnySlice[string](services, k) {
			if err := val(); err != nil {
				panic("程序服务启动失败:" + err.Error())
			}
			BootedService = append(BootedService, k)
		}
	}
}

// keys 获取BootServiceMap中所有键值
func (m bootServiceMap) keys() []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// bootMysql 装配数据库连接
func bootMysql() error {
	if variable.DB != nil {
		return nil
	}
	if variable.Config.GetBool("mysql.enabled") == false {
		return nil
	}
	variable.DB, err = mysql.New(mysql.DatabaseConfig{
		Host:         variable.Config.GetString("mysql.write.host"),
		Port:         variable.Config.GetString("mysql.write.port"),
		User:         variable.Config.GetString("mysql.write.user"),
		Pass:         variable.Config.GetString("mysql.write.password"),
		DbName:       variable.Config.GetString("mysql.write.dbname"),
		Prefix:       variable.Config.GetString("mysql.write.prefix"),
		MaxIdleConns: variable.Config.GetInt("mysql.maxIdleConns"),
		MaxOpenConns: variable.Config.GetInt("mysql.maxOpenConns"),
		MaxLifeTime:  variable.Config.GetInt("mysql.maxLifeTime"),
	})
	if err == nil {
		variable.Log.Info("程序载入MySQL服务成功")
	}
	return err
}

// bootRedis 装配redis服务
func bootRedis() error {
	if variable.Redis != nil {
		return nil
	}
	if variable.Config.GetBool("redis.enabled") == false {
		return nil
	}
	variable.Redis, err = redis.New(redis.Config{
		Addr:     fmt.Sprintf("%s:%s", variable.Config.GetString("redis.host"), variable.Config.GetString("redis.port")),
		Password: variable.Config.GetString("redis.password"),
		DbNum:    variable.Config.GetInt("redis.dbname"),
	})
	if err == nil {
		variable.Log.Info("程序载入Redis服务成功")
	}
	return err
}

// bootConfig 初始化配置
func bootConfig() error {
	var err error
	variable.Config, err = config.New(config.NewViper(), config.Options{
		BasePath: variable.BasePath,
		Filename: "config." + config.ConfEnv,
	})
	if err == nil {
		slog.Info("初始化配置成功")
	}
	return err
}

// bootLogger ...
func bootLogger() error {
	if variable.Log != nil {
		return nil
	}
	if variable.Config.GetBool("log.enabled") == false {
		return nil
	}
	variable.Log = logger.New(
		variable.Config.GetString("log.dirPath"),
		variable.Config.GetString("log.fileName"),
		slog.Level(variable.Config.GetInt("log.level")),
	)
	variable.Log.Info(fmt.Sprintf("程序载入Logger服务成功 [ 日志名：%s 日志路径：%s ]",
		variable.Config.GetString("log.fileName"),
		variable.Config.GetString("log.dirPath"),
	))
	return nil
}
