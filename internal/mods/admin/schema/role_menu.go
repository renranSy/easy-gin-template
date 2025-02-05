package schema

import (
	"easy-gin-template/pkg/db"
)

type RoleMenu struct {
	db.Model
	RoleID int
	MenuID int
}

func (RoleMenu) TableName() string {
	return "role_menu"
}

type RoleMenuForm struct {
	IDList string `json:"idList"`
}
