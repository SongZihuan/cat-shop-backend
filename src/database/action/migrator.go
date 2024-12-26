package action

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
)

func AutoMigrate() error {
	if !database.IsReady() {
		return errors.New("db is not ready")
	}

	db := database.DB()
	return db.AutoMigrate(model.AutoCreateModelList...)
}

func CreateEmptyClass() error {
	if !database.IsReady() {
		return errors.New("db is not ready")
	}

	db := database.DB()
	cls := model.NewEmptyClass()
	err := db.Model(&model.ClassM{}).Where("id = ?", modeltype.ClassEmptyID).Limit(1).FirstOrCreate(cls).Error
	if err != nil {
		return err
	}

	if !cls.IsEmptyWithCheck() {
		cls.ResetEmpty()
		err := db.Save(cls).Error
		if err != nil {
			return err
		}
	}

	return nil
}
