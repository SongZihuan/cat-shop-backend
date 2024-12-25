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

func GetBuyRecordListByUser(user *model.User, limit int, offset int) ([]model.BuyRecord, error) {
	return GetBuyRecordListByUserID(user.ID, limit, offset)
}

func GetBuyRecordListByUserID(userID uint, limit int, offset int) ([]model.BuyRecord, error) {
	var res []model.BuyRecord

	db := database.DB()
	err := db.Model(model.BuyRecord{}).Joins("Wupin").Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetBuyRecordListByPageByUser(user *model.User, page int, pagesize int) ([]model.BuyRecord, error) {
	return GetBuyRecordListByPageByUserID(user.ID, page, pagesize)
}

func GetBuyRecordListByPageByUserID(userID uint, page int, pagesize int) ([]model.BuyRecord, error) {
	var res []model.BuyRecord

	db := database.DB()
	err := db.Model(model.BuyRecord{}).Joins("Wupin").Where("user_id = ?", userID).Limit(pagesize).Offset((page - 1) * pagesize).Find(res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetBuyRecordLCountByPageByUser(user *model.User) (int, error) {
	return GetBuyRecordCountByPageByUserID(user.ID)
}

func GetBuyRecordCountByPageByUserID(userID uint) (int, error) {
	type count struct {
		count int `gorm:"column:count"`
	}

	var res count

	db := database.DB()
	err := db.Model(model.BuyRecord{}).Select("COUNT(*) as count").Where("user_id = ?", userID).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return res.count, nil
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

func SetBuyRecordPaySuccess(user *model.User, record *model.BuyRecord) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	}

	ok := record.PaySuccess()
	if !ok {
		return fmt.Errorf("pay error")
	}

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
	}

	record.Repay()

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

func BuyRecordChangeUser(user *model.User, record *model.BuyRecord, username, userphone, userlocation, userwechat, useremail, userremark string) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	}

	ok := record.ChangeUser(username, userphone, userlocation, userwechat, useremail, userremark)
	if !ok {
		return NewBuyRecordStatusError("状态错误")
	}

	return database.DB().Save(record).Error
}

func BuyRecordDaoHuo(user *model.User, record *model.BuyRecord) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	}

	ok := record.DaoHuo()
	if !ok {
		return NewBuyRecordStatusError("状态错误")
	}

	return database.DB().Save(record).Error
}

func BuyRecordPingJia(user *model.User, record *model.BuyRecord, isGood bool) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	}

	ok := record.PingJia(isGood)
	if !ok {
		return NewBuyRecordStatusError("状态错误")
	}

	return database.DB().Save(record).Error
}

func BuyRecordQuXiaoFahuo(user *model.User, record *model.BuyRecord) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	}

	ok := record.QuXiaoFahuo()
	if !ok {
		return NewBuyRecordStatusError("状态错误")
	}

	return database.DB().Save(record).Error
}

func BuyRecordQuXiaoPay(user *model.User, record *model.BuyRecord) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	}

	ok := record.QuXiaoPay()
	if !ok {
		return NewBuyRecordStatusError("状态错误")
	}

	return database.DB().Save(record).Error
}

func BuyRecordTuiHuoShenQing(user *model.User, record *model.BuyRecord) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	}

	ok := record.TuiHuoShenQing()
	if !ok {
		return NewBuyRecordStatusError("状态错误")
	}

	return database.DB().Save(record).Error
}

func BuyRecordTuiHuoDengJi(user *model.User, record *model.BuyRecord, kuaidi string, kuaidinum string) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	}

	ok := record.TuiHuoDengJi(kuaidi, kuaidinum)
	if !ok {
		return NewBuyRecordStatusError("状态错误")
	}

	return database.DB().Save(record).Error
}
