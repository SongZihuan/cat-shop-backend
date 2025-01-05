package internal

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB = nil

func ConnectToMySql() error {
	if !config.IsReady() {
		panic("config must be ready")
	}

	if IsReady() {
		return nil
	}

	return connectToMySql(config.Config().Yaml.Mysql.UserName,
		config.Config().Yaml.Mysql.Password,
		config.Config().Yaml.Mysql.Address,
		config.Config().Yaml.Mysql.Port,
		config.Config().Yaml.Mysql.DBName)
}

func connectToMySql(username string, password string, address string, port int, dbname string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, port, dbname)
	tmp, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	_db = tmp
	return nil
}

func CloseMySql() {
	_db = nil
}

func IsReady() bool {
	return _db != nil
}

func DB() *gorm.DB {
	if !IsReady() {
		panic("db not connect")
	}
	return _db
}
