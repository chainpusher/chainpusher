package datasource

type Datasource struct {
	Dsn    string `yaml:"dsn"`
	Driver string `yaml:"driver"`
}
