package resourcepath

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/ginplus"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/resource/image"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/resource/video"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"net/http"
)

var imagePath = ""
var videosPath = ""

func LoadImagePath(engine *ginplus.Router) {
	cfg := config.Config().Yaml.Http

	path, ok := engine.FindURL(image.Handler, http.MethodGet)
	if !ok {
		path = cfg.BasePath + cfg.ResourcePath + modeltype.ImagePathV1
	}

	imagePath = path
	engine.DebugMsg("[INFO] image path: %s", path)
}

func LoadVideoPath(engine *ginplus.Router) {
	cfg := config.Config().Yaml.Http

	path, ok := engine.FindURL(video.Handler, http.MethodGet)
	if !ok {
		path = cfg.BasePath + cfg.ResourcePath + modeltype.VideoPathV1
	}

	videosPath = path
	engine.DebugMsg("[INFO] video path: %s", path)
}

func GetImagePath() string {
	return imagePath
}

func GetVideoPath() string {
	return videosPath
}
