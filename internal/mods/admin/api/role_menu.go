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

type RoleMenu struct {
	RoleMenuBIZ biz.RoleMenu
}

func (a *RoleMenu) Query(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.ResError(c, errors.BadRequest(fmt.Sprintf("请求参数错误[%s]", err.Error())))
		return
	}
	list := a.RoleMenuBIZ.Query(id)

	util.ResOK(c, gin.H{
		"list": list,
	}, "")
}

func (a *RoleMenu) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.ResError(c, errors.BadRequest(fmt.Sprintf("请求参数错误[%s]", err.Error())))
		return
	}
	item := new(schema.RoleMenuForm)
	if err = c.ShouldBindJSON(item); err != nil {
		util.ResError(c, errors.BadRequest(err.Error()))
		return
	}
	if err = a.RoleMenuBIZ.Update(id, item); err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c, nil, "保存成功")
}
