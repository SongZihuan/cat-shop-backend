package config

import (
	"github.com/SuperH-0630/cat-shop-back/src/flagparser"
	"gopkg.in/yaml.v3"
	"os"
)

type YamlConfig struct {
	Mysql MySQLConfig `yaml:"mysql"`
	File  FileConfig  `yaml:"file"`
	Http  HttpConfig  `yaml:"http"`
	Jwt   JwtConfig   `yaml:"jwt"`
}

func (y *YamlConfig) setDefault() {
	y.Mysql.setDefault()
	y.File.setDefault()
	y.Http.setDefault()
	y.Jwt.setDefault()
}

func (y *YamlConfig) check() (err ConfigError) {
	err = y.Mysql.check()
	if err != nil && err.IsError() {
		return err
	}

	err = y.File.check()
	if err != nil && err.IsError() {
		return err
	}

	err = y.Http.check()
	if err != nil && err.IsError() {
		return err
	}

	err = y.Jwt.check()
	if err != nil && err.IsError() {
		return err
	}

	return nil
}

func (y *YamlConfig) parser() ParserError {
	file, err := os.ReadFile(flagparser.ConfigFile())
	if err != nil {
		return NewParserError(err, err.Error())
	}

	err = yaml.Unmarshal(file, y)
	if err != nil {
		return NewParserError(err, err.Error())
	}

	return nil
}
