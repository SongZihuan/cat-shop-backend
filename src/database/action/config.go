package action

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

const AdminMaxLimit = 100
const UserMaxLimit = 30

func GetConfigLst() (res []model.Config, err error) {
	_min := min(len(modeltype.ConfigKey), UserMaxLimit)

	db := internal.DB()
	err = db.Model(&model.Config{}).Where("key in ?", modeltype.ConfigKey).Limit(_min).Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func AdminGetConfigLst(limit int) (res []model.Config, err error) {
	_max := max(len(modeltype.ConfigKey), AdminMaxLimit)

	if limit <= 0 || limit > _max {
		limit = _max
	}

	db := internal.DB()
	err = db.Model(&model.Config{}).Where("key in ?", modeltype.ConfigKey).Limit(_max).Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
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
		var cfg = model.NewConfig(key, value)
		err := tx.Model(&model.Config{}).Where("key = ?", key).Order("create_at desc").FirstOrCreate(cfg).Error
		if err != nil {
			return err
		} else if cfg.Value == value {
			return nil
		}

		cfg.Value = value
		return tx.Save(cfg).Error
	})

}
