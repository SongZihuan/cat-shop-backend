package middleware

import (
	"errors"
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/adminaction"
	error2 "github.com/SongZihuan/cat-shop-backend/src/database/action/error"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/rootaction"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/contextkey"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func getAdminUser(c *gin.Context, self *model.User) (*model.User, error) {
	type Query struct {
		UserID uint `form:"userId"`
	}

	var err error
	var query Query
	if c.Request.Method == http.MethodGet {
		err = c.ShouldBindQuery(&query)
	} else if c.Request.Method == http.MethodPost {
		err = c.ShouldBindWith(&query, binding.FormMultipart)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if query.UserID <= 0 {
		return nil, fmt.Errorf("nad user id")
	}

	var user *model.User
	if self.IsRootAdmin() {
		user, err = rootaction.RootAdminGetUserByID(query.UserID)
	} else {
		user, err = adminaction.AdminGetUserByID(query.UserID)
	}
	if errors.Is(err, error2.ErrNotFound) {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return user, err
}

func AdminHasUserPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		self, ok := c.Value(contextkey.UserKey).(*model.User)
		if !ok {
			c.JSON(http.StatusOK, data.NewSystemUnknownError("管理员未找到"))
			return
		}

		user, err := getAdminUser(c, self)
		if err == nil || errors.Is(err, error2.ErrNotFound) {
			if c.Request.Method == http.MethodGet || user.HasPermission(self) {
				c.Set(contextkey.AdminUserIDKey, user.ID)
				c.Set(contextkey.AdminUserKey, user)
			} else {
				c.JSON(http.StatusOK, data.NewClientAdminUserNoPermission())
				return
			}
		} else {
			c.JSON(http.StatusOK, data.NewSystemUnknownError("用户未找到"))
			return
		}

		c.Next()
	}
}
