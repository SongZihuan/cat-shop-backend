package action

import (
	"github.com/SuperH-0630/cat-shop-back/src/database/action/internal"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
)

const ClassListLimit = 100

func GetClassList(limit int, showEmpty bool, showOne bool) ([]model.Class, error) {
	if limit > ClassListLimit {
		limit = ClassListLimit
	} else if limit <= 0 {
		limit = ClassListLimit
	}

	var res = make([]model.Class, 0, 100)
	var err error
	if showEmpty {
		if showOne {
			err = internal.DB().Model(&model.Class{}).Limit(limit).Find(&res).Error
		} else {
			err = internal.DB().Model(&model.Class{}).Where("id != ?", modeltype.ClassEmptyID).Limit(limit).Find(&res).Error
		}
	} else {
		err = internal.DB().Model(&model.Class{}).Where("show = true").Where("id != ?", modeltype.ClassEmptyID).Limit(limit).Find(&res).Error
	}
	if err != nil {
		return nil, err
	}

	return res, nil
}
