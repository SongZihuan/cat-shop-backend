package automigrator

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"gorm.io/gorm"
)

func SystemMustCreateData() error {
	if !internal.IsReady() {
		panic("db is not ready")
	}

	return internal.DB().Transaction(func(tx *gorm.DB) error {
		err := systemCreateEmptyClass(tx)
		if err != nil {
			return err
		}

		return nil
	})
}

func SystemCreateData() error {
	if !config.IsReady() {
		panic("config is not ready")
	}

	if config.Config().Yaml.Mysql.FakeData.IsDisable() {
		return nil
	}

	if !internal.IsReady() {
		panic("db is not ready")
	}

	return internal.DB().Transaction(func(tx *gorm.DB) error {
		return nil
	})
}
