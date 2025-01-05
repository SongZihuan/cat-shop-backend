package model

import (
	"database/sql"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

type UserM struct {
	gorm.Model
	Status       modeltype.UserStatus `gorm:"type:uint;not null"`
	Type         modeltype.UserType   `gorm:"type:uint;not null"`
	Name         string               `gorm:"type:varchar(20);not null"`
	Phone        string               `gorm:"type:varchar(30);not null"`
	WeChat       sql.NullString       `gorm:"type:varchar(50);"`
	Email        sql.NullString       `gorm:"type:varchar(50);"`
	Location     sql.NullString       `gorm:"type:varchar(200);"`
	Avatar       sql.NullString       `gorm:"type:varchar(200);"`
	TotalPrice   modeltype.Total      `gorm:"type:uint;not null"`
	TotalBuy     modeltype.Total      `gorm:"type:uint;not null"`
	TotalGood    modeltype.Total      `gorm:"type:uint;not null"`
	TotalJian    modeltype.Total      `gorm:"type:uint;not null"`
	TotalShouHuo modeltype.Total      `gorm:"type:uint;not null"`
	TotalPingJia modeltype.Total      `gorm:"type:uint;not null"`
	PasswordHash string               `gorm:"type:char(64);not null"`
}

func (*UserM) TableName() string {
	return "user"
}

func init() {
	if !modelTest[User, UserM]() {
		panic("database error")
	}
}
