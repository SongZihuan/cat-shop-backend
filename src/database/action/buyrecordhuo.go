package action

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

func NewBuyRecord(user *model.User, wupin *model.Wupin, num modeltype.Total, username, userphone, userlocation, userwechat, useremail, userremark string) (*model.BuyRecord, error) {
	record := model.NewBuyRecord(user, wupin, num, username, userphone, userlocation, userwechat, useremail, userremark)

	db := internal.DB()
	err := db.Create(record).Error
	if err != nil {
		return nil, err
	}

	return record, nil
}

func NewBagBuyRecord(user *model.User, bag *model.Bag, username, userphone, userlocation, userwechat, useremail, userremark string) (*model.BuyRecord, error) {
	record := model.NewBagBuyRecord(user, bag, username, userphone, userlocation, userwechat, useremail, userremark)
	err := internal.DB().Transaction(func(tx *gorm.DB) (err error) {
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

func NewRepayRecord(user *model.User, record *model.BuyRecord) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	}

	record.Repay()

	db := internal.DB()
	err := db.Create(record).Error
	if err != nil {
		return err
	}

	return nil
}

func SetBuyRecordPayFail(user *model.User, record *model.BuyRecord) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	ok := record.PayFail()
	if !ok {
		return fmt.Errorf("pay error")
	}

	db := internal.DB()
	err := db.Save(record).Error
	if err != nil {
		return err
	}

	return nil
}

func SetBuyRecordPaySuccess(user *model.User, record *model.BuyRecord) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	if record.WupinID <= 0 || record.Wupin == nil {
		return fmt.Errorf("bad user")
	}

	ok := record.PaySuccess()
	if !ok {
		return fmt.Errorf("pay error")
	}

	return internal.DB().Transaction(func(tx *gorm.DB) error {
		err := tx.Save(record).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.Wupin).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.User).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func BuyRecordChangeUser(user *model.User, record *model.BuyRecord, username, userphone, userlocation, userwechat, useremail, userremark string) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	ok := record.ChangeUser(username, userphone, userlocation, userwechat, useremail, userremark)
	if !ok {
		return NewBuyRecordStatusError("状态错误")
	}

	return internal.DB().Transaction(func(tx *gorm.DB) error {
		err := tx.Save(record).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.Wupin).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.User).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func AdminBuyRecordChangeUser(user *model.User, record *model.BuyRecord, username, userphone, userlocation, userwechat, useremail, userremark string) error {
	return BuyRecordChangeUser(user, record, username, userphone, userlocation, userwechat, useremail, userremark)
}

func AdminBuyRecordChangeShop(user *model.User, record *model.BuyRecord, shopname, shopphone, shoplocation, shopwechat, shopemail, shopremark string) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	ok := record.ChangeShop(shopname, shopphone, shoplocation, shopwechat, shopemail, shopremark)
	if !ok {
		return NewBuyRecordStatusError("状态错误")
	}

	return internal.DB().Transaction(func(tx *gorm.DB) error {
		err := tx.Save(record).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.Wupin).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.User).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func AdminBuyRecordAcceptQuxiao(user *model.User, record *model.BuyRecord) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	ok := record.AcceptQuXiao()
	if !ok {
		return NewBuyRecordStatusError("状态错误")
	}

	return internal.DB().Transaction(func(tx *gorm.DB) error {
		err := tx.Save(record).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.Wupin).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.User).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func AdminBuyRecordAcceptTuiHuo(user *model.User, record *model.BuyRecord, accept bool) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	if accept {
		if ok := record.AcceptTuiHuo(); !ok {
			return NewBuyRecordStatusError("状态错误")
		}
	} else {
		if ok := record.NotAcceptTuiHuo(); !ok {
			return NewBuyRecordStatusError("状态错误")
		}
	}

	return internal.DB().Transaction(func(tx *gorm.DB) error {
		err := tx.Save(record).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.Wupin).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.User).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func AdminBuyRecordFaHuoCheHui(user *model.User, record *model.BuyRecord) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	ok := record.CheHuiFaHuo()
	if !ok {
		return NewBuyRecordStatusError("状态错误")
	}

	return internal.DB().Transaction(func(tx *gorm.DB) error {
		err := tx.Save(record).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.Wupin).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.User).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func AdminBuyRecordTuiHuoDaoHuo(user *model.User, record *model.BuyRecord) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	ok := record.TuiHuoDaoHuo()
	if !ok {
		return fmt.Errorf("status error")
	}

	return internal.DB().Transaction(func(tx *gorm.DB) error {
		err := tx.Save(record).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.Wupin).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.User).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func AdminBuyRecordDaoHuo(user *model.User, record *model.BuyRecord) error {
	return BuyRecordDaoHuo(user, record)
}

