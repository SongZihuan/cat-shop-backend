package adminaction

import (
	"errors"
	error2 "github.com/SongZihuan/cat-shop-backend/src/database/action/error"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"gorm.io/gorm"
)

func AdminGetWupinBagWithUser(user *model.User, wupin *model.Wupin) (*model.Bag, error) {
	var bag = model.NewBag(user, wupin, 1)
	var err error

	db := internal.DB()
	err = db.Model(model.Bag{}).Joins("Wupin").Joins("Class").Where("wupin_id = ?", wupin.ID).Where("user_id = ?", user.ID).Order("time desc").FirstOrCreate(bag).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, error2.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return bag, nil
}

func AdminGetBagListByUser(user *model.User, page int, pagesize int) ([]model.Bag, error) {
	var res []model.Bag
	var err error

	db := internal.DB()
	err = db.Model(model.Bag{}).Joins("Wupin").Joins("Class").Where("user_id = ?", user.ID).Where("num > ?", 0).Limit(pagesize).Offset((page - 1) * pagesize).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AdminGetBagCountByUser(user *model.User) (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	var res count
	err := internal.DB().Model(model.Bag{}).Select("COUNT(*) as count").Where("user_id = ?", user.ID).Where("num > ?", 0).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return res.count, nil
}

func AdminAddBag(bag *model.Bag, num int) (bool, error) {
	isNotEmpty := bag.Add(num)

	db := internal.DB()
	err := db.Save(bag).Error
	if err != nil {
		return false, err
	}

	return isNotEmpty, nil
}
