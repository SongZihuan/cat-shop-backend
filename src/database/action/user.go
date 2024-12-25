package action

import (
	"errors"
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/database"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
)

func GetUserByID(userID uint, isRoot bool) (*model.User, error) {
	var user = new(model.User)
	var err error

	if userID <= 0 {
		return nil, ErrNotFound
	}

	db := database.DB()
	if isRoot {
		err = db.Model(model.User{}).Where("id = ?", userID).Order("create_at desc").First(user).Error
	} else {
		err = db.Model(model.User{}).Where("id = ?", userID).Where("status != ?", modeltype.DeleteUserStatus).Order("create_at desc").First(user).Error
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByPhone(phone string, isRoot bool) (*model.User, error) {
	var user = new(model.User)
	var err error

	db := database.DB()
	if isRoot {
		err = db.Model(model.User{}).Where("phone = ?", phone).Order("create_at desc").First(user).Error
	} else {
		err = db.Model(model.User{}).Where("phone = ?", phone).Where("status != ?", modeltype.DeleteUserStatus).Order("create_at desc").First(user).Error
	}
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

func AdminCreateUser(user *model.User, name string, wechat string, email string, location string, status modeltype.UserStatus, tp modeltype.UserType, isRoot bool) (error, error, error) {
	user.UpdateInfo(name, wechat, email, location)
	ok := user.UpdateType(tp)
	if !ok {
		return fmt.Errorf("bad type"), nil, nil
	}

	ok = user.UpdateStatus(status, isRoot)
	if !ok {
		return nil, fmt.Errorf("bad status"), nil
	}

	return nil, nil, database.DB().Save(user).Error
}

func UpdateUserPasswordWithCheck(user *model.User, oldPassword string, newPassword string) error {
	ok := user.UpdatePasswordWithCheck(oldPassword, newPassword)
	if !ok {
		return fmt.Errorf("password error")
	}
	return database.DB().Save(user).Error
}

func AdminUpdateUserPhone(user *model.User, phone string) error {
	user.UpdatePhone(phone)
	return database.DB().Save(user).Error
}

func UpdateUserPassword(user *model.User, newPassword string) error {
	user.UpdatePassword(newPassword)
	return database.DB().Save(user).Error
}

func UpdateUserAvatar(user *model.User, avatar string) error {
	user.UpdateAvatar(avatar)
	return database.DB().Save(user).Error
}

func GetUserByPage(page int, pagesize int, isRoot bool) (res []model.User, err error) {
	db := database.DB()
	if isRoot {
		err = db.Model(&model.User{}).Joins("User").Limit(pagesize).Offset((page - 1) * pagesize).Order("create_at desc").Find(&res).Error
	} else {
		err = db.Model(&model.User{}).Joins("User").Where("status != ?", modeltype.DeleteUserStatus).Limit(pagesize).Offset((page - 1) * pagesize).Order("create_at desc").Find(&res).Error
	}
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetUserCount(isRoot bool) (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	var res count
	var err error
	db := database.DB()
	if isRoot {
		err = db.Model(&model.Msg{}).Select("COUNT(*) as count").First(&res).Error
	} else {
		err = db.Model(&model.Msg{}).Select("COUNT(*) as count").Where("status != ?", modeltype.DeleteUserStatus).First(&res).Error
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return res.count, nil
}
