package adminupdateconfigpic

import (
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"mime/multipart"
)

type Query struct {
	Key   modeltype.ConfigKeyType `form:"key"`
	Value *multipart.FileHeader   `form:"value"`
}
