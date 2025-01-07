package adminaddclass

type Query struct {
	Name string `form:"name"`
	Show bool   `form:"show"`
	Down bool   `form:"down"`
}
