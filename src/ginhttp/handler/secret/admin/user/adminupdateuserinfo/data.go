package adminupdateuserinfo

import "github.com/SuperH-0630/cat-shop-back/src/model/modeltype"

type Query struct {
	Name     string               `form:"name"`
	Wechat   string               `form:"wechat"`
	Email    string               `form:"email"`
	Location string               `form:"location"`
	Status   modeltype.UserStatus `form:"status"`
	Type     modeltype.UserType   `form:"type"`
}
