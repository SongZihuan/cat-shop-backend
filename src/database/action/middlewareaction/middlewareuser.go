package middlewareaction

import (
	"errors"
	error2 "github.com/SongZihuan/cat-shop-backend/src/database/action/error"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

func MiddlewareGetUserByID(userID uint) (*model.User, error) {
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
