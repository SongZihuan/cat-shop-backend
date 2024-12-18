package middleware

import "github.com/gin-gonic/gin"

type middlewareUser interface {
	Use(middleware ...gin.HandlerFunc) gin.IRoutes
}

func GlobalUse(mu middlewareUser) {
	mu.Use(XTokenMiddleware())
}
