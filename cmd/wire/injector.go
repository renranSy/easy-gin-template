package wire

import (
	"easy-gin-template/internal/mods"
	"easy-gin-template/pkg/casbinx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Injector struct {
	Mods    *mods.Mods
	DB      *gorm.DB
	Log     *zap.Logger
	Casbinx *casbinx.CasbinX
}
