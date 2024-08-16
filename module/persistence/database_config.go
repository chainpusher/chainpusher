package persistence

type DatabaseConfig struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func (d *DatabaseConfig) GetHost() string {
	return d.Host
}

func (d *DatabaseConfig) GetPort() int {
	return d.Port
}

func (d *DatabaseConfig) GetUser() string {
	return d.User
}

func (d *DatabaseConfig) GetPassword() string {
	return d.Password
}

func (d *DatabaseConfig) Dsn() string {
	return d.User + ":" + d.Password + "@tcp(" + d.Host + ":" + string(rune(d.Port)) + ")/" + d.Name + "?charset=utf8&parseTime=True&loc=Local"
}
