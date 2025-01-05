package action

import (
	"errors"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

const HotWupinLimit = 50

func GetWupinByIDWithShow(wupinID uint) (*model.WuPin, error) {
	var wupin = new(model.WuPin)

	if wupinID <= 0 {
		return nil, ErrNotFound
	}

	db := internal.DB()
	err := db.Model(&model.WuPin{}).Joins("Class").Where("id = ?", wupinID).Where("down = false").Where("class_down = false").First(wupin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return wupin, nil
}

func GetHotWupinListWithShow() (res []model.WuPin, err error) {
	db := internal.DB()
	err = db.Model(&model.WuPin{}).Joins("Class").Limit(HotWupinLimit).Where("hot = true").Where("down = false").Where("class_down = false").Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetSearchListWithShow(search string, selectClass uint, page int, pagesize int) (res []model.WuPin, err error) {
	db := internal.DB()
	sql := db.Model(&model.WuPin{}).Joins("Class").Where("down = false").Where("class_down = false").Limit(pagesize).Offset((page - 1) * pagesize)

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

func GetSearchCountWithShow(search string, selectClass uint) (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	db := internal.DB()
	sql := db.Model(&model.WuPin{}).Select("COUNT(*) as count").Where("down = false").Where("class_down = false")

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
