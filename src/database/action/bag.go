package action

import (
	"errors"
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/database/action/internal"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"gorm.io/gorm"
)

func GetBagByIDAndUser(user *model.User, bagID uint) (*model.Bag, error) {
	return GetBagByID(user.ID, bagID)
}

func GetBagByID(userID uint, bagID uint) (*model.Bag, error) {
	var bag = new(model.Bag)
	var err error

	if bagID <= 0 {
		return nil, ErrNotFound
	}

	db := internal.DB()
	err = db.Model(model.Bag{}).Joins("Wupin").Joins("Class").Where("id = ?", bagID).Where("user_id = ?", userID).Where("class_down = false").Where("wupin_down = false").First(bag).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return bag, nil
}

func GetBagByWupinIDWithUser(user *model.User, wupinID uint) (*model.Bag, error) {
	return GetBagByWupinIDWithUserID(user.ID, wupinID)
}

func GetBagByWupinIDWithUserID(userID uint, wupinID uint) (*model.Bag, error) {
	var bag = new(model.Bag)
	var err error

	if wupinID <= 0 {
		return nil, ErrNotFound
	}

	db := internal.DB()
	err = db.Model(model.Bag{}).Joins("Wupin").Joins("Class").Where("wu_pin_id = ?", wupinID).Where("user_id = ?", userID).Where("class_down = false").Where("wupin_down = false").Order("time desc").First(bag).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return bag, nil
}

func AdminGetBagByWupinIDWithUser(user *model.User, wupinID uint) (*model.Bag, error) {
	return AdminGetBagByWupinIDWithUserID(user.ID, wupinID)
}

func AdminGetBagByWupinIDWithUserID(userID uint, wupinID uint) (*model.Bag, error) {
	var bag = new(model.Bag)
	var err error

	if wupinID <= 0 {
		return nil, ErrNotFound
	}

	db := internal.DB()
	err = db.Model(model.Bag{}).Joins("Wupin").Joins("Class").Where("wu_pin_id = ?", wupinID).Where("user_id = ?", userID).Order("time desc").First(bag).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return bag, nil
}

func GetBagListByUser(user *model.User, limit int, offset int) ([]model.Bag, error) {
	return GetBagListByUserID(user.ID, limit, offset)
}

func GetBagListByUserID(userID uint, limit int, offset int) ([]model.Bag, error) {
	var res []model.Bag
	var err error

	db := internal.DB()
	err = db.Model(model.Bag{}).Joins("Wupin").Joins("Class").Joins("Class").Where("user_id = ?", userID).Where("class_down = false").Where("wupin_down = false").Where("num > ?", 0).Limit(limit).Offset(offset).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AdminGetBagListByUser(user *model.User, limit int, offset int) ([]model.Bag, error) {
	return AdminGetBagListByUserID(user.ID, limit, offset)
}

func AdminGetBagListByUserID(userID uint, limit int, offset int) ([]model.Bag, error) {
	var res []model.Bag
	var err error

	db := internal.DB()
	err = db.Model(model.Bag{}).Joins("Wupin").Joins("Class").Joins("Class").Where("user_id = ?", userID).Where("num > ?", 0).Limit(limit).Offset(offset).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AdminAddBag(user *model.User, bag *model.Bag, num int) (bool, error) {
	return AddBag(user, bag, num)
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
