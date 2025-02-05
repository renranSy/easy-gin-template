package dal

import (
	"easy-gin-template/internal/mods/admin/schema"
	"gorm.io/gorm"
)

type Admin struct {
	DB *gorm.DB
}

func (a *Admin) GetByUsernameAndPassword(username string, password string) (*schema.Admin, error) {
	var admin *schema.Admin
	if err := a.DB.Select("id, username, password").Where("username = ? AND password = ?", username, password).Take(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (a *Admin) Exists(query interface{}, args ...interface{}) (bool, error) {
	var count int64
	a.DB.Model(&schema.Admin{}).Where(query, args...).Count(&count)

	return count > 0, nil
}

func (a *Admin) Query() (*[]schema.Admin, error) {
	var list []schema.Admin
	if err := a.DB.Select("id", "username", "email", "status", "created_at", "updated_at").Find(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

func (a *Admin) Create(admin *schema.Admin) error {
	if err := a.DB.Create(admin).Error; err != nil {
		return err
	}
	return nil
}

func (a *Admin) Update(admin *schema.Admin) error {
	if err := a.DB.Model(admin).Select("username", "email", "status").Updates(*admin).Error; err != nil {
		return err
	}
	return nil
}

func (a *Admin) Delete(id int) error {
	if err := a.DB.Delete(&schema.Admin{}, id).Error; err != nil {
		return err
	}
	return nil
}
