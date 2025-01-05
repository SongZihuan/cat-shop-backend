package model

import (
	"database/sql"
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Status       modeltype.UserStatus `gorm:"type:uint;not null"`
	Type         modeltype.UserType   `gorm:"type:uint;not null"`
	Name         string               `gorm:"type:varchar(20);not null"`
	Phone        string               `gorm:"type:varchar(30);not null"`
	WeChat       sql.NullString       `gorm:"type:varchar(50);"`
	Email        sql.NullString       `gorm:"type:varchar(50);"`
	Location     sql.NullString       `gorm:"type:varchar(200);"`
	Avatar       sql.NullString       `gorm:"type:varchar(200);"`
	TotalPrice   modeltype.Price      `gorm:"type:uint;not null"`
	TotalBuy     modeltype.Total      `gorm:"type:uint;not null"`
	TotalGood    modeltype.Total      `gorm:"type:uint;not null"`
	TotalJian    modeltype.Total      `gorm:"type:uint;not null"`
	TotalShouHuo modeltype.Total      `gorm:"type:uint;not null"`
	TotalPingJia modeltype.Total      `gorm:"type:uint;not null"`
	PasswordHash string               `gorm:"type:char(64);not null"`
}

func (*User) TableName() string {
	return "user"
}

func NewUser(phone string, password string) *User {
	return &User{
		Status:       modeltype.NormalUserStatus,
		Type:         modeltype.NormalUserType,
		Phone:        phone,
		Name:         "新用户",
		PasswordHash: getPasswordHash(password),
	}
}

func (u *User) UpdateInfo(name string, wechat string, email string, location string) {
	if len(name) == 0 {
		name = "新用户"
	}

	u.Name = name
	u.WeChat = sql.NullString{String: wechat, Valid: len(wechat) != 0}
	u.Email = sql.NullString{String: email, Valid: len(email) != 0}
	u.Location = sql.NullString{String: location, Valid: len(location) != 0}
}

func (u *User) UpdateType(tp modeltype.UserType) bool {
	if u.Type == modeltype.RootAdminUserType {
		if tp != modeltype.RootAdminUserType {
			return false
		}

		return true
	} else if tp == modeltype.RootAdminUserType { // 不允许设置更多的根管理员
		return false
	} else {
		u.Type = tp
		return true
	}
}

func (u *User) UpdateStatus(st modeltype.UserStatus) bool {
	if st == modeltype.DeleteUserStatus {
		if u.Status == modeltype.DeleteUserStatus {
			return true
		} else {
			return false
		}
	} else {
		u.Status = st
		return true
	}
}

func (u *User) UpdateStatusWithRoot(st modeltype.UserStatus) bool {
	// 只有root可以设置delete
	u.Status = st
	return true
}

func (u *User) UpdatePhone(phone string) {
	u.Phone = phone
}

func (u *User) GetLongName() string {
	return u.Name + " - " + u.Name
}

func (u *User) IsRootAdmin() bool {
	u.Status = modeltype.NormalUserStatus // Root必须是Normal
	return u.Type == modeltype.RootAdminUserType
}

func (u *User) IsNormalAdmin() bool {
	return u.Status == modeltype.NormalUserStatus && u.Type == modeltype.AdminUserType
}

func (u *User) IsAdmin() bool {
	return u.IsRootAdmin() || u.IsNormalAdmin()
}

func (u *User) IsNormalUser() bool {
	return !u.IsAdmin()
}

func (u *User) CanLogin() bool {
	if u.IsRootAdmin() {
		return true
	}

	return u.Status == modeltype.NormalUserStatus
}

func (u *User) PasswordCheck(password string) bool {
	return u.PasswordHash == getPasswordHash(password)
}

func (u *User) UpdatePasswordWithCheck(oldPassword string, newPassword string) bool {
	if !u.PasswordCheck(oldPassword) {
		return false
	}
	u.UpdatePassword(newPassword)
	return true
}

func (u *User) UpdatePassword(newPassword string) {
	u.PasswordHash = getPasswordHash(newPassword)
}

func (u *User) UpdateAvatar(avatarUrl string) bool {
	u.Avatar = sql.NullString{String: avatarUrl, Valid: len(avatarUrl) != 0}
	return true
}

func (u *User) SetNewPassword(password string) bool {
	newPassword := getPasswordHash(password)
	if newPassword == u.PasswordHash {
		return false
	}

	u.PasswordHash = newPassword
	return true
}

func (u *User) IsDeleteUser() bool {
	return u.Status == modeltype.DeleteUserStatus
}

func (u *User) HasPermission(admin *User) bool {
	if admin.Type == modeltype.RootAdminUserType {
		return true
	} else if u.IsDeleteUser() {
		return false
	} else if admin.Type == modeltype.AdminUserType {
		return u.Type == modeltype.NormalUserType
	} else if admin.Type == modeltype.NormalUserType {
		return false
	} else {
		panic("error user")
	}
}

func getPasswordHash(password string) string {
	if !config.IsReady() {
		panic("config is not ready")
	}

	ps := fmt.Sprintf("%s:%s", config.Config().Yaml.Password.Backend, password)
	return utils.SHA256([]byte(ps))
}

func (u *User) BuyNow(r *BuyRecord) bool {
	if r.UserID != u.ID || r.User == nil || r.User.ID != u.ID {
		return false
	}

	u.TotalPrice += r.TotalPrice
	u.TotalBuy += 1
	u.TotalJian += r.Num
	return true
}

func (u *User) BackPayNow(r *BuyRecord) bool {
	if r.UserID != u.ID || r.User == nil || r.User.ID != u.ID {
		return false
	}

	u.TotalPrice -= r.TotalPrice
	u.TotalBuy -= 1
	u.TotalJian -= r.Num

	if u.TotalPrice <= 0 {
		u.TotalPrice = 0
	}

	if u.TotalBuy <= 0 {
		u.TotalBuy = 0
	}

	if u.TotalJian <= 0 {
		u.TotalJian = 0
	}

	return true
}

func (u *User) Daohuo(r *BuyRecord) bool {
	if r.WupinID != u.ID || r.Wupin == nil || r.Wupin.ID != u.ID {
		return true
	}
	u.TotalShouHuo += 1
	return false
}

func (u *User) PingJia(r *BuyRecord, isGood bool) bool {
	if r.WupinID != u.ID || r.Wupin == nil || r.Wupin.ID != u.ID {
		return true
	}

	u.TotalPingJia += 1

	if isGood {
		u.TotalGood += 1
	}

	return false
}
