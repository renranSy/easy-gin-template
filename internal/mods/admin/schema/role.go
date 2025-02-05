package schema

import (
	"easy-gin-template/pkg/db"
)

type Role struct {
	db.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Sequence    int    `json:"sequence"`
	Status      int    `json:"status"`
}

func (Role) TableName() string {
	return "role"
}

type RoleForm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Sequence    int    `json:"sequence"`
	Status      int    `json:"status"`
}
