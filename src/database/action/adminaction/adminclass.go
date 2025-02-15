package adminaction

import (
	"errors"
	error2 "github.com/SongZihuan/cat-shop-backend/src/database/action/error"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
	"time"
)

func AdminGetClassListByPage(page int, pagesize int) (res []model.Class, err error) {
	db := internal.DB()
	err = db.Model(&model.Class{}).Where("id != ?", modeltype.ClassEmptyID).Limit(pagesize).Offset((page - 1) * pagesize).Order("create_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AdminGetClassCount() (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	var res count
	err := internal.DB().Model(&model.Class{}).Select("COUNT(*) as count").Where("id != ?", modeltype.ClassEmptyID).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return res.count, nil
}

func AdminAddClass(name string, show bool, down bool) error {
	db := internal.DB()

	err := adminCreateEmptyClass(db)
	if err != nil {
		return err
	}

	cls := model.NewClass(name, show, down)
	return db.Create(cls).Error
}

func AdminGetClass(id uint) (*model.Class, error) {
	return adminGetClass(id, internal.DB())
}

func adminGetClass(id uint, db *gorm.DB) (*model.Class, error) {
	if id == modeltype.ClassEmptyID {
		return &model.Class{
			Model: gorm.Model{
				ID:        modeltype.ClassEmptyID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:      modeltype.ClassEmptyName,
			Show:      modeltype.ClassEmptyShow,
			ClassDown: modeltype.ClassEmptyDown,
		}, nil
	}

	var res = new(model.Class)
	err := db.Model(&model.Class{}).Where("id = ?", id).Order("create_at desc").First(res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AdminUpdateClass(id uint, name string, show bool, down bool) error {
	if id == modeltype.ClassEmptyID {
		return nil
	}

	err := internal.DB().Transaction(func(tx *gorm.DB) error {
		err := adminCreateEmptyClass(tx)
		if err != nil {
			return err
		}

		cls, err := adminGetClass(id, tx)
		if err != nil {
			return err
		}

		needUpdate := cls.UpdateInfo(name, show, down)
		err = tx.Save(cls).Error
		if err != nil {
			return err
		}

		if !needUpdate {
			return nil
		}

		err = tx.Model(&model.Class{}).Where("class_id = ?", cls.ID).Update("class_down", cls.ClassDown).Error
		if err != nil {
			return err
		}

		err = tx.Model(&model.Class{}).Where("class_id = ?", cls.ID).Update("class_show", cls.Show).Error
		if err != nil {
			return err
		}

		err = tx.Model(&model.Bag{}).Where("class_id = ?", cls.ID).Update("class_down", cls.ClassDown).Error
		if err != nil {
			return err
		}

		err = tx.Model(&model.Bag{}).Where("class_id = ?", cls.ID).Update("class_show", cls.Show).Error
		if err != nil {
			return err
		}

		err = tx.Model(&model.BuyRecord{}).Where("class_id = ?", cls.ID).Update("class_down", cls.ClassDown).Error
		if err != nil {
			return err
		}

		err = tx.Model(&model.BuyRecord{}).Where("class_id = ?", cls.ID).Update("class_show", cls.Show).Error
		if err != nil {
			return err
		}

		return nil
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return error2.ErrNotFound
	}
	return err
}

// 内部使用
func adminCreateEmptyClass(db *gorm.DB) error {
	cls := model.NewEmptyClass()
	err := db.Model(&model.ClassM{}).Where("id = ?", modeltype.ClassEmptyID).Limit(1).FirstOrCreate(cls).Error
	if err != nil {
		return err
	}

	if cls.ResetIsEmpty() {
		err := db.Save(cls).Error
		if err != nil {
			return err
		}
	}

	return nil
}
