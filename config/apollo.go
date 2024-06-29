package config

import (
	"github.com/philchia/agollo/v4"
	"gopkg.in/yaml.v3"
)

func MustInitApolloClient(logger agollo.Logger) {
	err := InitApolloClient(logger)
	if err != nil {
		panic(err)
	}
}

func InitApolloClient(logger agollo.Logger) error {
	var a = ReadEnvConfig[Apollo]()
	appConfig := agollo.Conf{
		AppID:           a.AppId,
		Cluster:         a.Cluster,
		NameSpaceNames:  a.NamespaceNames,
		MetaAddr:        a.MetaAddr,
		AccesskeySecret: a.AccesskeySecret,
	}
	return agollo.Start(&appConfig, agollo.WithLogger(logger), agollo.SkipLocalCache())
}

func ApolloMustGet[T any](namespace string) *T {
	var content = agollo.GetContent(agollo.WithNamespace(namespace))
	var t = new(T)
	err := yaml.Unmarshal([]byte(content), t)
	if err != nil {
		panic(err)
	}
	return t
}
