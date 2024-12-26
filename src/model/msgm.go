package model

import (
	"gorm.io/gorm"
	"time"
)

type MsgM struct {
	gorm.Model
	UserID uint      `gorm:"not null"`
	Msg    string    `gorm:"type:varchar(200);not null"`
	Time   time.Time `gorm:"type:datetime;not null"`
}

func (*MsgM) TableName() string {
	return "msg"
}

func init() {
	if !modelTest[Msg, MsgM]() {
		panic("database error")
	}
}
