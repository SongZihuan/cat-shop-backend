package useraction

import (
	"errors"
	error2 "github.com/SongZihuan/cat-shop-backend/src/database/action/error"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

const HotWupinLimit = 50

func GetWupinByID(wupinID uint) (*model.Wupin, error) {
	var wupin = new(model.Wupin)

	if wupinID <= 0 {
		return nil, error2.ErrNotFound
	}

	db := internal.DB()
	err := db.Model(&model.Wupin{}).Joins("Class").Where("id = ?", wupinID).Where("wupin_down = false").Where("class_down = false").First(wupin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, error2.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return wupin, nil
}

func GetHotWupinList(limit int) (res []model.Wupin, err error) {
	if limit > HotWupinLimit || limit <= 0 {
		limit = HotWupinLimit
	}

	db := internal.DB()
	err = db.Model(&model.Wupin{}).Joins("Class").Limit(HotWupinLimit).Where("hot = true").Where("wupin_down = false").Where("class_down = false").Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetSearchList(search string, selectClass uint, page int, pagesize int) (res []model.Wupin, err error) {
	db := internal.DB()
	sql := db.Model(&model.Wupin{}).Joins("Class").Where("wupin_down = false").Where("class_down = false").Limit(pagesize).Offset((page - 1) * pagesize)

	if search != "" {
		sql = sql.Where("name LIKE ?", "%"+search+"%")
	}

	if selectClass != 0 && selectClass != modeltype.ClassEmptyID {
		sql = sql.Where("class_id = ?", selectClass)
	}

	err = sql.Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetSearchCount(search string, selectClass uint) (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	db := internal.DB()
	sql := db.Model(&model.Wupin{}).Select("COUNT(*) as count").Where("wupin_down = false").Where("class_down = false")

	if search != "" {
		sql = sql.Where("name LIKE ?", "%"+search+"%")
	}

	if selectClass != 0 && selectClass != modeltype.ClassEmptyID {
		sql = sql.Where("class_id = ?", selectClass)
	}

	var res count
	err := sql.First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return res.count, nil
}
