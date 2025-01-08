package model

import (
	"database/sql"
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"net/url"
	"time"
)

func NewBuyRecord(user *User, wupin *Wupin, num modeltype.Total, username, userphone, userlocation, userwechat, useremail, userremark string) *BuyRecord {
	if wupin.IsWupinCanNotSale() {
		panic("wupin not sale")
	}

	if num <= 0 {
		panic("num is bad")
	}

	return &BuyRecord{
		Status:             modeltype.WaitPay,
		UserID:             user.ID,
		WupinID:            wupin.ID,
		ClassID:            wupin.ClassID,
		Num:                num,
		Price:              wupin.GetPrice(),
		TotalPrice:         wupin.GetPriceTotal(num),
		XiaDanTime:         time.Now(),
		FuKuanTime:         utils.SqlTimeNow(),
		FaHuoTime:          utils.SqlTimeNow(),
		ShouHuoTime:        utils.SqlTimeNow(),
		PingJiaTime:        utils.SqlTimeNow(),
		TuiHuoShenQingTime: utils.SqlTimeNow(),
		QueRenTuiHuoTime:   utils.SqlTimeNow(),
		DengJiTuiHuoTime:   utils.SqlTimeNow(),
		TuiHuoTime:         utils.SqlTimeNow(),
		QuXiaoTime:         utils.SqlTimeNow(),
		FaHuoKuaiDi:        sql.NullString{Valid: false},
		FaHuoKuaiDiNum:     sql.NullString{Valid: false},
		TuiHuoKuaiDi:       sql.NullString{Valid: false},
		TuiHuoKuaiDiNum:    sql.NullString{Valid: false},
		IsGood:             sql.NullBool{Valid: false},

		WupinName:    wupin.Name,
		WupinPic:     wupin.Pic,
		WupinClassID: wupin.ClassID,
		WupinClass:   wupin.Class,
		WupinTag:     wupin.Tag,

		WupinHotPrice:  wupin.HotPrice,
		WupinRealPrice: wupin.RealPrice,

		WupinInfo:     wupin.Info,
		WupinRen:      wupin.Ren,
		WupinPhone:    wupin.Phone,
		WupinWeChat:   wupin.WeChat,
		WupinEmail:    wupin.Email,
		WupinLocation: wupin.Location,

		WupinBuyTotal:   wupin.BuyTotal,
		WupinBuyDaoHuo:  wupin.BuyDaoHuo,
		WupinBuyGood:    wupin.BuyGood,
		WupinBuyPrice:   wupin.BuyPrice,
		WupinBuyPingJia: wupin.BuyPingjia,
		WupinBuyJian:    wupin.BuyJian,
		WupinHot:        wupin.Hot,

		UserName:     username,
		UserPhone:    userphone,
		UserLocation: userlocation,
		UserWeChat:   sql.NullString{String: userwechat, Valid: len(userwechat) != 0},
		UserEmail:    sql.NullString{String: useremail, Valid: len(useremail) != 0},
		UserRemark:   sql.NullString{String: userremark, Valid: len(userremark) != 0},

		ShopName:     wupin.Ren,
		ShopPhone:    wupin.Phone,
		ShopWeChat:   wupin.WeChat,
		ShopEmail:    wupin.Email,
		ShopLocation: userlocation,
		ShopRemark:   sql.NullString{Valid: false},

		WupinDown: wupin.IsWupinDown(),
	}
}

