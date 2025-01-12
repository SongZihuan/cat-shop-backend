package adminadduser

type Query struct {
	Phone    string `form:"phone"`
	Password string `form:"password"`
}
