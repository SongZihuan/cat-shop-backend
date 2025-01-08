package adminfahuochangeshop

type Query struct {
	ID uint `form:"id"`

	ShopName     string `form:"username"`
	ShopPhone    string `form:"userphone"`
	ShopLocation string `form:"userlocation"`
	ShopWechat   string `form:"userwechat"`
	ShopEmail    string `form:"useremail"`
	ShopRemark   string `form:"userremark"`
}
