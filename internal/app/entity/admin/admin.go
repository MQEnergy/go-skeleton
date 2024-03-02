package admin

import "gorm.io/gen"

type Querier interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(id int) (gen.T, error)

	// SELECT * FROM @@table WHERE account = @account
	GetByAccount(account string) (*gen.T, error)
}
