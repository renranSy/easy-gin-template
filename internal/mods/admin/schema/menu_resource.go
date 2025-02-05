package schema

import (
	"easy-gin-template/pkg/db"
)

type MenuResource struct {
	db.Model
	MenuID int    `json:"menuId"`
	Method string `json:"method"`
	Path   string `json:"path"`
}

func (MenuResource) TableName() string {
	return "menu_resource"
}
