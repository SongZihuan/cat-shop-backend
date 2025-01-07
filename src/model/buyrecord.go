package model

import (
	"database/sql"
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"gorm.io/gorm"
	"net/url"
	"time"
)

type BuyRecord struct {
	gorm.Model
	Status             modeltype.BuyStatus `gorm:"type:uint;not null;"`
	UserID             uint                `gorm:"not null"`
	User               *User               `gorm:"foreignKey:UserID"`
	WupinID            uint                `gorm:"not null"`
	Wupin              *Wupin              `gorm:"foreignKey:WupinID"`
	ClassID            uint                `gorm:"not null"`
	Class              *Class              `gorm:"foreignKey:ClassID"`
	Num                modeltype.Total     `gorm:"type:uint;not null"`
	Price              modeltype.Price     `gorm:"type:uint;not null"`
	TotalPrice         modeltype.Price     `gorm:"type:uint;not null"`
	XiaDanTime         time.Time           `gorm:"type:datetime;not null"`
	FuKuanTime         sql.NullTime        `gorm:"type:datetime;"`
	FaHuoTime          sql.NullTime        `gorm:"type:datetime;"`
	ShouHuoTime        sql.NullTime        `gorm:"type:datetime;"`
	PingJiaTime        sql.NullTime        `gorm:"type:datetime;"`
	TuiHuoShenQingTime sql.NullTime        `gorm:"type:datetime;"`
	DengJiTuiHuoTime   sql.NullTime        `gorm:"type:datetime;"`
	QueRenTuiHuoTime   sql.NullTime        `gorm:"type:datetime;"`
	TuiHuoTime         sql.NullTime        `gorm:"type:datetime;"`
	QuXiaoTime         sql.NullTime        `gorm:"type:datetime;"`
	FaHuoKuaiDi        sql.NullString      `gorm:"type:varchar(20);"`
	FaHuoKuaiDiNum     sql.NullString      `gorm:"type:varchar(50);"`
	TuiHuoKuaiDi       sql.NullString      `gorm:"type:varchar(20);"`
	TuiHuoKuaiDiNum    sql.NullString      `gorm:"type:varchar(50);"`
	IsGood             sql.NullBool        `gorm:"type:boolean;"`

	UserName     string         `gorm:"type:varchar(20);not null;"`
	UserPhone    string         `gorm:"type:varchar(30);not null"`
	UserWeChat   sql.NullString `gorm:"type:varchar(50);"`
	UserEmail    sql.NullString `gorm:"type:varchar(50);"`
	UserLocation string         `gorm:"type:varchar(200);not null;"`
	UserRemark   sql.NullString `gorm:"type:varchar(200);"`

	ShopName     string         `gorm:"type:varchar(20);not null;"`
	ShopPhone    string         `gorm:"type:varchar(30);not null"`
	ShopWeChat   sql.NullString `gorm:"type:varchar(50);"`
	ShopEmail    sql.NullString `gorm:"type:varchar(50);"`
	ShopLocation string         `gorm:"type:varchar(200);not null;"`
	ShopRemark   sql.NullString `gorm:"type:varchar(200);"`

	WupinName    string         `gorm:"type:varchar(20);not null"`
	WupinPic     string         `gorm:"type:varchar(150);not null"`
	WupinClassID uint           `gorm:"not null"`
	WupinClass   *Class         `gorm:"foreignKey:ClassID"`
	WupinTag     sql.NullString `gorm:"type:varchar(20)"`

	WupinHotPrice  modeltype.PriceNull `gorm:"type:uint;"`
	WupinRealPrice modeltype.Price     `gorm:"type:uint;not null;"`

	WupinInfo     string         `gorm:"type:text;not null"`
	WupinRen      string         `gorm:"type:varchar(20);not null"`
	WupinPhone    string         `gorm:"type:varchar(30);not null"`
	WupinWeChat   sql.NullString `gorm:"type:varchar(50);"`
	WupinEmail    sql.NullString `gorm:"type:varchar(50);"`
	WupinLocation string         `gorm:"type:varchar(200);not null"`

	WupinBuyTotal  modeltype.Total `gorm:"type:uint;not null"`
	WupinBuyDaoHuo modeltype.Total `gorm:"type:uint;not null"`
	WupinPingJia   modeltype.Total `gorm:"type:uint;not null"`
	WupinBuyGood   modeltype.Total `gorm:"type:uint;not null"`

	WupinHot  bool `gorm:"type:boolean;not null"`
	WupinDown bool `gorm:"type:boolean;not null"`
	ClassShow bool `gorm:"type:boolean;not null;"`
	ClassDown bool `gorm:"type:boolean;not null;"`
}

