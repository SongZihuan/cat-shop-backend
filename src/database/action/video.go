package action

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
)

func AdminNewVideo(tp modeltype.VideoType, file []byte) (*model.Video, error, error) {
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
