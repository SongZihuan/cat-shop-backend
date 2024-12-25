package adminupdateuseravtar

import "mime/multipart"

type Query struct {
	File *multipart.FileHeader `form:"file"`
}
