package biz

import (
	"easy-gin-template/internal/mods/admin/dal"
	"easy-gin-template/internal/mods/admin/schema"
	"easy-gin-template/pkg/casbinx"
	"easy-gin-template/pkg/db"
	"easy-gin-template/pkg/errors"
)

type Menu struct {
	MenuDAL      *dal.Menu
	MenuResource *dal.MenuResource
	CasbinX      *casbinx.CasbinX
}

func (a *Menu) Query() *[]*schema.Menu {
	list := a.MenuDAL.Query()

	menuMap := make(map[int]*schema.Menu)
	menus := make([]*schema.Menu, 0)
	for _, v := range *list {
		if _, exists := menuMap[v.ID]; exists {
			children := menuMap[v.ID].Children
			menuMap[v.ID] = &v
			menuMap[v.ID].Children = children
		} else {
			menuMap[v.ID] = &v
		}

		// 存入跟节点
		if v.ParentID == 0 {
			menus = append(menus, menuMap[v.ID])
			continue
		}

		// 父节点存在时，子节点存入父节点的子级
		if parent, exists := menuMap[v.ParentID]; exists {
			if menuMap[v.ParentID].Children == nil {
				parent.Children = []*schema.Menu{menuMap[v.ID]}
			} else {
				parent.Children = append(menuMap[v.ParentID].Children, menuMap[v.ID])
			}
			continue
		}

		// 父节点不存在时，创建父节点并存入 map，子节点存入该父节点的子级
		menuMap[v.ParentID] = &schema.Menu{Children: []*schema.Menu{menuMap[v.ID]}}
	}

	return &menus
}

func (a *Menu) Create(item *schema.MenuForm) error {
	// TODO 使用事务
	menu := new(schema.Menu)
	item.FillTo(menu)
	a.MenuDAL.Create(menu)

	for _, v := range item.MenuResources {
		menuResource := new(schema.MenuResource)
		menuResource.Path = v.Path
		menuResource.Method = v.Method
		menuResource.MenuID = menu.ID
		a.MenuResource.Create(menuResource)
	}

	return nil
}

func (a *Menu) Get(id int) (*schema.Menu, error) {
	menu, err := a.MenuDAL.Get(id)
	if err != nil {
		return nil, errors.NotFound(err.Error())
	}
	if menu.ParentID != 0 {
		parentMenu, err := a.MenuDAL.Get(menu.ParentID)
		if err != nil {
			return nil, errors.NotFound(err.Error())
		}
		menu.ParentName = parentMenu.Name
	}
	menu.MenuResources = a.MenuResource.Query("menu_id = ?", menu.ID)
	return menu, nil
}

func (a *Menu) Update(id int, item *schema.MenuForm) error {
	// TODO 使用事务
	exists, err := a.MenuDAL.Exists("id = ?", id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.NotFound("当前菜单不存在")
	}
	if err = a.MenuDAL.Update(&schema.Menu{
		Model:       db.Model{ID: id},
		Name:        item.Name,
		Code:        item.Code,
		Description: item.Description,
		Sequence:    item.Sequence,
		Type:        item.Type,
		Path:        item.Path,
		Status:      item.Status,
	}); err != nil {
		return err
	}

	err = a.MenuResource.DeleteBy("menu_id = ?", id)
	if err != nil {
		return err
	}
	for _, v := range item.MenuResources {
		a.MenuResource.Create(&schema.MenuResource{
			MenuID: id,
			Method: v.Method,
			Path:   v.Path,
		})
	}

	if err = a.CasbinX.ResetPolicy(); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}

func (a *Menu) Delete(id int) error {
	// TODO 使用事务
	if err := a.MenuDAL.Delete(id); err != nil {
		return err
	}

	if err := a.MenuResource.DeleteBy("menu_id = ?", id); err != nil {
		return err
	}

	if err := a.CasbinX.ResetPolicy(); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}
