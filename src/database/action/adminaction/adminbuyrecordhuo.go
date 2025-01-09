package adminaction

import (
	"fmt"
	error2 "github.com/SongZihuan/cat-shop-backend/src/database/action/error"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"gorm.io/gorm"
)

func AdminBuyRecordChangeUser(user *model.User, record *model.BuyRecord, username, userphone, userlocation, userwechat, useremail, userremark string) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	ok := record.ChangeUser(username, userphone, userlocation, userwechat, useremail, userremark)
	if !ok {
		return error2.NewBuyRecordStatusError("状态错误")
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

func AdminBuyRecordChangeShop(user *model.User, record *model.BuyRecord, shopname, shopphone, shoplocation, shopwechat, shopemail, shopremark string) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	ok := record.ChangeShop(shopname, shopphone, shoplocation, shopwechat, shopemail, shopremark)
	if !ok {
		return error2.NewBuyRecordStatusError("状态错误")
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
		return error2.NewBuyRecordStatusError("状态错误")
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
			return error2.NewBuyRecordStatusError("状态错误")
		}
	} else {
		if ok := record.NotAcceptTuiHuo(); !ok {
			return error2.NewBuyRecordStatusError("状态错误")
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
		return error2.NewBuyRecordStatusError("状态错误")
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

func AdminBuyRecordQuXiaoPay(user *model.User, record *model.BuyRecord) error {
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}
	ok := record.QuXiaoPay()
	if !ok {
		return error2.NewBuyRecordStatusError("状态错误")
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
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	ok := record.TuiHuoShenQing()
	if !ok {
		return error2.NewBuyRecordStatusError("状态错误")
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
		return error2.NewBuyRecordStatusError("状态错误")
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
	if record.UserID != user.ID {
		return fmt.Errorf("bad user")
	} else {
		record.BindUser(user)
	}

	ok := record.TuiHuoDengJi(kuaidi, kuaidinum)
	if !ok {
		return error2.NewBuyRecordStatusError("状态错误")
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
