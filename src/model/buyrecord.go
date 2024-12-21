package model

import (
	"database/sql"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
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

	ShopName     string         `gorm:"type:varchar(20);not null;"`
	ShopPhone    string         `gorm:"type:varchar(30);not null"`
	ShopWeChat   sql.NullString `gorm:"type:varchar(50);"`
	ShopEmail    sql.NullString `gorm:"type:varchar(50);"`
	ShopLocation string         `gorm:"type:varchar(200);not null;"`
}
