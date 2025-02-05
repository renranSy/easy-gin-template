package dal

import (
	"easy-gin-template/internal/mods/admin/schema"
	"gorm.io/gorm"
)

type RoleMenu struct {
	DB *gorm.DB
}

func (a *RoleMenu) Exists(query interface{}, args ...interface{}) (bool, error) {
	var count int64
	a.DB.Model(&schema.RoleMenu{}).Where(query, args...).Count(&count)

	return count > 0, nil
}

func (a *RoleMenu) Query(query interface{}, args ...interface{}) *[]schema.RoleMenu {
	var list []schema.RoleMenu
	a.DB.Select("id", "role_id", "menu_id", "created_at", "updated_at").Where(query, args...).Find(&list)
	return &list
}

func (a *RoleMenu) Create(item *schema.RoleMenu) error {
	if err := a.DB.Create(item).Error; err != nil {
		return err
	}
	return nil
}

func (a *RoleMenu) Delete(id int) error {
	if err := a.DB.Delete(&schema.RoleMenu{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (a *RoleMenu) DeleteBy(query interface{}, args ...interface{}) error {
	if err := a.DB.Where(query, args).Delete(&schema.RoleMenu{}).Error; err != nil {
		return err
	}
	return nil
}