func NewBagBuyRecord(user *User, bag *Bag, username, userphone, userlocation, userwechat, useremail, userremark string) *BuyRecord {
	if bag.IsBagCanNotSale() {
		panic("bag not sale")
	}

	wupin := bag.Wupin
	if wupin == nil {
		panic("wupin is nil")
	} else if wupin.IsWupinCanNotSale() {
		panic("wupin not sale")
	}

	return &BuyRecord{
		Status:             modeltype.WaitPay,
		UserID:             user.ID,
		WupinID:            wupin.ID,
		ClassID:            wupin.ClassID,
		Num:                bag.Num,
		Price:              wupin.GetPrice(),
		TotalPrice:         wupin.GetPriceTotal(bag.Num),
		XiaDanTime:         time.Now(),
		FuKuanTime:         utils.SqlTimeNow(),
		FaHuoTime:          utils.SqlTimeNow(),
		ShouHuoTime:        utils.SqlTimeNow(),
		PingJiaTime:        utils.SqlTimeNow(),
		TuiHuoShenQingTime: utils.SqlTimeNow(),
		QueRenTuiHuoTime:   utils.SqlTimeNow(),
		DengJiTuiHuoTime:   utils.SqlTimeNow(),
		TuiHuoTime:         utils.SqlTimeNow(),
		QuXiaoTime:         utils.SqlTimeNow(),
		FaHuoKuaiDi:        sql.NullString{Valid: false},
		FaHuoKuaiDiNum:     sql.NullString{Valid: false},
		TuiHuoKuaiDi:       sql.NullString{Valid: false},
		TuiHuoKuaiDiNum:    sql.NullString{Valid: false},
		IsGood:             sql.NullBool{Valid: false},

		WupinName:    wupin.Name,
		WupinPic:     wupin.Pic,
		WupinClassID: wupin.ClassID,
		WupinClass:   wupin.Class,
		WupinTag:     wupin.Tag,

		WupinHotPrice:  wupin.HotPrice,
		WupinRealPrice: wupin.RealPrice,

		WupinInfo:     wupin.Info,
		WupinRen:      wupin.Ren,
		WupinPhone:    wupin.Phone,
		WupinWeChat:   wupin.WeChat,
		WupinEmail:    wupin.Email,
		WupinLocation: wupin.Location,

		WupinBuyTotal:   wupin.BuyTotal,
		WupinBuyDaoHuo:  wupin.BuyDaoHuo,
		WupinBuyGood:    wupin.BuyGood,
		WupinBuyPrice:   wupin.BuyPrice,
		WupinBuyPingJia: wupin.BuyPingjia,
		WupinBuyJian:    wupin.BuyJian,
		WupinHot:        wupin.Hot,

		UserName:     username,
		UserPhone:    userphone,
		UserLocation: userlocation,
		UserWeChat:   sql.NullString{String: userwechat, Valid: len(userwechat) != 0},
		UserEmail:    sql.NullString{String: useremail, Valid: len(useremail) != 0},
		UserRemark:   sql.NullString{String: userremark, Valid: len(userremark) != 0},

		ShopName:     wupin.Ren,
		ShopPhone:    wupin.Phone,
		ShopWeChat:   wupin.WeChat,
		ShopEmail:    wupin.Email,
		ShopLocation: userlocation,
		ShopRemark:   sql.NullString{Valid: false},

		WupinDown: wupin.IsWupinDown(),
		ClassDown: wupin.isClassDown(),
	}
}

func (r *BuyRecord) BindUser(u *User) {
	if r.UserID != u.ID {
		panic("bad user id")
	} else if r.User == nil {
		r.User = u
	}
}

func (r *BuyRecord) Repay() {
	if r.Status == modeltype.PayCheckFail {
		r.Status = modeltype.WaitPay
	}
}

func (r *BuyRecord) GetNewPayUrl(pft modeltype.PayFromType, redirect string) string {
	if !config.IsReady() {
		panic("config is not ready")
	}

	if r.Status != modeltype.WaitPay {
		return ""
	}

	cfg := config.Config()
	basePath := cfg.Yaml.Front.BasePath + cfg.Yaml.Front.TestPath + cfg.Yaml.Front.TestPayPath
	query := r.getPayUrlQuery(pft, modeltype.NewPay, redirect)

	return basePath + "?" + query
}

func (r *BuyRecord) GetRepayPayUrl(pft modeltype.PayFromType, redirect string) string {
	if !config.IsReady() {
		panic("config is not ready")
	}

	if r.Status != modeltype.WaitPay { // 此处应该是WaitPayCheck，因为在调用该函数之前旧调用了Repay函数，该函数会重置Status为WaitPayCheck
		return ""
	}

	cfg := config.Config()
	basePath := cfg.Yaml.Front.BasePath + cfg.Yaml.Front.TestPath + cfg.Yaml.Front.TestPayPath
	query := r.getPayUrlQuery(pft, modeltype.Repay, redirect)

	return basePath + "?" + query
}

func (r *BuyRecord) GetBagPayUrl(pft modeltype.PayFromType, redirect string) string {
	if !config.IsReady() {
		panic("config is not ready")
	}

	if r.Status != modeltype.WaitPay {
		return ""
	}

	cfg := config.Config()
	basePath := cfg.Yaml.Front.BasePath + cfg.Yaml.Front.TestPath + cfg.Yaml.Front.TestPayPath
	query := r.getPayUrlQuery(pft, modeltype.BagPay, redirect)

	return basePath + "?" + query
}

func (r *BuyRecord) getPayUrlQuery(pft modeltype.PayFromType, pt modeltype.PayType, redirect string) string {
	pftn, ok := modeltype.PayFromTypeToName[pft]
	if !ok {
		return ""
	}

	ptn, ok := modeltype.PayTypeToName[pt]
	if !ok {
		return ""
	}

	v := url.Values{}
	v.Add(string(modeltype.PayFromTypeKey), pftn)
	v.Add(string(modeltype.PayBuyRecordIdKey), fmt.Sprintf("%d", r.ID))
	v.Add(string(modeltype.PayRedirectKey), redirect)
	v.Add(string(modeltype.PayTypeKey), ptn)

	return v.Encode()
}

