package automigrator

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
)

func SystemAutoMigrate() error {
	if !internal.IsReady() {
		panic("db is not ready")
	}

	db := internal.DB()
	return db.AutoMigrate(model.AutoCreateModelList...)
}
