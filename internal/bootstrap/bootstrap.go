package bootstrap

import (
	"fmt"
	"go-skeleton/internal/variable"
	"go-skeleton/pkg/cache/redis"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/database/driver/mysql"
	"go-skeleton/pkg/helper"
	"go-skeleton/pkg/logger"
	mysql2 "gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log/slog"
	"time"
)

// Define service list
const (
	MysqlService = `Mysql`
	RedisService = `Redis`
)

type bootServiceMap map[string]func() error

var (
	BootedService []string
	err           error
	serviceMap    = bootServiceMap{
		MysqlService: initMysql,
		RedisService: initRedis,
	}
)

// BootService Load service
func BootService(services ...string) {
	if err = initConfig(); err != nil {
		panic("Failed to load config：" + err.Error())
	}
	if err = initLogger(); err != nil {
		panic("Failed to load logger：" + err.Error())
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
			variable.Log.Info("Loading " + k + " service successfully")
			BootedService = append(BootedService, k)
		}
	}
}

// keys ...
func (m bootServiceMap) keys() []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// initMysql ...
func initMysql() error {
	if variable.DB != nil {
		return nil
	}
	if variable.Config.GetBool("database.mysql.enabled") == false {
		return nil
	}
	dbContainer := func(dns string) *mysql.Mysql {
		return mysql.New(func(opts *mysql2.Config) {
			opts.DSN = dns
		})
	}
	masterDsn := variable.Config.GetString("database.mysql.master")
	d, err := database.New(
		dbContainer(masterDsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   variable.Config.GetString("database.mysql.prefix"),
				SingularTable: true, // 是否设置单数表名，设置为 是
			},
			Logger: logger2.Default.LogMode(logger2.LogLevel(variable.Config.GetInt("database.mysql.loglevel"))), // Todo
		},
		database.WithMaxIdleConn(variable.Config.GetInt("database.mysql.minIdleConn")),
		database.WithMaxOpenConn(variable.Config.GetInt("database.mysql.maxOpenConn")),
		database.WithConnMaxIdleTime(variable.Config.GetDuration("database.mysql.maxIdleTime")*time.Second),
		database.WithConnMaxLifetime(variable.Config.GetDuration("database.mysql.maxLifetime")*time.Minute),
	)
	if err != nil {
		return err
	}
	// write read seperate
	if variable.Config.GetBool("database.mysql.seperation") {
		var replicas []gorm.Dialector
		for _, slave := range variable.Config.GetStringSlice("database.mysql.slaves") {
			replicas = append(replicas, dbContainer(slave).Instance())
		}
		if err := d.WithSlaveDB([]gorm.Dialector{dbContainer(masterDsn).Instance()}, replicas); err != nil {
			return err
		}
	}
	variable.DB = d.DB
	return nil
}

// initRedis ...
func initRedis() error {
	if variable.Redis != nil {
		return nil
	}
	if variable.Config.GetBool("redis.enabled") == false {
		return nil
	}
	variable.Redis, err = redis.New(redis.Config{
		Addr:        fmt.Sprintf("%s:%s", variable.Config.GetString("redis.host"), variable.Config.GetString("redis.port")),
		Password:    variable.Config.GetString("redis.password"),
		DbNum:       variable.Config.GetInt("redis.dbname"),
		PoolSize:    variable.Config.GetInt("redis.poolSize"),
		MinIdleConn: variable.Config.GetInt("redis.minIdleConn"),
		MaxIdleConn: variable.Config.GetInt("redis.maxIdleConn"),
		MaxLifetime: variable.Config.GetDuration("redis.maxLifeTime") * time.Second,
		MaxIdleTime: variable.Config.GetDuration("redis.maxIdleTime") * time.Minute,
	})
	return err
}

// initConfig ...
func initConfig() error {
	var err error
	variable.Config, err = config.New(config.NewViper(), config.Options{
		BasePath: variable.BasePath,
		FileName: "config." + config.ConfEnv,
	})
	if err == nil {
		slog.Info("Loading configuration successfully")
	}
	return err
}

// initLogger ...
func initLogger() error {
	if variable.Log != nil {
		return nil
	}
	variable.Log = logger.New(
		variable.Config.GetString("log.dirPath"),
		variable.Config.GetString("log.fileName"),
		slog.Level(variable.Config.GetInt("log.level")),
	)
	variable.Log.Info(fmt.Sprintf("Loading logger successfully [ name：%s path：%s ]",
		variable.Config.GetString("log.fileName"),
		variable.Config.GetString("log.dirPath"),
	))
	return nil
}
