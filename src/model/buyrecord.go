package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type BuyStatus int

const (
	WaitPayCheck      BuyStatus = 1
	PayCheckFail      BuyStatus = 2
	WaitFahuo         BuyStatus = 3
	WaitShouHuo       BuyStatus = 4
	WaitPingJia       BuyStatus = 5
	YiPingJia         BuyStatus = 6
	TuiHuoCheck       BuyStatus = 7
	WaitTuiHuoFahuo   BuyStatus = 8
	WaitTuiHuoShouHuo BuyStatus = 9
	TuiHuoFail        BuyStatus = 10
	TuiHuo            BuyStatus = 11
	CheckQuXiao       BuyStatus = 12
	QuXiao            BuyStatus = 13
)

var StatusToName = map[BuyStatus]string{
	WaitPayCheck:      "已下单，支付代确认",
	PayCheckFail:      "已下单，支付失败",
	WaitFahuo:         "待发货",
	WaitShouHuo:       "已发货",
	WaitPingJia:       "已收货",
	YiPingJia:         "已评价",
	TuiHuoCheck:       "申请退货审核中",
	WaitTuiHuoFahuo:   "申请退货成功，等待寄回",
	WaitTuiHuoShouHuo: "申请退货成功，待收货",
	TuiHuoFail:        "申请退货失败",
	TuiHuo:            "已退货",
	CheckQuXiao:       "申请取消订单审核中",
	QuXiao:            "已取消订单",
}

type BuyRecord struct {
	gorm.Model
	Status           BuyStatus      `gorm:"type:uint;not null;"`
	UserID           uint           `gorm:"not null"`
	User             *User          `gorm:"foreignKey:UserID"`
	WuPinID          uint           `gorm:"not null"`
	WuPin            *WuPin         `gorm:"foreignKey:WuPinID"`
	ClassID          uint           `gorm:"not null"`
	Class            *Class         `gorm:"foreignKey:ClassID"`
	Num              Total          `gorm:"type:uint;not null"`
	Price            Price          `gorm:"type:uint;not null"`
	TotalPrice       Price          `gorm:"type:uint;not null"`
	XiaDanTime       time.Time      `gorm:"type:datetime;not null"`
	FuKuanTime       sql.NullTime   `gorm:"type:datetime;"`
	FaHuoTime        sql.NullTime   `gorm:"type:datetime;"`
	ShouHuoTime      sql.NullTime   `gorm:"type:datetime;"`
	PingJiaTime      sql.NullTime   `gorm:"type:datetime;"`
	DengJiTuiHuoTime sql.NullTime   `gorm:"type:datetime;"`
	QueRenTuiHuoTime sql.NullTime   `gorm:"type:datetime;"`
	TuiHuoTime       sql.NullTime   `gorm:"type:datetime;"`
	QuXiaoTime       sql.NullTime   `gorm:"type:datetime;"`
	FaHuoKuaiDi      sql.NullString `gorm:"type:varchar(20);"`
	FaHuoKuaiDiNum   sql.NullString `gorm:"type:varchar(50);"`
	TuiHuoKuaiDi     sql.NullString `gorm:"type:varchar(20);"`
	TuiHuoKuaiDiNum  sql.NullString `gorm:"type:varchar(50);"`
	IsGood           sql.NullBool   `gorm:"type:boolean;"`

	UserName     string         `gorm:"type:varchar(20);not null;"`
	UserPhone    string         `gorm:"type:varchar(30);not null"`
	UserWeChat   sql.NullString `gorm:"type:varchar(50);"`
	UserEmail    sql.NullString `gorm:"type:varchar(50);"`
	UserLocation string         `gorm:"type:varchar(200);not null;"`

	ShopName     string         `gorm:"type:varchar(20);not null;"`
	ShopPhone    string         `gorm:"type:varchar(30);not null"`
	ShopWeChat   sql.NullString `gorm:"type:varchar(50);"`
	ShopEmail    sql.NullString `gorm:"type:varchar(50);"`
	ShopLocation string         `gorm:"type:varchar(200);not null;"`
}
