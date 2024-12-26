package model

import (
	"database/sql"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
	"time"
)

type BuyRecordM struct {
	gorm.Model
	Status             modeltype.BuyStatus `gorm:"type:uint;not null;"`
	UserID             uint                `gorm:"not null"`
	WuPinID            uint                `gorm:"not null"`
	ClassID            uint                `gorm:"not null"`
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
}

func (*BuyRecordM) TableName() string {
	return "buyrecord"
}

func init() {
	if !modelTest[BuyRecord, BuyRecordM]() {
		panic("database error")
	}
}
