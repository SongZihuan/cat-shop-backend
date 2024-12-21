package action

import (
	"github.com/SuperH-0630/cat-shop-back/src/database"
	"github.com/SuperH-0630/cat-shop-back/src/model"
)

func GetXieYi() (*model.Xieyi, error) {
	xieyi := new(model.Xieyi)

	db := database.DB()
	err := db.Model(&model.Xieyi{}).Limit(1).Order("created_at desc").FirstOrCreate(xieyi, model.Xieyi{Data: ""}).Error
	if err != nil {
		return nil, err
	}

	return xieyi, nil
}