func BuyRecordDaoHuo(user *model.User, record *model.BuyRecord) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	ok := record.DaoHuo()
	if !ok {
		return fmt.Errorf("status error")
	}

	return internal.DB().Transaction(func(tx *gorm.DB) error {
		err := tx.Save(record).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.Wupin).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.User).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func BuyRecordPingJia(user *model.User, record *model.BuyRecord, isGood bool) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	ok := record.PingJia(isGood)
	if !ok {
		return fmt.Errorf("status error")
	}

	return internal.DB().Transaction(func(tx *gorm.DB) error {
		err := tx.Save(record).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.Wupin).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.User).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func BuyRecordQuXiaoFahuo(user *model.User, record *model.BuyRecord, accept bool) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	if accept {
		if ok := record.QuXiaoFaHuo(); !ok {
			return fmt.Errorf("status error")
		}
	} else {
		if ok := record.NotQuXiaoFaHuo(); !ok {
			return fmt.Errorf("status error")
		}
	}

	return internal.DB().Transaction(func(tx *gorm.DB) error {
		err := tx.Save(record).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.Wupin).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.User).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func AdminBuyRecordQuXiaoPay(user *model.User, record *model.BuyRecord) error {
	return BuyRecordQuXiaoPay(user, record)
}

func BuyRecordQuXiaoPay(user *model.User, record *model.BuyRecord) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}
	ok := record.QuXiaoPay()
	if !ok {
		return NewBuyRecordStatusError("状态错误")
	}

	return internal.DB().Transaction(func(tx *gorm.DB) error {
		err := tx.Save(record).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.Wupin).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.User).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func AdminBuyRecordTuiHuoShenQing(user *model.User, record *model.BuyRecord) error {
	return BuyRecordTuiHuoShenQing(user, record)
}

func BuyRecordTuiHuoShenQing(user *model.User, record *model.BuyRecord) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	ok := record.TuiHuoShenQing()
	if !ok {
		return NewBuyRecordStatusError("状态错误")
	}

	return internal.DB().Transaction(func(tx *gorm.DB) error {
		err := tx.Save(record).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.Wupin).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.User).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func AdminBuyRecordFaHuoDengJi(user *model.User, record *model.BuyRecord, kuaidi string, kuaidinum string) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	ok := record.FaHuoDengJi(kuaidi, kuaidinum)
	if !ok {
		return NewBuyRecordStatusError("状态错误")
	}

	return internal.DB().Transaction(func(tx *gorm.DB) error {
		err := tx.Save(record).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.Wupin).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.User).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func BuyRecordTuiHuoDengJi(user *model.User, record *model.BuyRecord, kuaidi string, kuaidinum string) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	ok := record.TuiHuoDengJi(kuaidi, kuaidinum)
	if !ok {
		return NewBuyRecordStatusError("状态错误")
	}

	return internal.DB().Transaction(func(tx *gorm.DB) error {
		err := tx.Save(record).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.Wupin).Error
		if err != nil {
			return err
		}

		err = tx.Save(record.User).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func AdminBuyRecordTuiHuoDengJi(user *model.User, record *model.BuyRecord, kuaidi string, kuaidinum string) error {
	return BuyRecordTuiHuoDengJi(user, record, kuaidi, kuaidinum)
}
