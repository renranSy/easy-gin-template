package db

import (
	"easy-gin-template/internal/config"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
	"time"
)

type Model struct {
	ID        int            `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func InitDB(log *zap.Logger) *gorm.DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		config.CFG.Mysql.User,
		config.CFG.Mysql.Password,
		config.CFG.Mysql.Addr,
		config.CFG.Mysql.Database,
		config.CFG.Mysql.Other)), &gorm.Config{
		Logger: zapgorm2.New(log),
	})
	if err != nil {
		panic(err)
	}

	debug := config.CFG.Mysql.Debug
	if debug {
		db = db.Debug()
	}
	return db
}
