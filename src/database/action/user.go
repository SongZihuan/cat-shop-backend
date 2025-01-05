package action

import (
	"errors"
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/database/action/internal"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
)

func RootAdminGetUserByID(userID uint) (*model.User, error) {
	var user = new(model.User)
	var err error

	if userID <= 0 {
		return nil, ErrNotFound
	}

	db := internal.DB()
	err = db.Model(model.User{}).Where("id = ?", userID).Order("create_at desc").First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func AdminGetUserByID(userID uint) (*model.User, error) {
	return GetUserByID(userID)
}

func MiddlewareGetUserByID(userID uint) (*model.User, error) {
	return GetUserByID(userID)
}

func GetUserByID(userID uint) (*model.User, error) {
	var user = new(model.User)
	var err error

	if userID <= 0 {
		return nil, ErrNotFound
	}

	db := internal.DB()
	err = db.Model(model.User{}).Where("id = ?", userID).Where("status != ?", modeltype.DeleteUserStatus).Order("create_at desc").First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func RootAdminGetUserByPhone(phone string) (*model.User, error) {
	var user = new(model.User)
	var err error

	db := internal.DB()
	err = db.Model(model.User{}).Where("phone = ?", phone).Order("create_at desc").First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func AdminGetUserByPhone(phone string) (*model.User, error) {
	return GetUserByPhone(phone)
}

func GetUserByPhone(phone string) (*model.User, error) {
	var user = new(model.User)
	var err error

	db := internal.DB()
	err = db.Model(model.User{}).Where("phone = ?", phone).Where("status != ?", modeltype.DeleteUserStatus).Order("create_at desc").First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func AdminCreateUser(phone string, password string) (*model.User, error) {
	user := model.NewUser(phone, password)
	err := internal.DB().Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *model.User, name string, wechat string, email string, location string) error {
	user.UpdateInfo(name, wechat, email, location)
	return internal.DB().Save(user).Error
}

func RootAdminUpdateUser(user *model.User, name string, wechat string, email string, location string, status modeltype.UserStatus, tp modeltype.UserType) (error, error, error) {
	user.UpdateInfo(name, wechat, email, location)
	ok := user.UpdateType(tp)
	if !ok {
		return fmt.Errorf("bad type"), nil, nil
	}

	ok = user.UpdateStatusWithRoot(status)
	if !ok {
		return nil, fmt.Errorf("bad status"), nil
	}

	return nil, nil, internal.DB().Save(user).Error
}

func AdminUpdateUser(user *model.User, name string, wechat string, email string, location string, status modeltype.UserStatus, tp modeltype.UserType) (error, error, error) {
	user.UpdateInfo(name, wechat, email, location)
	ok := user.UpdateType(tp)
	if !ok {
		return fmt.Errorf("bad type"), nil, nil
	}

	ok = user.UpdateStatus(status)
	if !ok {
		return nil, fmt.Errorf("bad status"), nil
	}

	return nil, nil, internal.DB().Save(user).Error
}

func UpdateUserPasswordWithCheck(user *model.User, oldPassword string, newPassword string) error {
	ok := user.UpdatePasswordWithCheck(oldPassword, newPassword)
	if !ok {
		return fmt.Errorf("password error")
	}
	return internal.DB().Save(user).Error
}

func AdminUpdateUserPhone(user *model.User, phone string) error {
	user.UpdatePhone(phone)
	return internal.DB().Save(user).Error
}

func UpdateUserPassword(user *model.User, newPassword string) error {
	user.UpdatePassword(newPassword)
	return internal.DB().Save(user).Error
}

func AdminUpdateUserPassword(user *model.User, newPassword string) error {
	return UpdateUserPassword(user, newPassword)
}

func UpdateUserAvatar(user *model.User, avatar string) error {
	user.UpdateAvatar(avatar)
	return internal.DB().Save(user).Error
}

func AdminUpdateUserAvatar(user *model.User, avatar string) error {
	return UpdateUserAvatar(user, avatar)
}

func RootAdminGetUserByPage(page int, pagesize int) (res []model.User, err error) {
	db := internal.DB()
	err = db.Model(&model.User{}).Joins("User").Limit(pagesize).Offset((page - 1) * pagesize).Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AdminGetUserByPage(page int, pagesize int) (res []model.User, err error) {
	db := internal.DB()
	err = db.Model(&model.User{}).Joins("User").Where("status != ?", modeltype.DeleteUserStatus).Limit(pagesize).Offset((page - 1) * pagesize).Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func RootAdminGetUserCount() (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	var res count
	var err error
	db := internal.DB()
	err = db.Model(&model.Msg{}).Select("COUNT(*) as count").First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return res.count, nil
}

func AdminGetUserCount() (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	var res count
	var err error
	db := internal.DB()
	err = db.Model(&model.Msg{}).Select("COUNT(*) as count").Where("status != ?", modeltype.DeleteUserStatus).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return res.count, nil
}
