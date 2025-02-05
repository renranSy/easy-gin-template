package schema

import (
	"easy-gin-template/pkg/db"
)

type Admin struct {
	db.Model
	Username   string `json:"username"`
	Password   string `json:"-"`
	Email      string `json:"email"`
	Status     int8   `json:"status" gorm:"default:1"`
	RoleIdList []int  `json:"roleIdList" gorm:"-"`
}

func (Admin) TableName() string {
	return "admin"
}

type LoginReq struct {
	Username string `form:"username" bind:"required"`
	Password string `form:"password" bind:"required"`
}

type LoginRes struct {
	UserInfo AdminInfo `json:"userInfo"`
	Token    string    `json:"token"`
}

type AdminInfo struct {
	Username string   `json:"username"`
	Menus    []string `json:"menus"`
	Buttons  []string `json:"buttons"`
}

type AdminAddForm struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RoleIDList []int  `json:"roleIdList"`
	Email      string `json:"email"`
}

type AdminEditForm struct {
	Username   string `json:"username" binding:"required"`
	Email      string `json:"email"`
	RoleIDList []int  `json:"roleIdList"`
	Status     int8   `json:"status"`
}

func (a Admin) FillTo(adminInfo *AdminInfo) {
	adminInfo.Username = a.Username
}
