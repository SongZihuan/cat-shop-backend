package adminaction

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
)

func AdminGetXieYi(xieyiType modeltype.XieYiType) (*model.Xieyi, error) {
	xieyi := new(model.Xieyi)

	if xieyiType == "" {
		xieyiType = modeltype.XieYiDefault
	}

	db := internal.DB()
	err := db.Model(&model.Xieyi{}).Where("type = ?", xieyiType).Limit(1).Order("created_at desc").FirstOrCreate(xieyi, model.Xieyi{Data: ""}).Error
	if err != nil {
		return nil, err
	}

	return xieyi, nil
}

func AdminUpdateXieYi(xieyiType modeltype.XieYiType, content string) error {
	if xieyiType == "" {
		xieyiType = modeltype.XieYiDefault
	}

	xieyi := model.NewXieyi(xieyiType, content)
	return internal.DB().Create(xieyi).Error
}
