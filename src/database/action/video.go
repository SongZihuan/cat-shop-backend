package action

import (
	"github.com/SuperH-0630/cat-shop-back/src/database/action/internal"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
)

func NewVideo(tp modeltype.VideoType, file []byte) (*model.Video, error, error) {
	vid, err := model.NewVideo(tp, file)
	if err != nil {
		return nil, nil, err
	}

	err = internal.DB().Create(&vid).Error
	if err != nil {
		return nil, err, nil
	}

	return vid, nil, nil
}
