package db

import (
	"easy-gin-template/internal/config"
	"easy-gin-template/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var DB *gorm.DB

func init() {
	DB = initDB()
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.CFG.Mysql.DSN), &gorm.Config{
		Logger: zapgorm2.New(log.Logger),
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
