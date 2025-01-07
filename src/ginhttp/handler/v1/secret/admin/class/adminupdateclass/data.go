package adminupdateclass

type Query struct {
	ID   uint   `form:"id"`
	Name string `form:"name"`
	Show bool   `form:"show"`
	Down bool   `form:"down"`
}
