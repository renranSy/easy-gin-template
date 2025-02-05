package api

import (
	"easy-gin-template/internal/mods/admin/biz"
	"easy-gin-template/internal/mods/admin/schema"
	"easy-gin-template/pkg/errors"
	"easy-gin-template/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Role struct {
	RoleBIZ *biz.Role
}

func (a *Role) Query(c *gin.Context) {
	data := a.RoleBIZ.Query()
	util.ResOK(c, data, "")
}

func (a *Role) Create(c *gin.Context) {
	form := new(schema.RoleForm)
	if err := c.ShouldBindJSON(form); err != nil {
		util.ResError(c, errors.BadRequest(fmt.Sprintf("请求参数错误[%s]", err.Error())))
		return
	}
	if err := a.RoleBIZ.Create(form); err != nil {
		util.ResError(c, err)
		return
	}

	util.ResOK(c, nil, fmt.Sprintf("创建角色[%s]成功", form.Name))
}

func (a *Role) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.ResError(c, errors.BadRequest(fmt.Sprintf("请求参数错误[%s]", err.Error())))
		return
	}
	item := new(schema.RoleForm)
	if err = c.ShouldBindJSON(item); err != nil {
		util.ResError(c, errors.BadRequest(err.Error()))
		return
	}
	if err = a.RoleBIZ.Update(id, item); err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c, nil, "更新成功")
}

func (a *Role) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.ResError(c, errors.BadRequest(fmt.Sprintf("请求参数错误[%s]", err.Error())))
		return
	}

	if err = a.RoleBIZ.Delete(id); err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c, nil, "删除成功")
}