func (r *BuyRecord) PaySuccess() bool {
	if r.Status == modeltype.WaitPay {
		ok := r.Wupin.BuyNow(r)
		if !ok {
			return false
		}

		ok = r.User.BuyNow(r)
		if !ok {
			return false
		}

		r.Status = modeltype.WaitFahuo
		r.FuKuanTime = utils.SqlTimeNow()
		return true
	}
	return false
}

func (r *BuyRecord) PayFail() bool {
	if r.Status == modeltype.WaitPay {
		r.Status = modeltype.PayCheckFail
		return false
	}
	return false
}

func (r *BuyRecord) ChangeUser(username, userphone, userlocation, userwechat, useremail, userremark string) bool {
	if r.Status == modeltype.WaitShouHuo || r.Status == modeltype.WaitPay || r.Status == modeltype.PayCheckFail {
		r.UserName = username
		r.UserPhone = userphone
		r.UserLocation = userlocation
		r.UserWeChat = sql.NullString{String: userwechat, Valid: len(userwechat) != 0}
		r.UserEmail = sql.NullString{String: useremail, Valid: len(useremail) != 0}
		r.UserRemark = sql.NullString{String: userremark, Valid: len(userremark) != 0}
		return true
	}
	return false
}

func (r *BuyRecord) ChangeShop(shopname, shopphone, shoplocation, shopwechat, shopemail, shopremark string) bool {
	r.ShopName = shopname
	r.ShopPhone = shopphone
	r.ShopLocation = shoplocation
	r.ShopWeChat = sql.NullString{String: shopwechat, Valid: len(shopwechat) != 0}
	r.ShopEmail = sql.NullString{String: shopemail, Valid: len(shopemail) != 0}
	r.ShopRemark = sql.NullString{String: shopremark, Valid: len(shopremark) != 0}
	return true
}

func (r *BuyRecord) AcceptTuiHuo() bool {
	if r.Status == modeltype.WaitTuiHuoShouHuo {
		return true
	} else if r.Status == modeltype.TuiHuoCheck {
		r.Status = modeltype.WaitTuiHuoShouHuo
		r.QueRenTuiHuoTime = utils.SqlTimeNow()
		return true
	}
	return false
}

func (r *BuyRecord) NotAcceptTuiHuo() bool {
	if r.Status == modeltype.WaitTuiHuoShouHuo {
		return true
	} else if r.Status == modeltype.TuiHuoCheck {
		r.Status = modeltype.TuiHuoFail
		r.QueRenTuiHuoTime = utils.SqlTimeNow()
		return true
	}
	return false
}

func (r *BuyRecord) AcceptQuXiao() bool {
	if r.Status == modeltype.QuXiao {
		return true
	} else if r.Status == modeltype.PayCheckFail || r.Status == modeltype.WaitPay {
		r.Status = modeltype.QuXiao
		r.QuXiaoTime = utils.SqlTimeNow()
		return true
	} else if r.Status == modeltype.WaitFahuo || r.Status == modeltype.WaitQuXiao {
		ok := r.Wupin.BuyQuXiao(r)
		if !ok {
			return false
		}

		ok = r.User.BuyQuXiao(r)
		if !ok {
			return false
		}

		r.Status = modeltype.QuXiao
		r.QuXiaoTime = utils.SqlTimeNow()
		return true
	}
	return false
}

func (r *BuyRecord) CheHuiFaHuo() bool {
	if r.Status == modeltype.WaitFahuo {
		return true
	} else if r.Status == modeltype.WaitShouHuo {
		r.Status = modeltype.WaitFahuo
		r.FaHuoTime = utils.SqlTimeNull()
		return true
	}
	return false
}

func (r *BuyRecord) DaoHuo() bool {
	if r.Status == modeltype.WaitPingJia {
		return true
	} else if r.Status == modeltype.WaitShouHuo {
		ok := r.Wupin.Daohuo()
		if !ok {
			return false
		}

		ok = r.User.Daohuo()
		if !ok {
			return false
		}

		r.Status = modeltype.WaitPingJia
		r.ShouHuoTime = utils.SqlTimeNow()
		return true
	}
	return false
}

func (r *BuyRecord) TuiHuoDaoHuo() bool {
	if r.Status == modeltype.TuiHuo {
		return true
	} else if r.Status == modeltype.WaitTuiHuoShouHuo {
		ok := r.Wupin.TuiHuoAfterFaHuo(r)
		if !ok {
			return false
		}

		ok = r.User.TuiHuoAfterFaHuo(r)
		if !ok {
			return false
		}

		r.Status = modeltype.TuiHuo
		r.TuiHuoTime = utils.SqlTimeNow()
		return true
	}
	return false
}

