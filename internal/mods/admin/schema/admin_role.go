package schema

import (
	"easy-gin-template/pkg/db"
)

type AdminRole struct {
	db.Model
	AdminId int
	RoleId  int
}

func (AdminRole) TableName() string {
	return "admin_role"
}
