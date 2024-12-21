package action

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"gorm.io/gorm"
)

func GetUserByID(userID uint) (*model.User, error) {
	var user = new(model.User)

	if userID <= 0 {
		return nil, ErrNotFound
	}

	db := database.DB()
	err := db.Model(model.User{}).Where("id = ?", userID).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByPhone(phone string) (*model.User, error) {
	var user = new(model.User)

	db := database.DB()
	err := db.Model(model.User{}).Where("phone = ?", phone).Order("create_at desc").First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(user *model.User) error {
	return database.DB().Create(user).Error
}
