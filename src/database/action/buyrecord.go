package action

import (
	"errors"
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/database"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
)

func GetBuyRecordByIDAndUser(user *model.User, recordID uint) (*model.BuyRecord, error) {
	record, err := GetBuyRecordByIDAndUserID(user.ID, recordID)
	if err != nil {
		return nil, err
	}

	record.User = user
	return record, nil
}

func GetBuyRecordByIDAndUserID(userID uint, recordID uint) (*model.BuyRecord, error) {
	var record = new(model.BuyRecord)

	if recordID <= 0 {
		return nil, ErrNotFound
	}

	db := database.DB()
	err := db.Model(model.BuyRecord{}).Joins("Wupin").Where("id = ?", recordID).Where("user_id = ?", userID).First(record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	if record.ClassID > 0 && record.WuPin.ClassID > 0 && record.ClassID == record.WuPin.ClassID {
		record.Class = record.WuPin.Class
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

func NewBuyRecord(user *model.User, wupin *model.WuPin, num modeltype.Total, username, userphone, userlocation, userwechat, useremail, userremark string) (*model.BuyRecord, error) {
	record := model.NewBuyRecord(user, wupin, num, username, userphone, userlocation, userwechat, useremail, userremark)

	db := database.DB()
	err := db.Create(record).Error
	if err != nil {
		return nil, err
	}

	return record, nil
}

func NewRepayRecord(user *model.User, record *model.BuyRecord) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.Repay()
	}

	db := database.DB()
	err := db.Create(record).Error
	if err != nil {
		return err
	}

	return nil
}

func NewBagBuyRecord(user *model.User, bag *model.Bag, username, userphone, userlocation, userwechat, useremail, userremark string) (*model.BuyRecord, error) {
	record := model.NewBagBuyRecord(user, bag, username, userphone, userlocation, userwechat, useremail, userremark)
	err := database.DB().Transaction(func(tx *gorm.DB) (err error) {
		bag.Num = 0

		err = tx.Create(record).Error
		if err != nil {
			return err
		}

		err = tx.Save(bag).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return record, nil
}
