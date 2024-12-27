package action

import (
	"github.com/SuperH-0630/cat-shop-back/src/database/action/internal"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
)

const ClassListLimit = 100

func GetClassList(limit int, isDown bool, isShow bool, isEmptyClass bool) ([]model.Class, error) {
	if limit > ClassListLimit {
		limit = ClassListLimit
	} else if limit <= 0 {
		limit = ClassListLimit
	}

	var res = make([]model.Class, 0, 100)

	sql := internal.DB().Model(&model.Class{})
	if isDown {
		isShow = true
	} else {
		sql = sql.Where("down = false")
	}

	if isShow {
		sql = sql.Where("show = true")
	} else {
		isEmptyClass = false
	}

	if !isEmptyClass {
		sql = sql.Where("id != ?", modeltype.ClassEmptyID)
	}

	err := sql.Limit(limit).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
