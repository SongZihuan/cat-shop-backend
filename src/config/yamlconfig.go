package config

import (
	"github.com/SongZihuan/cat-shop-backend/src/flagparser"
	"gopkg.in/yaml.v3"
	"os"
)

type YamlConfig struct {
	Global   GlobalConfig   `yaml:"global"`
	Mysql    MySQLConfig    `yaml:"mysql"`
	File     FileConfig     `yaml:"file"`
	Http     HttpConfig     `yaml:"http"`
	Front    FrontConfig    `yaml:"front"`
	Jwt      JwtConfig      `yaml:"jwt"`
	Password PasswordConfig `yaml:"password"`
}

func (y *YamlConfig) setDefault() {
	y.Global.setDefault()
	y.Mysql.setDefault()
	y.File.setDefault()
	y.Http.setDefault()
	y.Front.setDefault()
	y.Jwt.setDefault()
	y.Password.setDefault()
}

func (y *YamlConfig) check(fl *FileLocationConfig) (err ConfigError) {
	err = y.Global.check()
	if err != nil && err.IsError() {
		return err
	}

	err = y.Mysql.check()
	if err != nil && err.IsError() {
		return err
	}

	err = y.File.check(fl)
	if err != nil && err.IsError() {
		return err
	}

	err = y.Http.check()
	if err != nil && err.IsError() {
		return err
	}

	err = y.Front.check()
	if err != nil && err.IsError() {
		return err
	}

	err = y.Jwt.check()
	if err != nil && err.IsError() {
		return err
	}

	err = y.Password.check()
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
