//go:build wireinject
// +build wireinject

package wire

import (
	"context"
	"easy-gin-template/internal/middleware"
	"easy-gin-template/internal/mods"
	"easy-gin-template/pkg"
	"easy-gin-template/pkg/casbinx"
	"easy-gin-template/pkg/db"
	"easy-gin-template/pkg/log"
	"github.com/google/wire"
)

func SetupInjector(ctx context.Context) *Injector {
	wire.Build(
		log.InitLog,
		db.InitDB,
		casbinx.InitEnforcer,
		middleware.InitMiddleware,
		wire.NewSet(wire.Struct(new(Injector), "*")),
		mods.Set,
		pkg.Set,
	)
	return new(Injector)
}
