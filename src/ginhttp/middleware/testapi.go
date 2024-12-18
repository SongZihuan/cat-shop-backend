package middleware

import (
	"github.com/SuperH-0630/cat-shop-back/src/config"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/gin-gonic/gin"
)

func TestApiMiddleware() gin.HandlerFunc {
	if !config.IsReady() {
		panic("config is not ready")
	}

	if config.Config().Yaml.Http.TestApi {
		return func(c *gin.Context) {
			c.Next()
		}
	} else {
		return func(c *gin.Context) {
			c.JSON(200, data.NewData(data.GlobalErrorCodeNotTest, nil, "并非处于测试模式", "并非处于测试模式"))
		}
	}
}
