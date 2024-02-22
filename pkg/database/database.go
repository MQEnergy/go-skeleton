package database

import (
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

type Database struct {
	DB              *gorm.DB
	maxIdleConn     int
	maxOpenConn     int
	connMaxLifetime time.Duration
	connMaxIdleTime time.Duration
}

type DriverInterface interface {
	Instance() gorm.Dialector
}

type Options interface {
	apply(db *Database)
}

type OptionsFunc struct {
	fun func(db *Database)
}

func New(drive DriverInterface, config *gorm.Config, options ...Options) (*Database, error) {
	db, err := gorm.Open(drive.Instance(), config)
	if err != nil {
		return nil, err
	}
	dbContainer := &Database{
		DB: db,
	}
	for _, option := range options {
		option.apply(dbContainer)
	}
	// 新版本发现日志里出现大量record not found，屏蔽这里的日志 参考：https://github.com/go-gorm/gorm/issues/3789
	dbContainer.DB.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", func(d *gorm.DB) {
		d.Statement.RaiseErrorOnNotFound = false
	})

	s, err := dbContainer.DB.DB()
	if err != nil {
		return nil, err
	}
	s.SetMaxIdleConns(dbContainer.maxIdleConn)
	s.SetMaxOpenConns(dbContainer.maxOpenConn)
	s.SetConnMaxIdleTime(dbContainer.connMaxIdleTime)
	s.SetConnMaxLifetime(dbContainer.connMaxLifetime)
	return dbContainer, nil
}

// WithSlaveDB ...
func (d Database) WithSlaveDB(sources, replicas []gorm.Dialector) error {
	return d.DB.Use(
		dbresolver.Register(dbresolver.Config{
			Sources:  sources,
			Replicas: replicas,
			Policy:   dbresolver.RandomPolicy{},
		}).
			SetMaxIdleConns(d.maxIdleConn).
			SetMaxOpenConns(d.maxOpenConn).
			SetConnMaxIdleTime(d.connMaxIdleTime).
			SetConnMaxLifetime(d.connMaxLifetime),
	)
}

func (f OptionsFunc) apply(db *Database) {
	f.fun(db)
}

func WithMaxIdleConn(maxIdleConn int) Options {
	return OptionsFunc{
		fun: func(db *Database) {
			db.maxIdleConn = maxIdleConn
		},
	}
}

func WithMaxOpenConn(maxOpenConn int) Options {
	return OptionsFunc{
		fun: func(db *Database) {
			db.maxOpenConn = maxOpenConn
		},
	}
}

func WithConnMaxLifetime(connMaxLifetime time.Duration) Options {
	return OptionsFunc{
		fun: func(db *Database) {
			db.connMaxLifetime = connMaxLifetime
		},
	}
}

func WithConnMaxIdleTime(connMaxIdleTime time.Duration) Options {
	return OptionsFunc{
		fun: func(db *Database) {
			db.connMaxIdleTime = connMaxIdleTime
		},
	}
}
