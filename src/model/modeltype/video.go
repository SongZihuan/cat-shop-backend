package modeltype

type VideoType int

const (
	XieYiVide VideoType = 1
	WuPinVide VideoType = 2
)

var NameToVideoType = map[string]VideoType{
	"XieYi": XieYiVide,
	"WuPin": WuPinVide,
}

var VideoTypeToName = map[VideoType]string{
	XieYiVide: "XieYi",
	WuPinVide: "WuPin",
}