func (r *BuyRecord) PingJia(isGood bool) bool {
	if r.Status == modeltype.YiPingJia {
		return true
	} else if r.Status == modeltype.WaitPingJia {
		ok := r.Wupin.PingJia(isGood)
		if !ok {
			return false
		}

		ok = r.User.PingJia(isGood)
		if !ok {
			return false
		}

		r.Status = modeltype.YiPingJia
		r.IsGood = sql.NullBool{Bool: isGood, Valid: true}
		r.PingJiaTime = utils.SqlTimeNow()
		return true
	}
	return false
}

func (r *BuyRecord) QuXiaoFaHuo() bool {
	if r.Status == modeltype.QuXiao {
		return true
	} else if r.Status == modeltype.PayCheckFail || r.Status == modeltype.WaitPay {
		r.Status = modeltype.QuXiao
		r.QuXiaoTime = utils.SqlTimeNow()
		return true
	} else if r.Status == modeltype.WaitFahuo {
		r.Status = modeltype.WaitQuXiao
		return true
	}
	return false
}

func (r *BuyRecord) NotQuXiaoFaHuo() bool {
	if r.Status == modeltype.WaitFahuo || r.Status == modeltype.PayCheckFail || r.Status == modeltype.WaitPay {
		r.QuXiaoTime = utils.SqlTimeNull()
		return true
	} else if r.Status == modeltype.WaitQuXiao { // 若WaitQuXiao前的状态时未支付，那么用户申请取消会直接进入QuXiao状态，不会进入WaitQuXiao状态
		r.Status = modeltype.WaitFahuo
		r.QuXiaoTime = utils.SqlTimeNull()
		return true
	}
	return false
}

func (r *BuyRecord) QuXiaoPay() bool {
	if r.Status == modeltype.QuXiao {
		return true
	} else if r.Status == modeltype.PayCheckFail || r.Status == modeltype.WaitPay {
		r.Status = modeltype.QuXiao
		r.QuXiaoTime = utils.SqlTimeNow()
		return true
	}
	return false
}

func (r *BuyRecord) TuiHuoShenQing() bool {
	if r.Status == modeltype.TuiHuoCheck {
		return true
	} else if r.Status == modeltype.WaitPingJia || r.Status == modeltype.YiPingJia || r.Status == modeltype.TuiHuoFail {
		r.Status = modeltype.TuiHuoCheck
		r.TuiHuoShenQingTime = utils.SqlTimeNow()
		return true
	} else if r.Status == modeltype.WaitFahuo { // wait 收货不支持直接退款
		r.Status = modeltype.TuiHuo

		ok := r.Wupin.TuiHuoBeforeFaHuo(r)
		if !ok {
			return false
		}

		ok = r.User.TuiHuoBeforeFaHuo(r)
		if !ok {
			return false
		}

		r.TuiHuoKuaiDi = sql.NullString{String: "未发货退货", Valid: true}
		r.TuiHuoKuaiDiNum = sql.NullString{String: "无快递单号", Valid: true}
		r.DengJiTuiHuoTime = utils.SqlTimeNow()
		r.QueRenTuiHuoTime = r.DengJiTuiHuoTime
		r.TuiHuoShenQingTime = r.DengJiTuiHuoTime
		r.TuiHuoTime = r.DengJiTuiHuoTime
		return true
	}
	return false
}

func (r *BuyRecord) TuiHuoDengJi(kuaidi string, kuaidinum string) bool {
	if r.Status == modeltype.WaitTuiHuoShouHuo {
		return true
	} else if r.Status == modeltype.WaitTuiHuoFahuo {
		r.Status = modeltype.WaitTuiHuoShouHuo
		r.DengJiTuiHuoTime = utils.SqlTimeNow()
		r.TuiHuoKuaiDi = sql.NullString{String: kuaidi, Valid: true}
		r.TuiHuoKuaiDiNum = sql.NullString{String: kuaidinum, Valid: true}
		return true
	}
	return false
}

func (r *BuyRecord) FaHuoDengJi(kuaidi string, kuaidinum string) bool {
	if r.Status == modeltype.WaitShouHuo {
		return true
	} else if r.Status == modeltype.WaitFahuo {
		r.Status = modeltype.WaitShouHuo
		r.FaHuoTime = utils.SqlTimeNow()
		r.FaHuoKuaiDi = sql.NullString{String: kuaidi, Valid: true}
		r.FaHuoKuaiDiNum = sql.NullString{String: kuaidinum, Valid: true}
		return true
	}
	return false
}
