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

	if bagID <= 0 {
		return nil, ErrNotFound
	}

	db := internal.DB()
	err := db.Model(model.Bag{}).Joins("Wupin").Where("id = ?", bagID).Where("user_id = ?", userID).First(bag).Error
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

	if wupinID <= 0 {
		return nil, ErrNotFound
	}

	db := internal.DB()
	err := db.Model(model.Bag{}).Joins("Wupin").Where("wu_pin_id = ?", wupinID).Where("user_id = ?", userID).Order("time desc").First(bag).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return bag, nil
}

func GetBagListByUser(user *model.User, limit int, offset int, isAdmin bool) ([]model.Bag, error) {
	return GetBagListByUserID(user.ID, limit, offset, isAdmin)
}

func GetBagListByUserID(userID uint, limit int, offset int, isAdmin bool) ([]model.Bag, error) {
	var res []model.Bag
	var err error

	db := internal.DB()
	if isAdmin {
		err = db.Model(model.Bag{}).Joins("Wupin").Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&res).Error
	} else {
		err = db.Model(model.Bag{}).Joins("Wupin").Where("user_id = ?", userID).Where("wu_pin_show = true").Limit(limit).Offset(offset).Find(&res).Error
	}
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
