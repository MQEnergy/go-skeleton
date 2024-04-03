package bootstrap

import (
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/gogf/gf/v2/util/gconv"

	logger2 "github.com/MQEnergy/go-skeleton/pkg/logger"
	"gorm.io/gorm/logger"

	"github.com/MQEnergy/go-skeleton/internal/vars"
	"github.com/MQEnergy/go-skeleton/pkg/cache/redis"
	"github.com/MQEnergy/go-skeleton/pkg/config"
	"github.com/MQEnergy/go-skeleton/pkg/database"
	"github.com/MQEnergy/go-skeleton/pkg/database/driver/mysql"
	"github.com/MQEnergy/go-skeleton/pkg/helper"

	mysql2 "gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		MysqlService: initMultiMysql,
		RedisService: initRedis,
	}
)

// BootService Load service
func BootService(services ...string) {
	// loading configuration
	if err := InitConfig(); err != nil {
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
	// load dao
	LoadDao()
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
		slog.Info("Server.mode: " + vars.Config.GetString("server.mode"))
		slog.Info("Loading Configuration successfully")
	}
	return err
}

// LoadDao 如果多数据库需要手动配置...
func LoadDao() {
	// dao default set
	// if vars.DB != nil {
	//	 dao.SetDefault(vars.DB)
	// }

	// 此处自行配置其他dao配置 ... 执行 go run cmd/cli/main.go genModel -a=demo 会生成daodemo目录
	// if vars.MDB["demo"] != nil {
	//	 daodemo.SetDefault(vars.MDB["demo"])
	// }
}

// initMultiMysql ...
func initMultiMysql() error {
	if vars.MDB != nil {
		return nil
	}
	sources := vars.Config.Get("database.mysql.sources")
	sourceList, ok := sources.([]interface{})
	if !ok {
		return nil
	}
	if len(sourceList) == 0 {
		return nil
	}
	if vars.Config.GetBool("database.mysql.enabled") == false {
		return nil
	}
	vars.MDB = make(map[string]*gorm.DB, len(sourceList))
	for _, m := range sourceList {
		sm := gconv.Map(m)
		alias := sm["alias"].(string)
		d, err := handleMysql(sm)
		if err != nil {
			slog.Error("Failed to start mysql connection err: ", err.Error())
			continue
		}
		if alias == database.DefaultAlias {
			vars.DB = d.DB
		}
		vars.MDB[alias] = d.DB
		slog.Info("Starting mysql connection db:" + alias)
	}
	return nil
}

// handleMysql ...
func handleMysql(sourceMaps map[string]interface{}) (*database.Database, error) {
	fileName := sourceMaps["filename"].(string)
	logLevel := sourceMaps["loglevel"].(int)
	masterDsn := sourceMaps["master"].(string)
	prefix := sourceMaps["prefix"].(string)
	seperation := sourceMaps["seperation"].(bool)
	slaves := sourceMaps["slave"].([]interface{})

	dbContainer := func(dns string) *mysql.Mysql {
		return mysql.New(func(opts *mysql2.Config) {
			opts.DSN = dns
		})
	}
	newLogger := logger.New(
		log.New(logger2.ApplyWriter(
			fileName,
			vars.Config,
		), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,               // Slow SQL threshold
			LogLevel:                  logger.LogLevel(logLevel), // Log level
			IgnoreRecordNotFoundError: true,                      // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,                      // Don't include params in the SQL log
			Colorful:                  false,                     // Disable color
		},
	)
	d, err := database.New(
		dbContainer(masterDsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   prefix,
				SingularTable: true, // 是否设置单数表名，设置为 是
			},
			Logger: newLogger,
		},
		database.WithMaxIdleConn(vars.Config.GetInt("database.mysql.minIdleConn")),
		database.WithMaxOpenConn(vars.Config.GetInt("database.mysql.maxOpenConn")),
		database.WithConnMaxIdleTime(vars.Config.GetDuration("database.mysql.maxIdleTime")*time.Second),
		database.WithConnMaxLifetime(vars.Config.GetDuration("database.mysql.maxLifetime")*time.Minute),
	)
	if err != nil {
		return nil, err
	}
	// write read seperate
	if seperation {
		var replicas []gorm.Dialector
		for _, slave := range slaves {
			replicas = append(replicas, dbContainer(slave.(string)).Instance())
		}
		if err := d.WithSlaveDB([]gorm.Dialector{dbContainer(masterDsn).Instance()}, replicas); err != nil {
			return nil, err
		}
	}
	return d, nil
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
