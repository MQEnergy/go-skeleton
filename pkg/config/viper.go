package config

import (
	"github.com/spf13/viper"
	"strings"
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
	v2.AddConfigPath(option.BasePath + "/configs")
	if strings.TrimSpace(option.FileName) == "" {
		v2.SetConfigName("config")
	} else {
		v2.SetConfigName(option.FileName)
	}
	v2.SetConfigType(option.Ctype)
	if err := v2.ReadInConfig(); err != nil {
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
