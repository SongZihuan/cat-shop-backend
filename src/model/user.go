package model

import (
	"database/sql"
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/config"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Status       modeltype.UserStatus `gorm:"type:uint;not null"`
	Type         modeltype.UserType   `gorm:"type:uint;not null"`
	Name         sql.NullString       `gorm:"type:varchar(20);"`
	Phone        string               `gorm:"type:varchar(30);not null"`
	WeChat       sql.NullString       `gorm:"type:varchar(50);"`
	Email        sql.NullString       `gorm:"type:varchar(50);"`
	Location     sql.NullString       `gorm:"type:varchar(200);"`
	TotalPrice   modeltype.Total      `gorm:"type:uint;not null"`
	TotalBuy     modeltype.Total      `gorm:"type:uint;not null"`
	TotalGood    modeltype.Total      `gorm:"type:uint;not null"`
	TotalJian    modeltype.Total      `gorm:"type:uint;not null"`
	TotalShouHuo modeltype.Total      `gorm:"type:uint;not null"`
	PasswordHash string               `gorm:"type:char(64);not null"`
}

func NewUser(phone string, password string) *User {
	return &User{
		Status:       modeltype.NormalUserStatus,
		Type:         modeltype.NormalUserType,
		Phone:        phone,
		PasswordHash: getPasswordHash(password),
	}
}

func (u *User) GetName() string {
	if u.Name.Valid && u.Name.String != "" {
		return u.Name.String + " - " + u.Name.String
	}

	return u.Phone
}

func (u *User) CanLogin() bool {
	return u.Status == modeltype.NormalUserStatus
}

func (u *User) PasswordCheck(password string) bool {
	return u.PasswordHash == getPasswordHash(password)
}

func (u *User) SetNewPassword(password string) bool {
	newPassword := getPasswordHash(password)
	if newPassword == u.PasswordHash {
		return false
	}

	u.PasswordHash = newPassword
	return true
}

func getPasswordHash(password string) string {
	if !config.IsReady() {
		panic("config is not ready")
	}

	ps := fmt.Sprintf("%s:%s", config.Config().Yaml.Password.Backend, password)
	return utils.SHA256([]byte(ps))
}
