package config

type MySQLConfig struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

func (m *MySQLConfig) setDefault() {
	if m.Port == 0 {
		m.Port = 3306
	}

	if m.Address == "" {
		m.Address = "127.0.0.1"
	}
}

func (m *MySQLConfig) check() ConfigError {
	if m.UserName == "" {
		return NewConfigError("mysql username must be given")
	} else if m.UserName == "root" {
		return NewConfigError("mysql username can not be root")
	}

	if m.DBName == "" {
		return NewConfigError("mysql dbname must be given")
	}

	return nil
}
