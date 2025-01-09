package useraction

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
)

const ClassListLimit = 100

func GetClassList(limit int) ([]model.Class, error) {
	if limit > ClassListLimit {
		limit = ClassListLimit
	} else if limit <= 0 {
		limit = ClassListLimit
	}

	var res = make([]model.Class, 0, 100)
	err := internal.DB().Model(&model.Class{}).Where("id != ?", modeltype.ClassEmptyID).Where("class_down = false").Where("show = true").Limit(limit).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
