package adminfahuochuangeuser

type Query struct {
	ID uint `form:"id"`

	UserName     string `form:"username"`
	UserPhone    string `form:"userphone"`
	UserLocation string `form:"userlocation"`
	UserWechat   string `form:"userwechat"`
	UserEmail    string `form:"useremail"`
	UserRemark   string `form:"userremark"`
}
