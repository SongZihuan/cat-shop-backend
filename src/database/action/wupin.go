package action

import (
	"errors"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

const HotWupinLimit = 50

func AdminGetWupinByID(wupinID uint) (*model.Wupin, error) {
	return AdminGetWupinByID(wupinID)
}

func adminGetWupinByID(wupinID uint, db *gorm.DB) (*model.Wupin, error) {
	var wupin = new(model.Wupin)

	if wupinID <= 0 {
		return nil, ErrNotFound
	}

	err := db.Model(&model.Wupin{}).Joins("Class").Where("id = ?", wupinID).First(wupin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return wupin, nil
}

func GetWupinByID(wupinID uint) (*model.Wupin, error) {
	var wupin = new(model.Wupin)

	if wupinID <= 0 {
		return nil, ErrNotFound
	}

	db := internal.DB()
	err := db.Model(&model.Wupin{}).Joins("Class").Where("id = ?", wupinID).Where("wupin_down = false").Where("class_down = false").First(wupin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
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

func AdminGetWupinList(page int, pagesize int) (res []model.Wupin, err error) {
	db := internal.DB()
	err = db.Model(&model.Wupin{}).Joins("Class").Where("wupin_down = false").Where("class_down = false").Limit(pagesize).Offset((page - 1) * pagesize).Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AdminGetWupinCount() (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	var res count
	db := internal.DB()
	err := db.Model(&model.Wupin{}).Select("COUNT(*) as count").Where("wupin_down = false").Where("class_down = false").First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return res.count, nil
}

func AdminUpdateWupin(wupinID uint, name string, pic string, cls *model.Class, tag string, hotPrice modeltype.PriceNull, realPrice modeltype.Price, info string, ren string, phone string, email string, wechat string, location string, hot bool, down bool) error {
	err := internal.DB().Transaction(func(tx *gorm.DB) error {
		err := systemCreateEmptyClass(tx)
		if err != nil {
			return err
		}

		wp, err := adminGetWupinByID(wupinID, tx)
		if err != nil {
			return err
		}

		needUpdate := wp.UpdateInfo(name, pic, cls, tag, hotPrice, realPrice, info, ren, phone, email, wechat, location, hot, down)
		err = tx.Save(cls).Error
		if err != nil {
			return err
		}

		if !needUpdate {
			return nil
		}

		err = tx.Model(&model.Bag{}).Where("wupin_id = ?", wp.ID).Update("wupin_down", wp.WupinDown).Error
		if err != nil {
			return err
		}

		err = tx.Model(&model.Bag{}).Where("wupin_id = ?", wp.ID).Update("wupin_down", wp.WupinDown).Error
		if err != nil {
			return err
		}

		err = tx.Model(&model.BuyRecord{}).Where("wupin_id = ?", wp.ID).Update("wupin_down", wp.WupinDown).Error
		if err != nil {
			return err
		}

		err = tx.Model(&model.BuyRecord{}).Where("wupin_id = ?", wp.ID).Update("wupin_down", wp.WupinDown).Error
		if err != nil {
			return err
		}

		return nil
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}
	return err
}

func AdminAddWupin(name string, pic string, cls *model.Class, tag string, hotPrice modeltype.PriceNull, realPrice modeltype.Price, info string, ren string, phone string, email string, wechat string, location string, hot bool, down bool) error {
	wupin := model.NewWupin(name, pic, cls, tag, hotPrice, realPrice, info, ren, phone, email, wechat, location, hot, down)
	return internal.DB().Create(wupin).Error
}
