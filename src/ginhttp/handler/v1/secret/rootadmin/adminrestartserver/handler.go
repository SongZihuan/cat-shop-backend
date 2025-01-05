package adminrestartserver

import (
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/config"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/contextkey"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/httpstop"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	CodePasswordError = -3
	CodeSecretError   = -4
)

const MinWaitSec = 10
const MaxWaitSec = 60

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

	waitsec := query.WaitSec
	if waitsec < MinWaitSec {
		waitsec = MinWaitSec
	} else if waitsec > MaxWaitSec {
		waitsec = MaxWaitSec
	}

	if !user.PasswordCheck(query.Password) {
		c.JSON(http.StatusOK, data.NewCustomError(CodePasswordError, "用户密码错误"))
		return
	}

	if !config.Config().Yaml.Http.CheckStopSecret(query.Secret) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeSecretError, "密钥错误"))
		return
	}

	_, err = utils.Restart("-wait", fmt.Sprintf("%d", waitsec))
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemUnknownError(err))
		return
	}
	httpstop.SetStop()

	c.JSON(200, NewSuccessJsonData(waitsec))
}
