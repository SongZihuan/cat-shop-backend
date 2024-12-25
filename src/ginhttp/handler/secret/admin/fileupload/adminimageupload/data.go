package adminimageupload

import "mime/multipart"

type FileType string

const (
	FileTypeXieyi FileType = "admin-xieyi"
	FileTypeWupin FileType = "admin-wupin"
)

type Query struct {
	File *multipart.FileHeader `form:"file"`
	Type FileType              `form:"type"`
}

type Error struct {
	Errno   int    `json:"errno"`
	Message string `json:"message"`
}

func NewError(msg ...string) *Error {
	if len(msg) == 0 {
		return &Error{
			Errno:   1,
			Message: "image upload fail",
		}
	} else if len(msg) == 1 {
		return &Error{
			Errno:   1,
			Message: msg[0],
		}
	} else {
		panic("too many arguments")
	}
}

type SuccessData struct {
	Url  string `json:"url"`
	Alt  string `json:"alt,omitempty"`
	Href string `json:"href,omitempty"`
}

type Success struct {
	Errno int         `json:"errno"`
	Data  SuccessData `json:"data"`
}

func NewSuccess(url string, alt string) *Success {
	return &Success{
		Errno: 0,
		Data: SuccessData{
			Url: url,
			Alt: alt,
		},
	}
}
