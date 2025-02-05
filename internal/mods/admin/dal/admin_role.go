package dal

import (
	"easy-gin-template/internal/mods/admin/schema"
	"gorm.io/gorm"
)

type AdminRole struct {
	DB *gorm.DB
}

func (a *AdminRole) Exists(query interface{}, args ...interface{}) (bool, error) {
	var count int64
	a.DB.Model(&schema.AdminRole{}).Where(query, args...).Count(&count)

	return count > 0, nil
}

func (a *AdminRole) Query(query interface{}, args ...interface{}) *[]schema.AdminRole {
	var list []schema.AdminRole
	a.DB.Select("id", "role_id", "admin_id", "created_at", "updated_at").Where(query, args...).Find(&list)
	return &list
}

func (a *AdminRole) Create(item *schema.AdminRole) error {
	if err := a.DB.Create(item).Error; err != nil {
		return err
	}
	return nil
}

func (a *AdminRole) Delete(id int) error {
	if err := a.DB.Delete(&schema.AdminRole{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (a *AdminRole) DeleteBy(query interface{}, args ...interface{}) error {
	if err := a.DB.Where(query, args).Delete(&schema.AdminRole{}).Error; err != nil {
		return err
	}
	return nil
}
