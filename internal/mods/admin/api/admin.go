package api

import (
	"easy-gin-template/internal/enum"
	"easy-gin-template/internal/mods/admin/biz"
	"easy-gin-template/internal/mods/admin/schema"
	"easy-gin-template/pkg/errors"
	"easy-gin-template/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Admin struct {
	AdminBIZ *biz.Admin
}

func (a *Admin) Login(c *gin.Context) {
	var p schema.LoginReq

	if err := c.ShouldBindJSON(&p); err != nil {
		util.ResError(c, errors.BadRequest(fmt.Sprintf("请求参数错误[%s]", err.Error())))
		return
	}

	token, user, err := a.AdminBIZ.Login(&p)

	adminInfo := a.AdminBIZ.Get(user)

	if err != nil {
		util.ResError(c, err)
		return
	}

	util.ResOK(c, schema.LoginRes{
		Token:    token,
		UserInfo: *adminInfo,
	}, "登录成功")
}

func (a *Admin) Get(c *gin.Context) {
	var (
		admin = c.MustGet(enum.UserInfoContext).(*schema.Admin)
	)

	adminInfo := a.AdminBIZ.Get(admin)
	util.ResOK(c, adminInfo, "ok")
}

func (a *Admin) Query(c *gin.Context) {
	data, err := a.AdminBIZ.Query()
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c, data, "")
}

func (a *Admin) Create(c *gin.Context) {
	form := new(schema.AdminAddForm)
	if err := c.ShouldBindJSON(form); err != nil {
		util.ResError(c, errors.BadRequest(err.Error()))
		return
	}
	if err := a.AdminBIZ.Create(form); err != nil {
		util.ResError(c, err)
		return
	}

	util.ResOK(c, nil, fmt.Sprintf("创建用户[%s]成功", form.Username))
}

func (a *Admin) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.ResError(c, errors.BadRequest(fmt.Sprintf("请求参数错误[%s]", err.Error())))
		return
	}
	item := new(schema.AdminEditForm)
	if err = c.ShouldBindJSON(item); err != nil {
		util.ResError(c, errors.BadRequest(err.Error()))
		return
	}
	if err = a.AdminBIZ.Update(id, item); err != nil {
		util.ResError(c, err)
		return
	}

	util.ResOK(c, nil, "更新成功")
}

func (a *Admin) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.ResError(c, errors.BadRequest(fmt.Sprintf("请求参数错误[%s]", err.Error())))
		return
	}
	if err := a.AdminBIZ.Delete(id); err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c, nil, "删除成功")
}
