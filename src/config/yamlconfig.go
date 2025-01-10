package config

import (
	"github.com/SongZihuan/cat-shop-backend/src/flagparser"
	"gopkg.in/yaml.v3"
	"os"
)

type YamlConfig struct {
	GlobalConfig `yaml:",inline"`
	Mysql        MySQLConfig    `yaml:"mysql"`
	File         FileConfig     `yaml:"file"`
	Http         HttpConfig     `yaml:"http"`
	Front        FrontConfig    `yaml:"front"`
	Jwt          JwtConfig      `yaml:"jwt"`
	Password     PasswordConfig `yaml:"password"`
}

func (y *YamlConfig) init() error {
	return nil
}

func (y *YamlConfig) setDefault() {
	y.GlobalConfig.setDefault()
	y.Mysql.setDefault()
	y.File.setDefault()
	y.Http.setDefault(&y.GlobalConfig)
	y.Front.setDefault()
	y.Jwt.setDefault()
	y.Password.setDefault()
}

func (y *YamlConfig) check(fl *FileLocationConfig, co *CorsOrigin) (err ConfigError) {
	err = y.GlobalConfig.check()
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

	err = y.Http.check(co)
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
