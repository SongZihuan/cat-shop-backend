package admintuihuoaccept

type Query struct {
	ID     uint `form:"id"`
	Accept bool `form:"accept"`
}
