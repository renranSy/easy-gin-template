package biz

import (
	"easy-gin-template/internal/mods/admin/dal"
	"easy-gin-template/internal/mods/admin/schema"
	"easy-gin-template/pkg/db"
	"easy-gin-template/pkg/errors"
	"easy-gin-template/pkg/jwt"
	"fmt"
	"strconv"
	"time"
)

type Admin struct {
	AdminDAL     *dal.Admin
	AdminRoleDAL *dal.AdminRole
	RoleMenuDAL  *dal.RoleMenu
	MenuDAL      *dal.Menu
}

func (a *Admin) Login(form *schema.LoginReq) (string, *schema.Admin, error) {
	var user *schema.Admin
	user, err := a.AdminDAL.GetByUsernameAndPassword(form.Username, form.Password)
	if err != nil {
		return "", nil, errors.BadRequest("用户名或密码错误")
	}

	token, err := jwt.Jwt.GenToken(strconv.Itoa(user.ID), time.Now().Add(2*time.Hour))

	if err != nil {
		return "", nil, errors.InternalServerError("生成token失败")
	}
	return token, user, nil
}

func (a *Admin) Get(admin *schema.Admin) *schema.AdminInfo {
	roles := a.AdminRoleDAL.Query("admin_id = ?", admin.ID)
	menus := make([]string, 0)
	buttons := make([]string, 0)
	for _, role := range *roles {
		roleMenuList := a.RoleMenuDAL.Query("role_id = ?", role.RoleId)
		for _, roleMenu := range *roleMenuList {
			menu, _ := a.MenuDAL.Get("id = ?", roleMenu.MenuID)

			if menu.Type == "page" {
				menus = append(menus, menu.Code)
			} else {
				parentMenu, _ := a.MenuDAL.Get("id = ?", menu.ParentID)
				buttons = append(buttons, fmt.Sprintf("%s.%s", parentMenu.Code, menu.Code))
			}
		}
	}

	adminInfo := new(schema.AdminInfo)
	admin.FillTo(adminInfo)
	adminInfo.Buttons = buttons
	adminInfo.Menus = menus
	return adminInfo
}

func (a *Admin) Query() (*[]schema.Admin, error) {
	list, err := a.AdminDAL.Query()
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	for i, v := range *list {
		roleIds := a.AdminRoleDAL.Query("admin_id = ?", v.ID)
		ids := make([]int, 0)
		for _, roleId := range *roleIds {
			ids = append(ids, roleId.RoleId)
		}
		(*list)[i].RoleIdList = ids
	}
	return list, nil
}

func (a *Admin) Create(item *schema.AdminAddForm) error {
	// TODO 使用事务
	exists, err := a.AdminDAL.Exists("username = ?", item.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.BadRequest("用户名已存在")
	}

	admin := &schema.Admin{
		Username: item.Username,
		Password: item.Password,
		Email:    item.Email,
	}
	if err = a.AdminDAL.Create(admin); err != nil {
		return errors.InternalServerError(err.Error())
	}

	for _, v := range item.RoleIDList {
		if err = a.AdminRoleDAL.Create(&schema.AdminRole{
			AdminId: admin.ID,
			RoleId:  v,
		}); err != nil {
			return errors.InternalServerError(err.Error())
		}
	}
	return nil
}

func (a *Admin) Update(id int, item *schema.AdminEditForm) error {
	// TODO 使用事务
	exists, err := a.AdminDAL.Exists("id = ?", id)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	if !exists {
		return errors.NotFound("当前用户不存在")
	}
	exists, err = a.AdminDAL.Exists("username = ? AND id != ?", item.Username, id)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	if exists {
		return errors.BadRequest("用户名已存在")
	}

	if err = a.AdminDAL.Update(&schema.Admin{
		Model: db.Model{
			ID: id,
		},
		Username: item.Username,
		Email:    item.Email,
		Status:   item.Status,
	}); err != nil {
		return errors.InternalServerError(err.Error())
	}

	if err = a.AdminRoleDAL.DeleteBy("admin_id = ?", id); err != nil {
		return errors.InternalServerError(err.Error())
	}

	for _, v := range item.RoleIDList {
		if err = a.AdminRoleDAL.Create(&schema.AdminRole{
			AdminId: id,
			RoleId:  v,
		}); err != nil {
			return errors.InternalServerError(err.Error())
		}
	}

	return nil
}

func (a *Admin) Delete(id int) error {
	if err := a.AdminDAL.Delete(id); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}
