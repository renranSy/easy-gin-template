package casbinx

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"os"
	"path"
	"strconv"
)

type CasbinX struct {
	Enforcer *casbin.Enforcer
	DB       *gorm.DB
}

func InitEnforcer(db *gorm.DB) *casbin.Enforcer {
	adapter, _ := gormadapter.NewAdapterByDB(db)
	root, _ := os.Getwd()

	e, err := casbin.NewEnforcer(path.Join(root, "configs", "rbac_model.conf"), adapter)
	if err != nil {
		panic(err)
	}

	if err = e.LoadPolicy(); err != nil {
		panic(err)
	}
	return e
}

func (a CasbinX) ResetPolicy() error {
	type Rule struct {
		ID     int
		Path   string
		Method string
	}

	var list []Rule
	a.DB.Raw(`SELECT role.id, menu_resource.path, menu_resource.method FROM role JOIN role_menu ON role.id = role_menu.role_id JOIN menu_resource ON role_menu.menu_id = menu_resource.menu_id WHERE role.deleted_at IS NULL AND role_menu.deleted_at IS NULL AND menu_resource.deleted_at IS NULL AND menu_resource.path != ''`).Find(&list)
	policy, err := a.Enforcer.GetPolicy()
	if err != nil {
		return err
	}
	if _, err = a.Enforcer.RemovePolicies(policy); err != nil {
		return err
	}

	addPolicy := make([][]string, 0)
	seen := make(map[string]struct{})
	for _, v := range list {
		key := strconv.Itoa(v.ID) + ":" + v.Path + ":" + v.Method
		if _, exists := seen[key]; !exists {
			addPolicy = append(addPolicy, []string{strconv.Itoa(v.ID), v.Path, v.Method})
			seen[key] = struct{}{}
		}
	}

	if err = a.Enforcer.LoadPolicy(); err != nil {
		return err
	}

	if _, err = a.Enforcer.AddPolicies(addPolicy); err != nil {
		return err
	}

	return nil
}
