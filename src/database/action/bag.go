package action

import (
	"errors"
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
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
	err = db.Model(model.Bag{}).Joins("Wupin").Joins("Class").Where("wupin_id = ?", wupinID).Where("user_id = ?", userID).Where("class_down = false").Where("wupin_down = false").Order("time desc").First(bag).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return bag, nil
}

func AdminGetBagByWupinIDWithUser(user *model.User, wupin *model.Wupin) (*model.Bag, error) {
	var bag = model.NewBag(user, wupin, 1)
	var err error

	db := internal.DB()
	err = db.Model(model.Bag{}).Joins("Wupin").Joins("Class").Where("wupin_id = ?", wupin.ID).Where("user_id = ?", user.ID).Order("time desc").FirstOrCreate(bag).Error
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
	err = db.Model(model.Bag{}).Joins("Wupin").Joins("Class").Where("user_id = ?", userID).Where("class_down = false").Where("wupin_down = false").Where("num > ?", 0).Limit(limit).Offset(offset).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AdminGetBagListByUser(user *model.User, page int, pagesize int) ([]model.Bag, error) {
	return AdminGetBagListByUserID(user.ID, page, pagesize)
}

func AdminGetBagListByUserID(userID uint, page int, pagesize int) ([]model.Bag, error) {
	var res []model.Bag
	var err error

	db := internal.DB()
	err = db.Model(model.Bag{}).Joins("Wupin").Joins("Class").Where("user_id = ?", userID).Where("num > ?", 0).Limit(pagesize).Offset((page - 1) * pagesize).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AdminGetBagCountByUser(user *model.User) (int, error) {
	return AdminGetBagCountByUserID(user.ID)
}

func AdminGetBagCountByUserID(userID uint) (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	var res count
	err := internal.DB().Model(model.Bag{}).Select("COUNT(*) as count").Where("user_id = ?", userID).Where("num > ?", 0).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return res.count, nil
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
