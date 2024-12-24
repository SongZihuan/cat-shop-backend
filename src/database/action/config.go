package action

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
)

func GetConfigLst() (res []model.Config, err error) {
	db := database.DB()
	err = db.Model(&model.Config{}).Where("key in ?", modeltype.ConfigKey).Limit(len(modeltype.ConfigKey)).Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteConfig(key modeltype.ConfigKeyType) error {
	return UpdateConfig(key, "")
}

func UpdateConfigString(key modeltype.ConfigKeyType, value modeltype.ConfigValueType) error {
	return UpdateConfig(key, value)
}

func UpdateConfigPic(key modeltype.ConfigKeyType, img *model.Image) error {
	return UpdateConfig(key, modeltype.ConfigValueType(img.GetUrl()))
}

func UpdateConfig(key modeltype.ConfigKeyType, value modeltype.ConfigValueType) error {
	return database.DB().Transaction(func(tx *gorm.DB) error {
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
