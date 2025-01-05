package action

import (
	"errors"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

const ClassListLimit = 100

func AdminGetClassLst(limit int) ([]model.Class, error) {
	if limit > ClassListLimit {
		limit = ClassListLimit
	} else if limit <= 0 {
		limit = ClassListLimit
	}

	var res = make([]model.Class, 0, 100)

	sql := internal.DB().Model(&model.Class{})
	err := sql.Limit(limit).Where("id != ?", modeltype.ClassEmptyID).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AdminGetClassListByPage(page int, pagesize int) (res []model.Class, err error) {
	db := internal.DB()
	err = db.Model(&model.Class{}).Where("id != ?", modeltype.ClassEmptyID).Limit(pagesize).Offset((page - 1) * pagesize).Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AdminGetClassCount() (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	var res count
	db := internal.DB()
	err := db.Model(&model.Class{}).Select("COUNT(*) as count").First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return res.count, nil
}

func GetClassList(limit int) ([]model.Class, error) {
	if limit > ClassListLimit {
		limit = ClassListLimit
	} else if limit <= 0 {
		limit = ClassListLimit
	}

	var res = make([]model.Class, 0, 100)
	err := internal.DB().Model(&model.Class{}).Where("id != ?", modeltype.ClassEmptyID).Where("down = false").Where("show = true").Limit(limit).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
