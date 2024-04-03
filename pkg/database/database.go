package database

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var DefaultAlias = "default"

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
