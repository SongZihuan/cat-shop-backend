package addbag

type Query struct {
	WuPinID uint `form:"wupinId"`
	Num     int  `form:"num"`
}
