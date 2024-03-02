package command

import (
	"strings"

	"github.com/urfave/cli/v2"
	"go-skeleton/internal/app/entity"
	"go-skeleton/internal/bootstrap"
	"go-skeleton/internal/vars"
	"go-skeleton/pkg/command"
	"go-skeleton/pkg/config"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type GenModel struct{}

// Command ...
func (g *GenModel) Command() *cli.Command {
	var models string

	return &cli.Command{
		Name:  "genModel",
		Usage: "基于gorm的gen的代码生成器，生成数据表model，并生成model对应的方法。",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "env",
				Aliases:     []string{"e"},
				Value:       "dev",
				Usage:       "请选择配置文件 [dev | test | prod]",
				Destination: &config.ConfEnv,
			},
			&cli.StringFlag{
				Name:        "model",
				Aliases:     []string{"m"},
				Value:       "",
				Usage:       "请输入数据表名称 输入按照逗号分割 如：user,admin..., 如果不填生成全部模型",
				Destination: &models,
			},
		},
		Before: func(ctx *cli.Context) error {
			if err := bootstrap.InitConfig(); err != nil {
				return err
			}
			bootstrap.BootService(bootstrap.MysqlService)
			return nil
		},
		Action: func(ctx *cli.Context) error {
			return handleGenModel(models)
		},
	}
}

var _ command.Interface = (*GenModel)(nil)

// handleGenModel ...
func handleGenModel(models string) error {
	modelNames := make([]string, 0)
	if models != "" {
		modelNames = strings.Split(models, ",")
	}

	newGenCommand, err := command.NewGenCommand(vars.DB, gen.Config{
		OutPath:           "./internal/app/dao",
		OutFile:           "",
		ModelPkgPath:      "./internal/app/model",
		WithUnitTest:      false,
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
	},
		command.WithTables(modelNames...),
		command.WithTableMethods(entity.Load(modelNames)),
		command.WithIgnoreFields(),
		command.WithDataTypeMap(map[string]func(columnType gorm.ColumnType) (dataType string){
			"varchar":   func(columnType gorm.ColumnType) (dataType string) { return "string" },
			"int":       func(columnType gorm.ColumnType) (dataType string) { return "int" },
			"bigint":    func(columnType gorm.ColumnType) (dataType string) { return "int64" },
			"tinyint":   func(columnType gorm.ColumnType) (dataType string) { return "int8" },
			"smallint":  func(columnType gorm.ColumnType) (dataType string) { return "int16" },
			"mediumint": func(columnType gorm.ColumnType) (dataType string) { return "int32" },
			"decimal":   func(columnType gorm.ColumnType) (dataType string) { return "float64" },
			"timestamp": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		}),
	)
	if err != nil {
		return err
	}
	newGenCommand.GenModels()
	return nil
}
