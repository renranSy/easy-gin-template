package biz

import (
	"easy-gin-template/internal/mods/admin/dal"
	"easy-gin-template/internal/mods/admin/schema"
	"easy-gin-template/pkg/db"
	"easy-gin-template/pkg/errors"
)

type Role struct {
	RoleDAL *dal.Role
}

func (a *Role) Query() *[]schema.Role {
	return a.RoleDAL.Query()
}

func (a *Role) Create(item *schema.RoleForm) error {
	exists, err := a.RoleDAL.Exists("name = ?", item.Name)
	if err != nil {
		return err
	}
	if exists {
		return errors.BadRequest("角色名已存在")
	}

	if err = a.RoleDAL.Create(&schema.Role{
		Name:        item.Name,
		Description: item.Description,
		Sequence:    item.Sequence,
		Status:      item.Status,
	}); err != nil {
		return errors.InternalServerError(err.Error())
	}

	return nil
}

func (a *Role) Update(id int, item *schema.RoleForm) error {
	exists, err := a.RoleDAL.Exists("id = ?", id)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	if !exists {
		return errors.NotFound("当前角色不存在")
	}
	exists, err = a.RoleDAL.Exists("name = ? AND id != ?", item.Name, id)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	if exists {
		return errors.BadRequest("角色名已存在")
	}
	if err = a.RoleDAL.Update(&schema.Role{
		Model:       db.Model{ID: id},
		Name:        item.Name,
		Description: item.Description,
		Sequence:    item.Sequence,
		Status:      item.Status,
	}); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}

func (a *Role) Delete(id int) error {
	if err := a.RoleDAL.Delete(id); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}
