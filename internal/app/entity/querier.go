package entity

import "gorm.io/gen"

type Querier interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(id int) (gen.T, error)

	// SELECT * FROM @@table
	FindAll() ([]gen.T, error)

	// SELECT * FROM @@table LIMIT 1
	FindOne() gen.T
}
