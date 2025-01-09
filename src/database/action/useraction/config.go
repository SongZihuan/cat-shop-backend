package useraction

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
)

const UserMaxLimit = 30

func GetConfigLst(limit int) (res []model.Config, err error) {
	_min := min(len(modeltype.ConfigKey), UserMaxLimit)

	if limit <= 0 || limit > _min {
		limit = _min
	}

	db := internal.DB()
	err = db.Model(&model.Config{}).Where("key in ?", modeltype.ConfigKey).Limit(_min).Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
