package dal

import (
	"easy-gin-template/internal/mods/admin/schema"
	"gorm.io/gorm"
)

type Role struct {
	DB *gorm.DB
}

func (a *Role) Exists(query interface{}, args ...interface{}) (bool, error) {
	var count int64
	a.DB.Model(&schema.Role{}).Where(query, args...).Count(&count)

	return count > 0, nil
}

func (a *Role) Query() *[]schema.Role {
	var list []schema.Role
	a.DB.Select("id", "name", "description", "sequence", "status", "created_at", "updated_at").Find(&list)
	return &list
}

func (a *Role) Create(role *schema.Role) error {
	if err := a.DB.Create(role).Error; err != nil {
		return err
	}
	return nil
}

func (a *Role) Update(role *schema.Role) error {
	if err := a.DB.Model(role).Select("name", "description", "sequence", "status").Updates(role).Error; err != nil {
		return err
	}
	return nil
}

func (a *Role) Delete(id int) error {
	if err := a.DB.Delete(&schema.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}
