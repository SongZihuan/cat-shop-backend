package useraction

import (
	"errors"
	"fmt"
	error2 "github.com/SongZihuan/cat-shop-backend/src/database/action/error"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

func GetUserByID(userID uint) (*model.User, error) {
	var user = new(model.User)
	var err error

	if userID <= 0 {
		return nil, error2.ErrNotFound
	}

	db := internal.DB()
	err = db.Model(model.User{}).Where("id = ?", userID).Where("status != ?", modeltype.DeleteUserStatus).Order("create_at desc").First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, error2.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByPhone(phone string) (*model.User, error) {
	var user = new(model.User)
	var err error

	db := internal.DB()
	err = db.Model(model.User{}).Where("phone = ?", phone).Where("status != ?", modeltype.DeleteUserStatus).Order("create_at desc").First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, error2.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(user *model.User, name string, wechat string, email string, location string) error {
	user.UpdateInfo(name, wechat, email, location)
	return internal.DB().Save(user).Error
}

func UpdateUserPasswordWithCheck(user *model.User, oldPassword string, newPassword string) error {
	ok := user.UpdatePasswordWithCheck(oldPassword, newPassword)
	if !ok {
		return fmt.Errorf("password error")
	}
	return internal.DB().Save(user).Error
}

func UpdateUserPassword(user *model.User, newPassword string) error {
	user.UpdatePassword(newPassword)
	return internal.DB().Save(user).Error
}

func UpdateUserAvatar(user *model.User, avatar string) error {
	user.UpdateAvatar(avatar)
	return internal.DB().Save(user).Error
}
