package router

import (
	"easy-gin-template/internal/config"
	"easy-gin-template/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

var group *gin.RouterGroup

func SetupRouter() *gin.Engine {
	RouteEngin := gin.New()

	RouteEngin.Use(middleware.GinLogger())
	RouteEngin.Use(gin.Recovery())

	RouteEngin.GET("/ping", ping)

	group = RouteEngin.Group(config.CFG.Server.Prefix)

	setupUserRouter()

	return RouteEngin
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
