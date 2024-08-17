package config

type Gateway struct {
	Datasource map[string]Datasource `yaml:"datasource"`
}
