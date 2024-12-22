package model

import (
	"gorm.io/gorm"
	"time"
)

type Msg struct {
	gorm.Model
	UserID uint      `gorm:"not null"`
	User   *User     `gorm:"foreignKey:UserID"`
	Msg    string    `gorm:"type:varchar(200);not null"`
	Time   time.Time `gorm:"type:datetime;not null"`
}

func NewMsg(userID uint, msg string) *Msg {
	return &Msg{
		UserID: userID,
		Msg:    msg,
		Time:   time.Now(),
	}
}
