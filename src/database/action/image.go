package action

import (
	"github.com/SuperH-0630/cat-shop-back/src/database/action/internal"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
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
