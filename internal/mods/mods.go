package mods

import (
	"easy-gin-template/internal/config"
	"easy-gin-template/internal/middleware"
	"easy-gin-template/internal/mods/admin"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
)

var Set = wire.NewSet(
	wire.Struct(new(Mods), "*"),
	admin.Set,
)

type Mods struct {
	System     *admin.System
	Middleware *middleware.Middleware
}

func (a Mods) SetupRouter() *gin.Engine {
	RouteEngin := gin.New()

	RouteEngin.Use(a.Middleware.GinLogger())
	RouteEngin.Use(gin.Recovery())

	RouteEngin.GET("/ping", ping)

	var Group = RouteEngin.Group(config.CFG.Server.Prefix)

	a.System.SetupRouter(Group)

	return RouteEngin
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
