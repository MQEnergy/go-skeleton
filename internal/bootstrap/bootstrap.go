package bootstrap

import (
	"fmt"
	"log/slog"
	"time"

	"go-skeleton/internal/vars"

	"go-skeleton/pkg/cache/redis"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/database/driver/mysql"
	"go-skeleton/pkg/helper"

	mysql2 "gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
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
}

// keys ...
func (m bootServiceMap) keys() []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// InitConfig ...
func InitConfig() error {
	var err error
	vars.Config, err = config.New(config.NewViper(), config.Options{
		BasePath: vars.BasePath,
		FileName: "config." + config.ConfEnv,
	})
	if err == nil {
		slog.Info("Loading Configuration successfully")
	}
	return err
}

// initMysql ...
func initMysql() error {
	if vars.DB != nil {
		return nil
	}
	if vars.Config.GetBool("database.mysql.enabled") == false {
		return nil
	}
	dbContainer := func(dns string) *mysql.Mysql {
		return mysql.New(func(opts *mysql2.Config) {
			opts.DSN = dns
		})
	}
	masterDsn := vars.Config.GetString("database.mysql.master")
	d, err := database.New(
		dbContainer(masterDsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   vars.Config.GetString("database.mysql.prefix"),
				SingularTable: true, // 是否设置单数表名，设置为 是
			},
			Logger: logger2.Default.LogMode(logger2.LogLevel(vars.Config.GetInt("database.mysql.loglevel"))), // Todo
		},
		database.WithMaxIdleConn(vars.Config.GetInt("database.mysql.minIdleConn")),
		database.WithMaxOpenConn(vars.Config.GetInt("database.mysql.maxOpenConn")),
		database.WithConnMaxIdleTime(vars.Config.GetDuration("database.mysql.maxIdleTime")*time.Second),
		database.WithConnMaxLifetime(vars.Config.GetDuration("database.mysql.maxLifetime")*time.Minute),
	)
	if err != nil {
		return err
	}
	// write read seperate
	if vars.Config.GetBool("database.mysql.seperation") {
		var replicas []gorm.Dialector
		for _, slave := range vars.Config.GetStringSlice("database.mysql.slaves") {
			replicas = append(replicas, dbContainer(slave).Instance())
		}
		if err := d.WithSlaveDB([]gorm.Dialector{dbContainer(masterDsn).Instance()}, replicas); err != nil {
			return err
		}
	}
	vars.DB = d.DB
	return nil
}

// initRedis ...
func initRedis() error {
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
