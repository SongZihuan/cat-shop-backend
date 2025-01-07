package modeltype

type VideoType int

const (
	XieYiVideo VideoType = 1
	WupinVideo VideoType = 2
)

var NameToVideoType = map[string]VideoType{
	"XieYi": XieYiVideo,
	"WuPin": WupinVideo,
}

var VideoTypeToName = map[VideoType]string{
	XieYiVideo: "XieYi",
	WupinVideo: "WuPin",
}

const VideoPathV1 = "/v1/video"
