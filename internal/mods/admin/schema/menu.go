package schema

import (
	"easy-gin-template/pkg/db"
)

type Menu struct {
	db.Model
	Name          string          `json:"name"`
	Code          string          `json:"code"`
	Description   string          `json:"description"`
	Sequence      int             `json:"sequence"`
	Type          string          `json:"type"`
	Path          string          `json:"path"`
	Status        int             `json:"status"`
	ParentName    string          `json:"parentName" gorm:"-"`
	ParentID      int             `json:"parentId"`
	Children      []*Menu         `json:"children" gorm:"-"`
	MenuResources *[]MenuResource `json:"menuResources" gorm:"-"`
}

func (Menu) TableName() string {
	return "menu"
}

type MenuForm struct {
	Name          string `json:"name" binding:"required"`
	Code          string `json:"code" binding:"required"`
	Description   string `json:"description"`
	Sequence      int    `json:"sequence"`
	Type          string `json:"type"`
	Path          string `json:"path"`
	Status        int    `json:"status"`
	ParentID      int    `json:"parentId"`
	MenuResources []struct {
		ID     int    `json:"id"`
		Method string `json:"method" binding:"required"`
		Path   string `json:"path" binding:"required"`
	} `json:"menuResources" binding:"required"`
}

func (a MenuForm) FillTo(menu *Menu) {
	menu.Name = a.Name
	menu.Code = a.Code
	menu.Description = a.Description
	menu.Sequence = a.Sequence
	menu.Type = a.Type
	menu.Path = a.Path
	menu.Status = a.Status
	menu.ParentID = a.ParentID
}
