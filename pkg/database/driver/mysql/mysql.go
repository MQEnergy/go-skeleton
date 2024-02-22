package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Option func(opts *mysql.Config)

type Mysql struct {
	instance gorm.Dialector
	Config   mysql.Config
}

func New(opts ...Option) *Mysql {
	mysqlContainer := &Mysql{}
	for _, f := range opts {
		f(&mysqlContainer.Config)
	}
	mysqlContainer.instance = mysql.New(mysqlContainer.Config)
	return mysqlContainer
}

func (m *Mysql) Instance() gorm.Dialector {
	return m.instance
}
