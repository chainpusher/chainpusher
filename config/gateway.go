package config

import "github.com/chainpusher/chainpusher/module/datasource"

type Gateway struct {
	Datasource map[string]datasource.Datasource `yaml:"datasource"`
}
