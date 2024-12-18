package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type UserType int

const (
	NormalUserType    UserType = 1
	AdminUserType     UserType = 2
	RootAdminUserType UserType = 3
)

type UserStatus int

const (
	NormalUserStatus UserStatus = 1
	FreezeUserStatus UserStatus = 2
	DeleteUserStatus UserStatus = 3
)

type User struct {
	gorm.Model
	Status       UserStatus     `gorm:"type:uint;not null"`
	Type         UserType       `gorm:"type:uint;not null"`
	Name         sql.NullString `gorm:"type:varchar(20);"`
	Phone        string         `gorm:"type:varchar(30);not null"`
	WeChat       sql.NullString `gorm:"type:varchar(50);"`
	Email        sql.NullString `gorm:"type:varchar(50);"`
	Location     sql.NullString `gorm:"type:varchar(200);"`
	TotalPrice   Total          `gorm:"type:uint;not null"`
	TotalBuy     Total          `gorm:"type:uint;not null"`
	TotalGood    Total          `gorm:"type:uint;not null"`
	TotalJian    Total          `gorm:"type:uint;not null"`
	TotalShouHuo Total          `gorm:"type:uint;not null"`
	PasswordHash string         `gorm:"type:char(64);not null"`
}

func (u *User) GetName() string {
	if u.Name.Valid && u.Name.String != "" {
		return u.Name.String + " - " + u.Name.String
	}

	return u.Phone
}

func (u *User) CanLogin() bool {
	return u.Status == NormalUserStatus
}
