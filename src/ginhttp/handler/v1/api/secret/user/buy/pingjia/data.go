package pingjia

type Query struct {
	ID     uint `form:"id"`
	IsGood bool `form:"isgood"`
}
