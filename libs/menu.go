package libs

import (
	"fmt"
	"github.com/codyi/grabc/models"
	"strings"
)

//重新整理菜单，返还可以显示的菜单数据
type newMenu struct {
	Url  string
	Name string
	Icon string
}

type MenuGroup struct {
	Parent newMenu
	Child  []newMenu
}

//获取用户可以访问的菜单
func AccessMenus() []*MenuGroup {
	returnMenus := make([]*MenuGroup, 0)
	menu := models.Menu{}
	menus, err := menu.ListAll()
	if err != nil {
		return returnMenus
	}

	allAccessRoutes := AccessRoutes()
	type temp struct {
		Parent *models.Menu
		Child  []*models.Menu
	}

	tt := make(map[int]temp, 0)
	//归类子菜单，并检查完子菜单权限
	for _, m := range menus {
		if m.Parent == 0 {
			t := temp{}
			t.Parent = m
			tt[m.Id] = t
		} else {
			r := strings.Split(m.Url, "/")
			controllerName := r[0]
			routeName := r[1]
			if CheckAccess(controllerName, routeName, allAccessRoutes) {
				t := tt[m.Parent]
				if t.Parent == nil && len(t.Child) == 0 {
					t = temp{}
				}

				t.Child = append(t.Child, m)
				tt[m.Parent] = t
			}
		}
	}

	//检查完父级菜单权限，如果有子菜单，这个父级菜单将显示
	for i, t := range tt {
		if len(t.Child) > 0 {
			continue
		}

		r := strings.Split(t.Parent.Url, "/")
		controllerName := r[0]
		routeName := r[1]
		if !CheckAccess(controllerName, routeName, allAccessRoutes) {
			delete(tt, i)
		}
	}

	for _, t := range tt {
		m := MenuGroup{}
		p := newMenu{}
		p.Name = t.Parent.Name
		p.Url = "/" + t.Parent.Url

		m.Parent = p

		childMenus := make([]newMenu, 0)

		for _, m := range t.Child {
			nm := newMenu{}
			nm.Name = m.Name
			nm.Url = "/" + m.Url
			childMenus = append(childMenus, nm)
		}

		m.Child = childMenus
		returnMenus = append(returnMenus, &m)
	}

	return returnMenus
}

func ShowMenu(controllName, actionName string) string {
	html := `<ul class='sidebar-menu tree' data-widget='tree'>`
	for _, menu := range AccessMenus() {

		if len(menu.Child) > 0 {
			childHtml := ""
			isActiveChild := false
			for _, childMenu := range menu.Child {
				activeClass := ""

				if strings.ToLower("/"+controllName+"/"+actionName) == strings.ToLower(childMenu.Url) {
					activeClass = "active"
					isActiveChild = true
				}

				childHtml += fmt.Sprintf(`<li class="%s"><a href='%s'>%s%s</a></li>`, activeClass, childMenu.Url, childMenu.Icon, childMenu.Name)
			}

			s := `<li class='treeview %s'><a href='#'>%s<span>%s</span><span class='pull-right-container'><i class='fa fa-angle-left pull-right'></i></span></a><ul class='treeview-menu'>%s</ul></li>`
			if isActiveChild {
				html += fmt.Sprintf(s, "active menu-open", menu.Parent.Icon, menu.Parent.Name, childHtml)
			} else {
				html += fmt.Sprintf(s, "", menu.Parent.Icon, menu.Parent.Name, childHtml)
			}
		} else {
			activeClass := ""
			s := `<li class='treeview %s'><a href='%s'>%s<span>%s</span></a></li>`

			if strings.ToLower("/"+controllName+"/"+actionName) == strings.ToLower(menu.Parent.Url) {
				activeClass = "active"
			}

			html += fmt.Sprintf(s, activeClass, menu.Parent.Url, menu.Parent.Icon, menu.Parent.Name)
		}

	}

	html += `</ul>`
	return html
}
