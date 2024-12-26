package action

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action/internal"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
)

func AutoMigrate() error {
	if !internal.IsReady() {
		return errors.New("db is not ready")
	}

	db := internal.DB()
	return db.AutoMigrate(model.AutoCreateModelList...)
}

func CreateEmptyClass() error {
	if !internal.IsReady() {
		return errors.New("db is not ready")
	}

	db := internal.DB()
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
