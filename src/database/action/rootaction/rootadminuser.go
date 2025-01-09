package rootaction

import (
	"errors"
	"fmt"
	error2 "github.com/SongZihuan/cat-shop-backend/src/database/action/error"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

func RootAdminGetUserByID(userID uint) (*model.User, error) {
	var user = new(model.User)
	var err error

	if userID <= 0 {
		return nil, error2.ErrNotFound
	}

	db := internal.DB()
	err = db.Model(model.User{}).Where("id = ?", userID).Order("create_at desc").First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, error2.ErrNotFound
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
		return nil, error2.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
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

func RootAdminGetUserByPage(page int, pagesize int) (res []model.User, err error) {
	db := internal.DB()
	err = db.Model(&model.User{}).Joins("User").Limit(pagesize).Offset((page - 1) * pagesize).Order("create_at desc").Find(&res).Error
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
