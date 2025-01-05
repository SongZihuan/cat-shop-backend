package action

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
)

func GetUserXieYi() (*model.Xieyi, error) {
	xieyi := new(model.Xieyi)

	db := internal.DB()
	err := db.Model(&model.Xieyi{}).Where("type = ?", modeltype.XieyiUser).Limit(1).Order("created_at desc").FirstOrCreate(xieyi, model.Xieyi{Data: ""}).Error
	if err != nil {
		return nil, err
	}

	return xieyi, nil
}

func AdminUpdateUserXieyi(xieyiType modeltype.XieYiType, content string) error {
	xieyi := model.NewXieyi(xieyiType, content)
	db := internal.DB()
	return db.Create(xieyi).Error
}
