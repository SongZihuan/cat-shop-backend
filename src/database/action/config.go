package action

import (
	"errors"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

func GetConfigLst() (res []model.Config, err error) {
	db := internal.DB()
	err = db.Model(&model.Config{}).Where("key in ?", modeltype.ConfigKey).Limit(len(modeltype.ConfigKey)).Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func AdminGetConfigLst() (res []model.Config, err error) {
	return GetConfigLst()
}

func AdminDeleteConfig(key modeltype.ConfigKeyType) error {
	return AdminUpdateConfig(key, "")
}

func AdminUpdateConfigString(key modeltype.ConfigKeyType, value modeltype.ConfigValueType) error {
	return AdminUpdateConfig(key, value)
}

func AdminUpdateConfigPic(key modeltype.ConfigKeyType, img *model.Image) error {
	return AdminUpdateConfig(key, modeltype.ConfigValueType(img.GetUrl()))
}

func AdminUpdateConfig(key modeltype.ConfigKeyType, value modeltype.ConfigValueType) error {
	return internal.DB().Transaction(func(tx *gorm.DB) error {
		var cfg = new(model.Config)
		err := tx.Model(&model.Config{}).Where("key = ?", key).Order("create_at desc").First(cfg).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cfg.Key = key
			cfg.Value = value
			return tx.Create(cfg).Error
		} else if err != nil {
			return err
		} else {
			cfg.Value = value
			return tx.Save(cfg).Error
		}
	})

}