func NewBuyRecord(user *User, wupin *Wupin, num modeltype.Total, username, userphone, userlocation, userwechat, useremail, userremark string) *BuyRecord {
	return &BuyRecord{
		Status:     modeltype.WaitPayCheck,
		UserID:     user.ID,
		WupinID:    wupin.ID,
		ClassID:    wupin.ClassID,
		Num:        num,
		Price:      wupin.GetPrice(),
		TotalPrice: wupin.GetPriceTotal(num),
		XiaDanTime: time.Now(),

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

		WupinBuyTotal:  wupin.BuyTotal,
		WupinBuyDaoHuo: wupin.BuyDaoHuo,
		WupinPingJia:   wupin.BuyPingjia,
		WupinBuyGood:   wupin.BuyGood,

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
	}
}

func NewBagBuyRecord(user *User, bag *Bag, username, userphone, userlocation, userwechat, useremail, userremark string) *BuyRecord {
	wupin := bag.Wupin
	if wupin == nil {
		panic("wupin is nil")
	}

	return &BuyRecord{
		Status:     modeltype.WaitPayCheck,
		UserID:     user.ID,
		WupinID:    wupin.ID,
		ClassID:    wupin.ClassID,
		Num:        bag.Num,
		Price:      wupin.GetPrice(),
		TotalPrice: wupin.GetPriceTotal(bag.Num),
		XiaDanTime: time.Now(),

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

		WupinBuyTotal:  wupin.BuyTotal,
		WupinBuyDaoHuo: wupin.BuyDaoHuo,
		WupinPingJia:   wupin.BuyPingjia,
		WupinBuyGood:   wupin.BuyGood,

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
		r.Status = modeltype.WaitPayCheck
	}
}

