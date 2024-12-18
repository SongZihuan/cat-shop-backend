package action

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database"
	"github.com/SuperH-0630/cat-shop-back/src/model"
)

func AutoMigrate() error {
	if !database.IsReady() {
		return errors.New("db is not ready")
	}

	db := database.DB()
	return db.AutoMigrate(
		&model.Bag{},
		&model.BuyRecord{},
		&model.Class{},
		&model.Config{},
		&model.Image{},
		&model.Msg{},
		&model.User{},
		&model.WuPin{},
		model.Xieyi{},
	)
}
