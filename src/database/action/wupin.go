package action

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"gorm.io/gorm"
)

func GetWupinByID(wupinID uint) (*model.WuPin, error) {
	var wupin = new(model.WuPin)

	if wupinID <= 0 {
		return nil, ErrNotFound
	}

	db := database.DB()
	err := db.Model(model.WuPin{}).Where("id = ?", wupinID).First(wupin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return wupin, nil
}
