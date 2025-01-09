package adminaction

import (
	"errors"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"gorm.io/gorm"
)

func AdminGetMsgByPage(page int, pagesize int) (res []model.Msg, err error) {
	db := internal.DB()
	err = db.Model(&model.Msg{}).Joins("User").Limit(pagesize).Offset((page - 1) * pagesize).Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AdminGetMsgCount() (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	var res count
	db := internal.DB()
	err := db.Model(&model.Msg{}).Select("COUNT(*) as count").First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return res.count, nil
}

func AdminGetMsgByPageAndUser(user *model.User, page int, pagesize int) (res []model.Msg, err error) {
	db := internal.DB()
	err = db.Model(&model.Msg{}).Joins("User").Where("user_id = ?", user.ID).Limit(pagesize).Offset((page - 1) * pagesize).Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AdminGetMsgCountWithUser(user *model.User) (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	var res count
	db := internal.DB()
	err := db.Model(&model.Msg{}).Select("COUNT(*) as count").Where("user_id = ?", user.ID).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return res.count, nil
}
