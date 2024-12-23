package action

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database"
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

	db := database.DB()
	err := db.Model(model.Bag{}).Joins("Wupin").Where("id = ?", bagID).Where("user_id = ?", userID).First(bag).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return bag, nil
}
