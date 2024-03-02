package admin

import "gorm.io/gen"

type Querier interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(id int) (gen.T, error)

	// SELECT * FROM @@table WHERE name = @name
	GetByName(name string) (gen.T, error)
}
