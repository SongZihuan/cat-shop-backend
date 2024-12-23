package action

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"gorm.io/gorm"
)

const HotWupinLimit = 50

func GetWupinByIDWithShow(wupinID uint) (*model.WuPin, error) {
	var wupin = new(model.WuPin)

	if wupinID <= 0 {
		return nil, ErrNotFound
	}

	db := database.DB()
	err := db.Model(&model.WuPin{}).Joins("Class").Where("id = ?", wupinID).Where("is_show = true").First(wupin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return wupin, nil
}

func GetHotWupinListWithShow() (res []model.WuPin, err error) {
	db := database.DB()
	err = db.Model(&model.WuPin{}).Joins("Class").Limit(HotWupinLimit).Where("is_hot = true").Where("is_show = true").Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetSearchListWithShow(search string, selectClass []uint, page int, pagesize int) (res []model.WuPin, err error) {
	db := database.DB()
	sql := db.Model(&model.WuPin{}).Joins("Class").Where("is_show = true").Limit(pagesize).Offset((page - 1) * pagesize)

	if search != "" {
		sql = sql.Where("name LIKE ?", "%"+search+"%")
	}

	if len(selectClass) > 0 {
		sql = sql.Where("class_id IN (?)", selectClass)
	}

	err = sql.Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetSearchCountWithShow(search string, selectClass []uint) (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	db := database.DB()
	sql := db.Model(&model.WuPin{}).Select("count(*) as count").Where("is_show = true")

	if search != "" {
		sql = sql.Where("name LIKE ?", "%"+search+"%")
	}

	if len(selectClass) > 0 {
		sql = sql.Where("class_id IN (?)", selectClass)
	}

	var res count
	err := sql.Find(&res).Error
	if err != nil {
		return 0, err
	}

	return res.count, nil
}
