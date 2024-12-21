package modeltype

type ImageType int

const (
	XieYiImage  ImageType = 1
	WuPinImage  ImageType = 2
	ConfigImage ImageType = 3
	AvatarImage ImageType = 4
)

var NameToImageType = map[string]ImageType{
	"XieYi":  XieYiImage,
	"WuPin":  WuPinImage,
	"Config": ConfigImage,
	"Avatar": AvatarImage,
}

var ImageTypeToName = map[ImageType]string{
	XieYiImage:  "XieYi",
	WuPinImage:  "WuPin",
	ConfigImage: "Config",
	AvatarImage: "Avatar",
}
