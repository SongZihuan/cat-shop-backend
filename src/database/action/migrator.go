package action

import (
	"errors"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

func SystemAutoMigrate() error {
	if !internal.IsReady() {
		return errors.New("db is not ready")
	}

	db := internal.DB()
	return db.AutoMigrate(model.AutoCreateModelList...)
}

func SystemCreateEmptyClass() error {
	if !internal.IsReady() {
		return errors.New("db is not ready")
	}
	return systemCreateEmptyClass(internal.DB())
}

func systemCreateEmptyClass(db *gorm.DB) error {
	cls := model.NewEmptyClass()
	err := db.Model(&model.ClassM{}).Where("id = ?", modeltype.ClassEmptyID).Limit(1).FirstOrCreate(cls).Error
	if err != nil {
		return err
	}

	if cls.ResetIsEmpty() {
		err := db.Save(cls).Error
		if err != nil {
			return err
		}
	}

	return nil
}
