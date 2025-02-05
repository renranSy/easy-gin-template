package biz

import (
	"easy-gin-template/internal/mods/admin/dal"
	"easy-gin-template/internal/mods/admin/schema"
	"easy-gin-template/pkg/casbinx"
	"easy-gin-template/pkg/errors"
	"strconv"
	"strings"
)

type RoleMenu struct {
	RoleMenuDAL  *dal.RoleMenu
	MenuDAL      *dal.Menu
	MenuResource *dal.MenuResource
	CasbinX      *casbinx.CasbinX
}

func (a *RoleMenu) Query(id int) *[]int {
	roleMenuList := a.RoleMenuDAL.Query("role_id = ?", id)

	ids := make([]int, len(*roleMenuList))
	for i, v := range *roleMenuList {
		ids[i] = v.MenuID
	}

	return &ids
}

func (a *RoleMenu) Update(id int, item *schema.RoleMenuForm) error {
	// TODO 使用事务
	if err := a.RoleMenuDAL.DeleteBy("role_id = ?", id); err != nil {
		return errors.InternalServerError(err.Error())
	}

	menuIDList := strings.Split(item.IDList, ",")
	for _, v := range menuIDList {
		menuId, _ := strconv.Atoi(v)
		if err := a.RoleMenuDAL.Create(&schema.RoleMenu{
			RoleID: id,
			MenuID: menuId,
		}); err != nil {
			return errors.InternalServerError(err.Error())
		}
	}

	if err := a.CasbinX.ResetPolicy(); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}
