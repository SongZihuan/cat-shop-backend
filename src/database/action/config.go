package action

import (
	"github.com/SuperH-0630/cat-shop-back/src/database"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
)

func GetConfigLst() (res []model.Config, err error) {
	db := database.DB()
	err = db.Model(&model.Config{}).Where("key in ?", modeltype.ConfigKey).Limit(len(modeltype.ConfigKey)).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
