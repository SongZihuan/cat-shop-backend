package modeltype

type VideoType int

const (
	XieYiVideo VideoType = 1
	WuPinVideo VideoType = 2
)

var NameToVideoType = map[string]VideoType{
	"XieYi": XieYiVideo,
	"WuPin": WuPinVideo,
}

var VideoTypeToName = map[VideoType]string{
	XieYiVideo: "XieYi",
	WuPinVideo: "WuPin",
}
