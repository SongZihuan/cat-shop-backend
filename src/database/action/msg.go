package action

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action/internal"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"gorm.io/gorm"
)

func GetMsgByPage(page int, pagesize int) (res []model.Msg, err error) {
	db := internal.DB()
	err = db.Model(&model.Msg{}).Joins("User").Limit(pagesize).Offset((page - 1) * pagesize).Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetMsgCount() (int, error) {
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

func GetMsgByPageAndUser(user *model.User, page int, pagesize int) (res []model.Msg, err error) {
	return GetMsgByPageAndUserID(user.ID, page, pagesize)
}

func GetMsgCountWithUser(user *model.User) (int, error) {
	return GetMsgCountWithUserID(user.ID)
}

func GetMsgByPageAndUserID(userID uint, page int, pagesize int) (res []model.Msg, err error) {
	db := internal.DB()
	err = db.Model(&model.Msg{}).Joins("User").Where("user_id = ?", userID).Limit(pagesize).Offset((page - 1) * pagesize).Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetMsgCountWithUserID(userID uint) (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	var res count
	db := internal.DB()
	err := db.Model(&model.Msg{}).Select("COUNT(*) as count").Where("user_id = ?", userID).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return res.count, nil
}
