package action

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
)

func NewImage(tp modeltype.ImageType, file []byte) (*model.Image, error, error) {
	img, err := model.NewImage(tp, file)
	if err != nil {
		return nil, nil, err
	}

	err = internal.DB().Create(&img).Error
	if err != nil {
		return nil, err, nil
	}

	return img, nil, nil
}

func AdminNewImage(tp modeltype.ImageType, file []byte) (*model.Image, error, error) {
	return NewImage(tp, file)
}
