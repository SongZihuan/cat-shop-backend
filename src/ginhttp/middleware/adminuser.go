package middleware

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/contextkey"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func getAdminUser(c *gin.Context, self *model.User) *model.User {
	type Query struct {
		UserID uint `form:"userid"`
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

	user, err := action.GetUserByID(query.UserID, self.IsRootAdmin())
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
			if user.HasPermission(self) {
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
