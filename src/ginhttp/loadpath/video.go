package loadpath

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/ginplus"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/resource/video"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"net/http"
)

var videosPath = ""

func LoadVideoPath(engine *ginplus.Router) {
	cfg := config.Config().Yaml.Http

	path, ok := engine.FindURLByHandler(video.Handler, http.MethodGet)
	if !ok {
		resourcePath = GetResourcePath()
		if resourcePath == "" {
			path = cfg.BasePath + cfg.ResourcePath + modeltype.VideoPathV1 + "/video"
		} else {
			path = resourcePath + modeltype.VideoPathV1 + "/video"
		}
	}

	videosPath = path
	engine.DebugMsg("[INFO] video path: %s", path)
}

func GetVideoPath() string {
	return videosPath
}
