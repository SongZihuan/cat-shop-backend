package modeltype

type ImageType int

const (
	XieYiImage  ImageType = 1
	WupinImage  ImageType = 2
	ConfigImage ImageType = 3
	AvatarImage ImageType = 4
)

var NameToImageType = map[string]ImageType{
	"XieYi":  XieYiImage,
	"WuPin":  WupinImage,
	"Config": ConfigImage,
	"Avatar": AvatarImage,
}

var ImageTypeToName = map[ImageType]string{
	XieYiImage:  "XieYi",
	WupinImage:  "WuPin",
	ConfigImage: "Config",
	AvatarImage: "Avatar",
}

var ImageAlt = map[ImageType]string{
	XieYiImage:  "协议附图",
	WupinImage:  "商品图",
	ConfigImage: "",
	AvatarImage: "头像",
}

const ImagePathV1 = "/v1/image"