func (r *BuyRecord) GetNewPayUrl(pft modeltype.PayFromType, redirect string) string {
	if !config.IsReady() {
		panic("config is not ready")
	}

	if r.Status != modeltype.WaitPayCheck {
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

	if r.Status != modeltype.WaitPayCheck { // 此处应该是WaitPayCheck，因为在调用该函数之前旧调用了Repay函数，该函数会重置Status为WaitPayCheck
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

	if r.Status != modeltype.WaitPayCheck {
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
	if r.WupinID <= 0 || r.Wupin == nil {
		return false
	} else if r.UserID <= 0 || r.User == nil {
		return false
	}

	if r.Status == modeltype.WaitPayCheck {
		ok := r.Wupin.BuyNow(r)
		if !ok {
			return false
		}

		ok = r.User.BuyNow(r)
		if !ok {
			return false
		}

		r.Status = modeltype.WaitFahuo
		r.FuKuanTime = utils.SqlNullNow()
		return true
	}
	return false
}

func (r *BuyRecord) PayFail() bool {
	if r.Status == modeltype.WaitPayCheck {
		r.Status = modeltype.PayCheckFail
		return false
	}
	return false
}

func (r *BuyRecord) ChangeUser(username, userphone, userlocation, userwechat, useremail, userremark string) bool {
	if r.Status == modeltype.WaitShouHuo || r.Status == modeltype.WaitPayCheck || r.Status == modeltype.PayCheckFail {
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

func (r *BuyRecord) DaoHuo() bool {
	if r.WupinID <= 0 || r.Wupin == nil {
		return false
	} else if r.UserID <= 0 || r.User == nil {
		return false
	}

	if r.Status == modeltype.WaitPingJia {
		return true
	} else if r.Status == modeltype.WaitShouHuo {
		ok := r.Wupin.Daohuo(r)
		if !ok {
			return false
		}

		ok = r.User.Daohuo(r)
		if !ok {
			return false
		}

		r.Status = modeltype.WaitPingJia
		r.ShouHuoTime = utils.SqlNullNow()
		return true
	}
	return false
}

func (r *BuyRecord) PingJia(isGood bool) bool {
	if r.WupinID <= 0 || r.Wupin == nil {
		return false
	} else if r.UserID <= 0 || r.User == nil {
		return false
	}

	if r.Status == modeltype.YiPingJia {
		return true
	} else if r.Status == modeltype.WaitPingJia {
		ok := r.Wupin.PingJia(r, isGood)
		if !ok {
			return false
		}

		ok = r.User.PingJia(r, isGood)
		if !ok {
			return false
		}

		r.Status = modeltype.YiPingJia
		r.IsGood = sql.NullBool{Bool: isGood, Valid: true}
		r.PingJiaTime = utils.SqlNullNow()
		return true
	}
	return false
}

func (r *BuyRecord) QuXiaoFahuo() bool {
	if r.WupinID <= 0 || r.Wupin == nil {
		return false
	}

	if r.Status == modeltype.QuXiao {
		return true
	} else if r.Status == modeltype.PayCheckFail || r.Status == modeltype.WaitPayCheck {
		r.Status = modeltype.QuXiao
		r.QuXiaoTime = utils.SqlNullNow()
		return true
	} else if r.Status == modeltype.WaitFahuo {
		r.Status = modeltype.CheckQuXiao
		r.QuXiaoTime = utils.SqlNullNow()
		return true
	}
	return false
}

func (r *BuyRecord) QuXiaoPay() bool {
	if r.Status == modeltype.QuXiao {
		return true
	} else if r.Status == modeltype.PayCheckFail || r.Status == modeltype.WaitPayCheck {
		r.Status = modeltype.QuXiao
		r.QuXiaoTime = utils.SqlNullNow()
		return true
	}
	return false
}

func (r *BuyRecord) TuiHuoShenQing() bool {
	if r.Status == modeltype.WaitPingJia || r.Status == modeltype.YiPingJia || r.Status == modeltype.TuiHuoFail {
		r.Status = modeltype.TuiHuoCheck
		r.TuiHuoShenQingTime = utils.SqlNullNow()
		return true
	}
	return false
}

func (r *BuyRecord) TuiHuoDengJi(kuaidi string, kuaidinum string) bool {
	if r.Status == modeltype.WaitTuiHuoShouHuo {
		return true
	} else if r.Status == modeltype.WaitTuiHuoFahuo {
		r.Status = modeltype.WaitTuiHuoShouHuo
		r.DengJiTuiHuoTime = utils.SqlNullNow()
		r.TuiHuoKuaiDi = sql.NullString{String: kuaidi, Valid: true}
		r.TuiHuoKuaiDiNum = sql.NullString{String: kuaidinum, Valid: true}
		return true
	}
	return false
}

func (r *BuyRecord) isClassDown() bool {
	if r.Class == nil {
		return r.ClassDown
	} else {
		if r.Class.ID != r.ClassID {
			panic("wupin id not equal")
		}

		return r.Class.IsClassDown()
	}
}

func (r *BuyRecord) isWupinDown() bool {
	if r.Wupin == nil {
		return r.WupinDown || r.ClassDown
	} else {
		if r.WupinID != r.Wupin.ID {
			panic("wupin id not equal")
		}

		return r.Wupin.IsWupinDown()
	}
}

func (r *BuyRecord) IsBuyRecordCanNotPay() bool {
	return r.isWupinDown() || r.isClassDown()
}

func (r *BuyRecord) IsBuyRecordCanPay() bool {
	return !r.IsBuyRecordCanNotPay()
}

func (*BuyRecord) IsBuyRecordDown() bool {
	return false
}
