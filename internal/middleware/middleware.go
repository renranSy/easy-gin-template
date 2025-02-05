package middleware

import (
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Middleware struct {
	Log      *zap.Logger
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

func InitMiddleware(log *zap.Logger, db *gorm.DB, enforcer *casbin.Enforcer) *Middleware {
	return &Middleware{
		Log:      log,
		DB:       db,
		Enforcer: enforcer,
	}
}
