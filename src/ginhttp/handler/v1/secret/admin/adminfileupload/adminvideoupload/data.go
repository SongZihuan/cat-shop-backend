package adminvideoupload

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
	Url    string `json:"url"`
	Poster string `json:"poster,omitempty"`
}

type Success struct {
	Errno int         `json:"errno"`
	Data  SuccessData `json:"data"`
}

func NewSuccess(url string, poster string) *Success {
	return &Success{
		Errno: 0,
		Data: SuccessData{
			Url:    url,
			Poster: poster,
		},
	}
}
