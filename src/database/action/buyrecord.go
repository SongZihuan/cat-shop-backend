package action

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
)

func GetBuyRecordByID(recordID uint) (*model.BuyRecord, error) {
	var record = new(model.BuyRecord)

	if recordID <= 0 {
		return nil, ErrNotFound
	}

	db := database.DB()
	err := db.Model(model.BuyRecord{}).Where("id = ?", recordID).First(record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return record, nil
}

func SetBuyRecordPayFail(record *model.BuyRecord) error {
	record.Status = modeltype.PayCheckFail

	db := database.DB()
	err := db.Save(record).Error
	if err != nil {
		return err
	}

	return nil
}

var SetBuyRecordWaitFahuo = SetBuyRecordPaySuccess

func SetBuyRecordPaySuccess(record *model.BuyRecord) error {
	record.Status = modeltype.WaitFahuo

	db := database.DB()
	err := db.Save(record).Error
	if err != nil {
		return err
	}

	return nil
}
