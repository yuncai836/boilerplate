package config

import (
	"github.com/spf13/viper"
)

func InitViperLocalByEnv() error {
	var v = ReadEnvConfig[ViperConfig]()
	return InitViperLocalByPath(v.ConfigPath)
}

func InitViperLocalByPath(path string) error {
	viper.SetConfigFile(path)
	return viper.ReadInConfig()
}

func MustInitViperLocalByPath(path string) {
	if err := InitViperLocalByPath(path); err != nil {
		panic(err)
	}
}

func MustInitViperLocalByEnv() {
	if err := InitViperLocalByEnv(); err != nil {
		panic(err)
	}
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
