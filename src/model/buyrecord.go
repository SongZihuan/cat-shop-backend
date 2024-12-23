package model

import (
	"database/sql"
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/config"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
	"net/url"
	"time"
)

type BuyRecord struct {
	gorm.Model
	Status           modeltype.BuyStatus `gorm:"type:uint;not null;"`
	UserID           uint                `gorm:"not null"`
	User             *User               `gorm:"foreignKey:UserID"`
	WuPinID          uint                `gorm:"not null"`
	WuPin            *WuPin              `gorm:"foreignKey:WuPinID"`
	ClassID          uint                `gorm:"not null"`
	Class            *Class              `gorm:"foreignKey:ClassID"`
	Num              modeltype.Total     `gorm:"type:uint;not null"`
	Price            modeltype.Price     `gorm:"type:uint;not null"`
	TotalPrice       modeltype.Price     `gorm:"type:uint;not null"`
	XiaDanTime       time.Time           `gorm:"type:datetime;not null"`
	FuKuanTime       sql.NullTime        `gorm:"type:datetime;"`
	FaHuoTime        sql.NullTime        `gorm:"type:datetime;"`
	ShouHuoTime      sql.NullTime        `gorm:"type:datetime;"`
	PingJiaTime      sql.NullTime        `gorm:"type:datetime;"`
	DengJiTuiHuoTime sql.NullTime        `gorm:"type:datetime;"`
	QueRenTuiHuoTime sql.NullTime        `gorm:"type:datetime;"`
	TuiHuoTime       sql.NullTime        `gorm:"type:datetime;"`
	QuXiaoTime       sql.NullTime        `gorm:"type:datetime;"`
	FaHuoKuaiDi      sql.NullString      `gorm:"type:varchar(20);"`
	FaHuoKuaiDiNum   sql.NullString      `gorm:"type:varchar(50);"`
	TuiHuoKuaiDi     sql.NullString      `gorm:"type:varchar(20);"`
	TuiHuoKuaiDiNum  sql.NullString      `gorm:"type:varchar(50);"`
	IsGood           sql.NullBool        `gorm:"type:boolean;"`

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
}

func NewBuyRecord(user *User, wupin *WuPin, num modeltype.Total, username, userphone, userlocation, userwechat, useremail, userremark string) *BuyRecord {
	return &BuyRecord{
		Status:     modeltype.WaitPayCheck,
		UserID:     user.ID,
		WuPinID:    wupin.ID,
		ClassID:    wupin.ClassID,
		Num:        num,
		Price:      wupin.GetPrice(),
		TotalPrice: wupin.GetPriceTotal(num),
		XiaDanTime: time.Now(),

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
	wupin := bag.WuPin
	if wupin == nil {
		panic("wupin is nil")
	}

	return &BuyRecord{
		Status:     modeltype.WaitPayCheck,
		UserID:     user.ID,
		WuPinID:    wupin.ID,
		ClassID:    wupin.ClassID,
		Num:        bag.Num,
		Price:      wupin.GetPrice(),
		TotalPrice: wupin.GetPriceTotal(bag.Num),
		XiaDanTime: time.Now(),

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
	if r.Status == modeltype.WaitPingJia {
		return true
	} else if r.Status == modeltype.WaitShouHuo {
		r.Status = modeltype.WaitPingJia
		return true
	}
	return false
}

func (r *BuyRecord) PingJia(isGood bool) bool {
	if r.Status == modeltype.YiPingJia {
		return true
	} else if r.Status == modeltype.WaitPingJia {
		r.Status = modeltype.YiPingJia
		r.IsGood = sql.NullBool{Bool: isGood, Valid: true}
		return true
	}
	return false
}

func (r *BuyRecord) QuXiaoFahuo() bool {
	if r.Status == modeltype.QuXiao {
		return true
	} else if r.Status == modeltype.PayCheckFail || r.Status == modeltype.WaitPayCheck {
		r.Status = modeltype.QuXiao
		return true
	} else if r.Status == modeltype.WaitFahuo {
		r.Status = modeltype.CheckQuXiao
		return true
	}
	return false
}

func (r *BuyRecord) QuXiaoPay() bool {
	if r.Status == modeltype.QuXiao {
		return true
	} else if r.Status == modeltype.PayCheckFail || r.Status == modeltype.WaitPayCheck {
		r.Status = modeltype.QuXiao
		return true
	}
	return false
}
