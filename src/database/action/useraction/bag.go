package useraction

import (
	"errors"
	"fmt"
	error2 "github.com/SongZihuan/cat-shop-backend/src/database/action/error"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"gorm.io/gorm"
)

func GetBagByIDAndUser(user *model.User, bagID uint) (*model.Bag, error) {
	var bag = new(model.Bag)
	var err error

	if bagID <= 0 {
		return nil, error2.ErrNotFound
	}

	db := internal.DB()
	err = db.Model(model.Bag{}).Joins("Wupin").Joins("Class").Where("id = ?", bagID).Where("user_id = ?", user.ID).Where("class_down = false").Where("wupin_down = false").First(bag).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, error2.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return bag, nil
}

func GetUserWupinBag(user *model.User, wupin *model.Wupin) (*model.Bag, error) {
	var bag = new(model.Bag)
	var err error

	db := internal.DB()
	err = db.Model(model.Bag{}).Joins("Wupin").Joins("Class").Where("wupin_id = ?", wupin.ID).Where("user_id = ?", user.ID).Where("class_down = false").Where("wupin_down = false").Order("time desc").First(bag).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, error2.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return bag, nil
}

func GetUserBagList(user *model.User, limit int, offset int) ([]model.Bag, error) {
	var res []model.Bag
	var err error

	db := internal.DB()
	err = db.Model(model.Bag{}).Joins("Wupin").Joins("Class").Where("user_id = ?", user.ID).Where("num > ?", 0).Limit(limit).Offset(offset).Find(&res).Error // 不再销售的也返回
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AddBag(user *model.User, bag *model.Bag, num int) (bool, error) {
	if bag.UserID != user.ID {
		return false, fmt.Errorf("bad user")
	}

	isNotEmpty := bag.Add(num)

	db := internal.DB()
	err := db.Save(bag).Error
	if err != nil {
		return false, err
	}

	return isNotEmpty, nil
}
