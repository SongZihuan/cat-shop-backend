package admingetuserinfo

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
)

type User struct {
	ID           uint                 `json:"id"`
	Status       modeltype.UserStatus `json:"status"`
	Type         modeltype.UserType   `json:"type"`
	Name         string               `json:"name"`
	Wechat       string               `json:"wechat"`
	Email        string               `json:"email"`
	Location     string               `json:"location"`
	Avatar       string               `json:"avatar"`
	Phone        string               `json:"phone"`
	TotalPrice   modeltype.Price      `json:"totalPrice"`
	TotalBuy     modeltype.Total      `json:"totalBuy"`
	TotalGood    modeltype.Total      `json:"totalGood"`
	TotalJian    modeltype.Total      `json:"totalJian"`
	TotalShouHuo modeltype.Total      `json:"totalShouHuo"`
	TotalPingJia modeltype.Total      `json:"totalPingJia"`
}

func NewData(user *model.User) User {
	return User{
		ID:           user.ID,
		Status:       user.Status,
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
