package dal

import (
	"easy-gin-template/internal/mods/admin/schema"
	"gorm.io/gorm"
)

type Menu struct {
	DB *gorm.DB
}

func (a *Menu) Query() *[]schema.Menu {
	var list []schema.Menu
	a.DB.Select("id", "name", "code", "description", "sequence", "type", "path", "status", "parent_id", "created_at").Find(&list)
	return &list
}

func (a *Menu) Get(query interface{}, args ...interface{}) (*schema.Menu, error) {
	var menu schema.Menu

	if err := a.DB.Select("id", "name", "code", "description", "sequence", "type", "path", "status", "parent_id", "created_at").Where(query, args).Take(&menu).Error; err != nil {
		return nil, err
	}

	return &menu, nil
}

func (a *Menu) Exists(query interface{}, args ...interface{}) (bool, error) {
	var count int64
	a.DB.Model(&schema.Menu{}).Where(query, args).Count(&count)

	return count > 0, nil
}

func (a *Menu) Create(menu *schema.Menu) {
	a.DB.Create(menu)
}

func (a *Menu) Update(menu *schema.Menu) error {
	if err := a.DB.Model(menu).Select("name", "code", "description", "sequence", "type", "path", "status").Updates(*menu).Error; err != nil {
		return err
	}
	return nil
}

func (a *Menu) Delete(id int) error {
	if err := a.DB.Delete(&schema.Menu{}, id).Error; err != nil {
		return err
	}
	return nil
}
