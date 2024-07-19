package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64 `gorm:"primaryKey"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (User) TableName() string {
	return "user"
}
