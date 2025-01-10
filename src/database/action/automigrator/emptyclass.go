package automigrator

import (
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

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
