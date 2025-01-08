package model

import (
	"database/sql"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
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
	QueRenTuiHuoTime   sql.NullTime        `gorm:"type:datetime;"`
	DengJiTuiHuoTime   sql.NullTime        `gorm:"type:datetime;"`
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

	WupinBuyTotal   modeltype.Total `gorm:"type:uint;not null"`
	WupinBuyDaoHuo  modeltype.Total `gorm:"type:uint;not null"`
	WupinBuyGood    modeltype.Total `gorm:"type:uint;not null"`
	WupinBuyPrice   modeltype.Price `gorm:"type:uint;not null"`
	WupinBuyPingJia modeltype.Total `gorm:"type:uint;not null"`
	WupinBuyJian    modeltype.Total `gorm:"type:uint;not null"`
	WupinHot        bool            `gorm:"type:boolean;not null"`

	WupinDown bool `gorm:"type:boolean;not null"` // 并非物品Lock信息
	ClassDown bool `gorm:"type:boolean;not null;"`
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

func (r *BuyRecord) IsWupinSale() bool {
	return r.isWupinDown() || r.isClassDown()
}

func (r *BuyRecord) IsWupinNotSale() bool {
	return !r.IsWupinSale()
}

func (r *BuyRecord) IsBuyRecordCanRepay() bool {
	return r.IsWupinSale() && r.Status == modeltype.PayCheckFail
}

func (r *BuyRecord) IsBuyRecordCanNotRepay() bool {
	return !r.IsBuyRecordCanRepay()
}

func (r *BuyRecord) IsBuyRecordCanPay() bool {
	return r.IsWupinSale() && r.Status == modeltype.WaitPay
}

func (r *BuyRecord) IsBuyRecordCanNotPay() bool {
	return !r.IsBuyRecordCanPay()
}

func (*BuyRecord) IsBuyRecordDown() bool {
	return false
}
