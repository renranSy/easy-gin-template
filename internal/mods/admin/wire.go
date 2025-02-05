package admin

import (
	"easy-gin-template/internal/mods/admin/api"
	"easy-gin-template/internal/mods/admin/biz"
	"easy-gin-template/internal/mods/admin/dal"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	wire.Struct(new(System), "*"),
	wire.Struct(new(api.Admin), "*"),
	wire.Struct(new(biz.Admin), "*"),
	wire.Struct(new(dal.Admin), "*"),
	wire.Struct(new(dal.AdminRole), "*"),
	wire.Struct(new(api.Menu), "*"),
	wire.Struct(new(biz.Menu), "*"),
	wire.Struct(new(dal.Menu), "*"),
	wire.Struct(new(api.Role), "*"),
	wire.Struct(new(biz.Role), "*"),
	wire.Struct(new(dal.Role), "*"),
	wire.Struct(new(dal.MenuResource), "*"),
	wire.Struct(new(api.RoleMenu), "*"),
	wire.Struct(new(biz.RoleMenu), "*"),
	wire.Struct(new(dal.RoleMenu), "*"),
)
