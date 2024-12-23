package getuserinfo

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
)

type User struct {
	Type         modeltype.UserType `json:"type"`
	Name         string             `json:"name"`
	Wechat       string             `json:"wechat"`
	Email        string             `json:"email"`
	Location     string             `json:"location"`
	Avatar       string             `json:"avatar"`
	Phone        string             `json:"phone"`
	TotalPrice   modeltype.Total    `json:"totalPrice"`
	TotalBuy     modeltype.Total    `json:"totalBuy"`
	TotalGood    modeltype.Total    `json:"totalGood"`
	TotalJian    modeltype.Total    `json:"totalJian"`
	TotalShouHuo modeltype.Total    `json:"totalShouHuo"`
	TotalPingJia modeltype.Total    `json:"totalPingJia"`
}

func NewData(user *model.User) User {
	return User{
		Type:         user.Type,
		Name:         user.Name,
		Wechat:       utils.GetSQLNullString(user.WeChat),
		Email:        utils.GetSQLNullString(user.Email),
		Location:     utils.GetSQLNullString(user.Location),
		Avatar:       utils.GetSQLNullString(user.Avatar),
		Phone:        user.Phone,
		TotalPrice:   user.TotalPrice,
		TotalBuy:     user.TotalBuy,
		TotalGood:    user.TotalGood,
		TotalJian:    user.TotalJian,
		TotalShouHuo: user.TotalShouHuo,
		TotalPingJia: user.TotalPingJia,
	}
}

func NewJsonData(user *model.User) data.Data {
	return data.NewSuccessWithData(NewData(user))
}
