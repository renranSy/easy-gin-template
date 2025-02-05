package pkg

import (
	"easy-gin-template/pkg/casbinx"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	wire.Struct(new(casbinx.CasbinX), "*"),
)
