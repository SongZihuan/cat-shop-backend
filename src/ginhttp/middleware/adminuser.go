package middleware

import (
	"errors"
	"github.com/SongZihuan/cat-shop-backend/src/database/action"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/contextkey"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func getAdminUser(c *gin.Context, self *model.User) *model.User {
	type Query struct {
		UserID uint `form:"userId"`
	}

	var err error
	var query Query
	if c.Request.Method == "GET" {
		err = c.ShouldBindQuery(&query)
	} else if c.Request.Method == "POST" {
		err = c.ShouldBindWith(&query, binding.FormMultipart)
	} else {
		return nil
	}
	if err != nil {
		return nil
	}

	if query.UserID <= 0 {
		return nil
	}

	var user *model.User
	if self.IsRootAdmin() {
		user, err = action.RootAdminGetUserByID(query.UserID)
	} else {
		user, err = action.AdminGetUserByID(query.UserID)
	}
	if errors.Is(err, action.ErrNotFound) {
		return nil
	} else if err != nil {
		return nil
	}

	return user
}

func AdminUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		self, ok := c.Value(contextkey.UserKey).(*model.User)
		if !ok {
			c.JSON(http.StatusOK, data.NewSystemUnknownError("admin self not found"))
			return
		}

		user := getAdminUser(c, self)
		if user != nil {
			if c.Request.Method == http.MethodGet || user.HasPermission(self) {
				c.Set(contextkey.AdminUserIDKey, user.ID)
				c.Set(contextkey.AdminUserKey, user)
			} else {
				c.JSON(http.StatusOK, data.NewClientAdminUserNoPermission())
				return
			}
		}

		c.Next()
	}
}
