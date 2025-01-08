package admingetuserlst

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
)

type Query struct {
	Page     int `form:"page"`
	PageSize int `form:"pagesize"`
}

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

func NewUser(user model.User) User {
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
		TotalShouHuo: user.TotalDaohuo,
		TotalPingJia: user.TotalPingJia,
	}
}

type Data struct {
	List     []User `json:"list"`
	Total    int    `json:"total"`
	MaxCount int    `json:"maxpage"`
}

func NewData(u []model.User, maxcount int) Data {
	res := make([]User, 0, len(u))
	for _, v := range u {
		res = append(res, NewUser(v))
	}

	return Data{
		List:     res,
		Total:    len(res),
		MaxCount: maxcount,
	}
}

func NewJsonData(m []model.User, maxcount int) data.Data {
	return data.NewSuccessWithData(NewData(m, maxcount))
}
