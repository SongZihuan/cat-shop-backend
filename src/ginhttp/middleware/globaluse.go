package middleware

import "github.com/gin-gonic/gin"

type middlewareUser interface {
	Use(middleware ...gin.HandlerFunc) gin.IRoutes
}

func BaseUser(mu middlewareUser) {
	mu.Use(DBReady(), MustFormData(), XTokenMiddleware())
}

func BaseApi(mu middlewareUser) {
	mu.Use(MustFormData(), MustAccept(), ReturnContentJson())
}

func GlobalUse(mu middlewareUser) {
	BaseUser(mu)
	BaseApi(mu)
}

func ResourceUse(mu middlewareUser) {
	BaseUser(mu)
}

func SecretUse(mu middlewareUser) {
	BaseUser(mu)
	BaseApi(mu)
	mu.Use(MustXTokenMiddleware())
}

func testUse(mu middlewareUser) {
	BaseUser(mu)
	BaseApi(mu)
	mu.Use(TestApiMiddleware())
}

func TestSecretUse(mu middlewareUser) {
	testUse(mu)
	mu.Use(MustXTokenMiddleware())
}

func TestGlobalUse(mu middlewareUser) {
	testUse(mu)
}
