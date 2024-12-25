package testpay

type Query struct {
	ID       uint `form:"id"`
	FailRate int  `form:"failrate"`
}
