package config

import (
	"bytes"
	"github.com/MQEnergy/go-skeleton/configs"
	"strings"

	"github.com/spf13/viper"
)

type Viper struct {
	viper  *viper.Viper
	option Options
}

var _ CommonInterface = (*Viper)(nil)

func NewViper() *Viper {
	return &Viper{}
}

// Apply 创建实例
func (v *Viper) Apply(option Options) error {
	v.option = option

	v2 := viper.New()
	v2.SetConfigType(v.option.Ctype)
	v2.AddConfigPath(v.option.BasePath + "/configs")
	if strings.TrimSpace(v.option.FileName) == "" {
		v2.SetConfigName("config.dev")
	} else {
		v2.SetConfigName(v.option.FileName)
	}
	yamlConf, err := configs.TemplateFs.ReadFile(v.option.FileName + ".yaml")
	if err != nil {
		return err
	}
	if err := v2.ReadConfig(bytes.NewBuffer(yamlConf)); err != nil {
		return err
	}
	v.viper = v2
	return nil
}

func (v *Viper) Get(key string) any {
	return v.viper.Get(key)
}

func (v *Viper) Set(key string, value any) bool {
	v.viper.Set(key, value)
	return true
}

func (v *Viper) Has(key string) bool {
	return v.viper.IsSet(key)
}
