package boots

import (
	"log"
	"log/slog"
	"time"

	"github.com/MQEnergy/go-skeleton/internal/vars"
	"github.com/MQEnergy/go-skeleton/pkg/database"
	"github.com/MQEnergy/go-skeleton/pkg/database/driver/mysql"
	logger2 "github.com/MQEnergy/go-skeleton/pkg/logger"
	"github.com/gogf/gf/v2/util/gconv"
	mysql2 "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// InitMultiMysql ...
func InitMultiMysql() error {
	if vars.MDB != nil {
		return nil
	}
	if vars.Config.GetBool("database.mysql.enabled") == false {
		return nil
	}
	sources := vars.Config.Get("database.mysql.sources")
	sourceList, ok := sources.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(sourceList) == 0 {
		return nil
	}
	vars.MDB = make(map[string]*gorm.DB, len(sourceList))
	for alias, m := range sourceList {
		sm := gconv.Map(m)
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
