package adminstopserver

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/contextkey"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/httpstop"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	CodePasswordError = -3
	CodeSecretError   = -4
)

func Handler(c *gin.Context) {
	user, ok := c.Value(contextkey.UserKey).(*model.User)
	if !ok {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("用户未找到"))
		return
	}

	query := Query{}
	err := c.ShouldBindWith(&query, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

	if !user.PasswordCheck(query.Password) {
		c.JSON(http.StatusOK, data.NewCustomError(CodePasswordError, "用户密码错误"))
		return
	}

	if !config.Config().Yaml.Http.CheckStopSecret(query.Secret) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeSecretError, "密钥错误"))
		return
	}

	httpstop.SetStop()
	c.JSON(200, data.NewSuccess("退出信号已发出"))
}
