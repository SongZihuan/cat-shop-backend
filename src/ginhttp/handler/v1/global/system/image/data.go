package image

type Query struct {
	Type string `form:"type"`
	Hash string `form:"hash"`
	Time string `form:"time"`
}
