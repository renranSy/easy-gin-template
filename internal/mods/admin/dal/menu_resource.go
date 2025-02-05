package dal

import (
	"easy-gin-template/internal/mods/admin/schema"
	"gorm.io/gorm"
)

type MenuResource struct {
	DB *gorm.DB
}

func (a *MenuResource) Create(menuResource *schema.MenuResource) {
	a.DB.Create(menuResource)
}

func (a *MenuResource) Query(query interface{}, args ...interface{}) *[]schema.MenuResource {
	var list []schema.MenuResource
	a.DB.Where(query, args).Find(&list)
	return &list
}

func (a *MenuResource) Delete(id int) {
	a.DB.Delete(&schema.MenuResource{}, id)
}

func (a *MenuResource) DeleteBy(query interface{}, args ...interface{}) error {
	if err := a.DB.Where(query, args).Delete(&schema.MenuResource{}).Error; err != nil {
		return err
	}
	return nil
}
