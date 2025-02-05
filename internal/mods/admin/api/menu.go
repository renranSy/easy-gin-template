package api

import (
	"easy-gin-template/internal/mods/admin/biz"
	"easy-gin-template/internal/mods/admin/schema"
	"easy-gin-template/pkg/errors"
	"easy-gin-template/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
)

type Menu struct {
	MenuBIZ *biz.Menu
}

func (a *Menu) Query(c *gin.Context) {
	data := a.MenuBIZ.Query()
	util.ResOK(c, data, "")
}

func (a *Menu) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := a.MenuBIZ.Get(id)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c, data, "")
}

func (a *Menu) Create(c *gin.Context) {
	form := new(schema.MenuForm)
	if err := c.ShouldBindBodyWith(form, binding.JSON); err != nil {
		fmt.Println(err)
		util.ResError(c, errors.BadRequest(err.Error()))
		return
	}
	if err := a.MenuBIZ.Create(form); err != nil {
		util.ResError(c, err)
	}

	util.ResOK(c, nil, fmt.Sprintf("创建菜单【%s】成功", form.Name))
}

func (a *Menu) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.ResError(c, errors.BadRequest("请求参数错误"))
		return
	}
	item := new(schema.MenuForm)
	if err := c.ShouldBindJSON(item); err != nil {
		util.ResError(c, errors.BadRequest(err.Error()))
		return
	}
	if err = a.MenuBIZ.Update(id, item); err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c, nil, "更新成功")
}

func (a *Menu) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := a.MenuBIZ.Delete(id); err != nil {
		util.ResError(c, err)
		return
	}

	util.ResOK(c, nil, "删除成功")
}
