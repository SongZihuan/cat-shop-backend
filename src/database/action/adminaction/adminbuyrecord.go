package adminaction

import (
	"errors"
	error2 "github.com/SongZihuan/cat-shop-backend/src/database/action/error"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"gorm.io/gorm"
)

func AdminGetBuyRecordByID(user *model.User, recordID uint) (*model.BuyRecord, error) {
	var record = new(model.BuyRecord)

	if recordID <= 0 {
		return nil, error2.ErrNotFound
	}

	db := internal.DB()
	err := db.Model(model.BuyRecord{}).Joins("Wupin").Joins("Class").Where("id = ?", recordID).Where("user_id = ?", user.ID).First(record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, error2.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	if record.ClassID > 0 && record.Wupin.ClassID > 0 && record.ClassID == record.Wupin.ClassID {
		record.Class = record.Wupin.Class
	}
	return record, nil
}

func AdminGetBuyRecordListByPageByUser(user *model.User, page int, pagesize int) ([]model.BuyRecord, error) {
	var res []model.BuyRecord

	db := internal.DB()
	err := db.Model(model.BuyRecord{}).Joins("Wupin").Joins("Class").Where("user_id = ?", user.ID).Limit(pagesize).Offset((page - 1) * pagesize).Find(res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AdminGetBuyRecordLCountByPageByUser(user *model.User) (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	var res count

	db := internal.DB()
	err := db.Model(model.BuyRecord{}).Select("COUNT(*) as count").Where("user_id = ?", user.ID).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return res.count, nil
}

func AdminGetBuyRecordListByPage(page int, pagesize int) ([]model.BuyRecord, error) {
	var res []model.BuyRecord

	err := internal.DB().Model(model.BuyRecord{}).Joins("Wupin").Joins("Class").Limit(pagesize).Offset((page - 1) * pagesize).Find(res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AdminGetBuyRecordCountByPage() (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	var res count
	err := internal.DB().Model(model.BuyRecord{}).Select("COUNT(*) as count").First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return res.count, nil
}
