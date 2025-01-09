package internal

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB = nil

func ConnectToMySQL() error {
	if !config.IsReady() {
		panic("config must be ready")
	}

	if IsReady() {
		return nil
	}

	return connectToMySQL(config.Config().Yaml.Mysql.UserName,
		config.Config().Yaml.Mysql.Password,
		config.Config().Yaml.Mysql.Address,
		config.Config().Yaml.Mysql.Port,
		config.Config().Yaml.Mysql.DBName)
}

func connectToMySQL(username string, password string, address string, port int, dbname string) error {
	if _db != nil {
		panic("db is connect")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, port, dbname)
	tmp, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	_db = tmp
	return nil
}

func CloseMySQL() {
	if _db == nil {
		return
	}

	defer func() {
		_db = nil
	}()

	if !config.IsReady() {
		panic("config is not ready")
	}

	if config.Config().Yaml.Mysql.ActiveShutdown.IsEnable() {
		// https://github.com/go-gorm/gorm/issues/3145
		sqlDB, err := _db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.Close()
	}
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
