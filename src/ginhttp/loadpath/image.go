package loadpath

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/ginplus"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/resource/image"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"net/http"
)

var imagePath = ""

func LoadImagePath(engine *ginplus.Router) {
	cfg := config.Config().Yaml.Http

	path, ok := engine.FindURLByHandler(image.Handler, http.MethodGet)
	if !ok {
		resourcePath = GetResourcePath()
		if resourcePath == "" {
			path = cfg.BasePath + cfg.ResourcePath + modeltype.ImagePathV1 + "/image"
		} else {
			path = resourcePath + modeltype.ImagePathV1 + "/image"
		}
	}

	imagePath = path
	engine.DebugMsg("[INFO] image path: %s", path)
}

func GetImagePath() string {
	return imagePath
}
