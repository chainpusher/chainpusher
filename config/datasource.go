package config

type Datasource struct {
	Dsn    string `yaml:"dsn"`
	Driver string `yaml:"driver"`
}
