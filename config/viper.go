package config

import (
	"github.com/spf13/viper"
)

func InitViperLocalByEnv() error {
	var v = ReadEnvConfig[ViperConfig]()
	viper.SetConfigFile(v.ConfigPath)
	return viper.ReadInConfig()
}

func InitViperLocalByPath(path string) error {
	viper.SetConfigFile(path)
	return viper.ReadInConfig()
}

func ViperMustGetAll[T any]() *T {
	var t = new(T)
	err := viper.Unmarshal(t)
	if err != nil {
		panic(err)
	}
	return t
}

func ViperMustGetKey[T any](key string) *T {
	var t = new(T)
	err := viper.UnmarshalKey(key, t)
	if err != nil {
		return nil
	}
	return t
}
