package admin

import (
	"easy-gin-template/internal/middleware"
	"easy-gin-template/internal/mods/admin/api"
	"github.com/gin-gonic/gin"
)

type System struct {
	UserAPI     *api.Admin
	MenuAPI     *api.Menu
	RoleAPI     *api.Role
	RoleMenuAPI *api.RoleMenu
	Middleware  *middleware.Middleware
}

func (a System) SetupRouter(group *gin.RouterGroup) {
	group.POST("/admin/login", a.UserAPI.Login)
	group.GET("/admin/get", a.Middleware.Auth(), a.UserAPI.Get)
	group.GET("/admin", a.Middleware.Auth(), a.Middleware.Casbin(), a.UserAPI.Query)
	group.POST("/admin", a.Middleware.Auth(), a.Middleware.Casbin(), a.UserAPI.Create)
	group.PUT("/admin/:id", a.Middleware.Auth(), a.Middleware.Casbin(), a.UserAPI.Update)
	group.DELETE("/admin/:id", a.Middleware.Auth(), a.Middleware.Casbin(), a.UserAPI.Delete)

	group.GET("/menu/:id", a.Middleware.Auth(), a.Middleware.Casbin(), a.MenuAPI.Get)
	group.GET("/menu", a.Middleware.Auth(), a.Middleware.Casbin(), a.MenuAPI.Query)
	group.POST("/menu", a.Middleware.Auth(), a.Middleware.Casbin(), a.MenuAPI.Create)
	group.PUT("/menu/:id", a.Middleware.Auth(), a.Middleware.Casbin(), a.MenuAPI.Update)
	group.DELETE("/menu/:id", a.Middleware.Auth(), a.Middleware.Casbin(), a.MenuAPI.Delete)

	group.GET("/role", a.Middleware.Auth(), a.Middleware.Casbin(), a.RoleAPI.Query)
	group.POST("/role", a.Middleware.Auth(), a.Middleware.Casbin(), a.RoleAPI.Create)
	group.PUT("/role/:id", a.Middleware.Auth(), a.Middleware.Casbin(), a.RoleAPI.Update)
	group.DELETE("/role/:id", a.Middleware.Auth(), a.Middleware.Casbin(), a.RoleAPI.Delete)

	group.GET("/role/menu/:id", a.Middleware.Auth(), a.Middleware.Casbin(), a.RoleMenuAPI.Query)
	group.PUT("/role/menu/:id", a.Middleware.Auth(), a.Middleware.Casbin(), a.RoleMenuAPI.Update)
}
