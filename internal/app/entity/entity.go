package entity

import (
	"go-skeleton/internal/app/entity/admin"
	"go-skeleton/pkg/helper"
)

type MethodMaps map[string][]any

var methodMaps = MethodMaps{
	"yfo_admin": {
		func(Querier) {},
		func(admin.Querier) {},
	},
}

// Load ...
func Load(models []string) (tableMethods MethodMaps) {
	tableMethods = make(MethodMaps)
	if len(models) > 0 {
		for table, methods := range methodMaps {
			if helper.InAnySlice(models, table) {
				tableMethods[table] = methods
			} else {
				tableMethods[table] = []any{}
			}
		}
	} else {
		tableMethods = methodMaps
	}
	return
}
