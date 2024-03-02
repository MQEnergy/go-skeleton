package command

import (
	"go-skeleton/internal/app/entity"
	"go-skeleton/pkg/helper"
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
)

type GormGenCommand struct {
	generator    *gen.Generator
	db           *gorm.DB
	config       gen.Config
	ignoreFileds []string
	tables       []string
	tableMethods entity.MethodMaps
	dataTypeMap  map[string]func(detailType gorm.ColumnType) (dataType string)
}

type Options interface {
	apply(db *GormGenCommand)
}

type OptionsFunc struct {
	fun func(db *GormGenCommand)
}

// NewGenCommand ...
func NewGenCommand(db *gorm.DB, config gen.Config, options ...Options) (*GormGenCommand, error) {
	g := &GormGenCommand{
		db:           db,
		config:       config,
		ignoreFileds: make([]string, 0),
		tables:       make([]string, 0),
		tableMethods: make(entity.MethodMaps),
		dataTypeMap:  make(map[string]func(detailType gorm.ColumnType) (dataType string)),
	}
	generator := gen.NewGenerator(g.config)
	generator.UseDB(g.db)
	g.generator = generator

	for _, option := range options {
		option.apply(g)
	}
	return g, nil
}

// GenModels 生成模型和dao
func (c *GormGenCommand) GenModels() {
	if c.dataTypeMap != nil {
		c.generator.WithDataTypeMap(c.dataTypeMap)
	}
	jsonTagWithNS := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		if len(c.ignoreFileds) > 0 {
			for _, v := range c.ignoreFileds {
				if strings.Contains(v, columnName) {
					return "-"
				}
			}
		}
		return columnName
	})
	models := make([]interface{}, len(c.tables))
	if len(c.tables) > 0 {
		for _, table := range c.tables {
			model := c.generator.GenerateModelAs(table, strings.Title(helper.ToCamelCase(table)), jsonTagWithNS)
			models = append(models, model)
		}
	} else {
		models = c.generator.GenerateAllTable(jsonTagWithNS)
	}
	c.generator.ApplyBasic(models...)
	c.genModelsInterface()
	c.generator.Execute()
}

// genModelsInterface 生成模型接口
func (c *GormGenCommand) genModelsInterface() {
	if len(c.tableMethods) == 0 {
		return
	}
	for table, methods := range c.tableMethods {
		if len(methods) == 0 {
			continue
		}
		for _, method := range methods {
			c.generator.ApplyInterface(method, c.generator.GenerateModel(table))
		}
	}
}

func (f OptionsFunc) apply(db *GormGenCommand) {
	f.fun(db)
}

func WithIgnoreFields(ignoreFields ...string) *OptionsFunc {
	return &OptionsFunc{
		fun: func(db *GormGenCommand) {
			db.ignoreFileds = ignoreFields
		},
	}
}

func WithTableMethods(tableMethods entity.MethodMaps) *OptionsFunc {
	return &OptionsFunc{
		fun: func(db *GormGenCommand) {
			db.tableMethods = tableMethods
		},
	}
}

func WithDataTypeMap(dataTypeMap map[string]func(detailType gorm.ColumnType) (dataType string)) *OptionsFunc {
	return &OptionsFunc{
		fun: func(db *GormGenCommand) {
			db.dataTypeMap = dataTypeMap
		},
	}
}

func WithTables(tables ...string) *OptionsFunc {
	return &OptionsFunc{
		fun: func(db *GormGenCommand) {
			db.tables = tables
		},
	}
}
