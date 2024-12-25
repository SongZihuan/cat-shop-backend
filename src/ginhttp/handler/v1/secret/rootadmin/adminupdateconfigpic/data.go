package adminupdateconfigpic

import (
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"mime/multipart"
)

type Query struct {
	Key   modeltype.ConfigKeyType `form:"key"`
	Value *multipart.FileHeader   `form:"value"`
}
