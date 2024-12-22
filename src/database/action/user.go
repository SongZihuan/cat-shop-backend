package action

import (
	"errors"
	"fmt"
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

func CreateUser(phone string, password string) (*model.User, error) {
	user := model.NewUser(phone, password)
	err := database.DB().Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *model.User, name string, wechat string, email string, location string) error {
	user.UpdateInfo(name, wechat, email, location)
	return database.DB().Save(user).Error
}

func UpdateUserPassword(user *model.User, oldPassword string, newPassword string) error {
	ok := user.UpdatePassword(oldPassword, newPassword)
	if !ok {
		return fmt.Errorf("password error")
	}
	return database.DB().Save(user).Error
}

func UpdateUserAvatar(user *model.User, avatar string) error {
	user.UpdateAvatar(avatar)
	return database.DB().Save(user).Error
}
