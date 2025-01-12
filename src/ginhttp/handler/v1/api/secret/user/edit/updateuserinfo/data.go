package updateuserinfo

type Query struct {
	Name     string `form:"name"`
	Wechat   string `form:"wechat"`
	Email    string `form:"email"`
	Location string `form:"location"`
}
