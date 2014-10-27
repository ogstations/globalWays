// Package: models
// File: menu.go
// Created by mint
// Useage: 菜单管理
// DATE: 14-6-26 10:29
package models

import (
	"html/template"
	"strconv"
	"strings"
	"utils/container"
	"github.com/astaxie/beego/orm"
	"utils/convert"
	"gwsAdmin/utils"
	"fmt"
)

type Menu struct {
	Id      int64
	Pid     int64  `orm:"size(11);column(pid)"`
	Name    string `orm:"size(40);column(name)"`
	Url     string `orm:"size(100);column(url)"`
	Data    string `orm:"size(60);column(data)"`
	Order   int64  `orm:"size(11);column(order)"`
	Display bool   `orm:"column(display)"`
}

func (m *Menu) TableName() string {
	return "menu"
}

//根据Id获取菜单信息
func GetMenuById(id int64, ormer orm.Ormer) *Menu {
	menu := new(Menu)
	ormer.QueryTable(menu).Filter("id", id).One(menu)
	return menu
}

//添加菜单
func AddMenu(menu *Menu, ormer orm.Ormer) *Menu {
	id, _ := ormer.Insert(menu)
	menu.Id = id
	return menu
}

//编辑菜单
func UpdateMenu(menu *Menu, ormer orm.Ormer) bool {
	if _, err := ormer.Update(menu); err != nil {
		return false
	}

	return true
}

//删除菜单
func DelMenuById(id int64, ormer orm.Ormer) bool {

	menus := GetMenuByPid(id, utils.Founder, true, ormer)
	if len(menus) > 0 {
		for _, menu := range menus {
			DelMenuById(menu.Id, ormer)
		}
	}

	if _, err := ormer.QueryTable(new(Menu)).Filter("id", id).Delete(); err != nil {
		return false
	}

	return true
}

//按父ID查找菜单子项
func GetMenuByPid(pid, aid int64, whole bool, ormer orm.Ormer) []*Menu {

	menus := make([]*Menu, 0)
	admin := GetAdminById(aid, ormer)
	querySeter := ormer.QueryTable(new(Menu))

	if aid != utils.Founder && len(admin.Role.Data) > 0 {
		querySeter.Filter("pid", pid).Filter("display", whole).Filter("id__in", admin.Role.Data).All(&menus)
	} else {
		querySeter.Filter("pid", pid).Filter("display", whole).All(&menus)
	}

	return menus
}

//获取所有菜单
func GetMenuAll(ormer orm.Ormer)  map[int64][]*Menu {
	menus := make([]*Menu, 0)
	ormer.QueryTable(new(Menu)).OrderBy("order").All(&menus)

	menuList := make(map[int64][]*Menu)
	for _, menu := range menus {
		if _, ok := menuList[menu.Pid]; !ok {
			menuList[menu.Pid] = make([]*Menu, 0)
		}
		menuList[menu.Pid] = append(menuList[menu.Pid], menu)
	}

	return menuList
}

//获取左侧导航菜单
func GetLeftMenuHtml(pid, aid int64, ormer orm.Ormer) template.HTML {
	menus := make([]*Menu, 0)
	admin := GetAdminById(aid, ormer)
	querySeter := ormer.QueryTable(new(Menu))

	if aid != utils.Founder && len(admin.Role.Data) > 0 {
		querySeter.Filter("id__in", admin.Role.Data).OrderBy("order").All(&menus)
	} else {
		querySeter.OrderBy("order").All(&menus)
	}

	menusList := make(map[int64][]*Menu)
	for _, menu := range menus {
		if _, ok := menusList[menu.Pid]; !ok {
			menusList[menu.Pid] = make([]*Menu, 0)
		}
		menusList[menu.Pid] = append(menusList[menu.Pid], menu)
	}

	Html := ""
	for _, menuSec := range menusList[pid] {
		Html += "<h3 class=\"f14\"><span class=\"switchs cu on\" title=\"展开与收缩\"></span>" + menuSec.Name + "</h3>"

		Html += "<ul>"
		for _, menuLast := range menusList[menuSec.Id] {
			menuLastId := convert.Int642str(menuLast.Id)
			Html += "<li id=\"_MP" + menuLastId + "\" class=\"sub_menu\">"
			Html += "<a href=\"javascript:_MP(" + menuLastId + ",'" + menuLast.Url + "');\" hidefocus=\"true\" style=\"outline:none;\">" + menuLast.Name + "</a>"
			Html += "</li>"
		}

		// 快捷面板
		if menuSec.Id == utils.ShourtCutPanel {
			panelList := GetPanelList(aid, ormer)

			for _, panel := range panelList {
				menuInfo := GetMenuById(panel.Mid, ormer)
				menuId := convert.Int642str(menuInfo.Id)
				Html += "<li id=\"_MP" + menuId + "\" class=\"sub_menu\">"
				Html += "<a href=\"javascript:_MP(" + menuId + ",'" + menuInfo.Url + "');\" hidefocus=\"true\" style=\"outline:none;\">" + menuInfo.Name + "</a>"
				Html += "</li>"
			}
		}
		Html += "</ul>"
	}

	return template.HTML(Html)
}

//返回菜单Option的HTML
func GetMenuOptionHtml(id int64, ormer orm.Ormer) template.HTML {

	menuList := GetMenuAll(ormer)

	Html := ""
	for _, menu := range menuList[0] {
		menuId := convert.Int642str(menu.Id)
		if menu.Id == id {
			Html += "<option value=" + menuId + " selected ><b>" + menu.Name + "</b></option>"
		} else {
			Html += "<option value=" + menuId + "><b>" + menu.Name + "</b></option>"
		}

		for _, menuSec := range menuList[menu.Id] {
			menuSecId := convert.Int642str(menuSec.Id)
			if menuSec.Id == id {
				Html += "<option value=" + menuSecId + " selected >&#12288;&#8866;" + menuSec.Name + "</option>"
			} else {
				Html += "<option value=" + menuSecId + ">&#12288;&#8866;" + menuSec.Name + "</option>"
			}

			for _, menuLast := range menuList[menuSec.Id] {
				menuLastId := convert.Int642str(menuLast.Id)
				if menuLast.Id == id {
					Html += "<option value=" + menuLastId + " selected >&#12288;&#12288;&#8866;" + menuLast.Name + "</option>"
				} else {
					Html += "<option value=" + menuLastId + ">&#12288;&#12288;&#8866;" + menuLast.Name + "</option>"
				}
			}
		}
	}

	return template.HTML(Html)
}

//返回菜单树
func GetMenuTree(role string, ormer orm.Ormer) template.HTML {
	menuList := GetMenuAll(ormer)

	//解析权限
	arr_role := strings.Split(role, ",")

	Html := "<SCRIPT type=\"text/javascript\">var zNodes =["
	for _, menu := range menuList[0] {
		menuId := convert.Int642str(menu.Id)
		menuPid := convert.Int642str(menu.Pid)
		if container.InArray(menuId, arr_role) {
			Html += "{ id:" + menuId + ", pId:" + menuPid + ", name:'" + menu.Name + "', open:true, checked:true},"
		} else {
			Html += "{ id:" + menuId + ", pId:" + menuPid + ", name:'" + menu.Name + "', open:true},"
		}

		for _, menuSec := range menuList[menu.Id] {
			menuSecId := convert.Int642str(menuSec.Id)
			menuSecPid := convert.Int642str(menuSec.Pid)
			if container.InArray(menuSecId, arr_role) {
				Html += "{ id:" + menuSecId + ", pId:" + menuSecPid + ", name:'" + menuSec.Name + "', open:true, checked:true},"
			} else {
				Html += "{ id:" + menuSecId + ", pId:" + menuSecPid + ", name:'" + menuSec.Name + "', open:true},"
			}

			for _, menuLast := range menuList[menuSec.Id] {
				menuLastId := convert.Int642str(menuLast.Id)
				menuLastPid := convert.Int642str(menuLast.Pid)
				if container.InArray(menuLastId, arr_role) {
					Html += "{ id:" + menuLastId + ", pId:" + menuLastPid + ", name:'" + menuLast.Name + "', checked:true},"
				} else {
					Html += "{ id:" + menuLastId + ", pId:" + menuLastPid + ", name:'" + menuLast.Name + "'},"
				}
			}
		}
	}

	Html += "];</SCRIPT>"
	return template.HTML(Html)
}

//返回后台地图
func GetMenuMap(ormer orm.Ormer) template.HTML {
	menuList := GetMenuAll(ormer)

	Html := ""
	n := 1
	for _, menu := range menuList[0] {
		if n == 1 {
			Html += "<div class=\"map-menu lf\">"
		}

		Html += "<ul>"
		Html += "<li class=\"title\">" + menu.Name + "</li>"

		for _, menuSec := range menuList[menu.Id] {
			Html += "<li class=\"title2\">" + menuSec.Name + "</li>"

			for _, menuLast := range menuList[menuSec.Id] {
				Html += "<li><a href=\"javascript:Go(" + convert.Int642str(menuLast.Id) + ",'" + menuLast.Url + "')\">" + menuLast.Name + "</a></li>"
			}
		}

		Html += "</ul>"
		if n%2 == 0 {
			Html += "</div><div class=\"map-menu lf\">"
		}
		n++
	}

	return template.HTML(Html)
}

//获取所有菜单
//返回HTML
func GetMenuHtml(ormer orm.Ormer) template.HTML {

	menuList := GetMenuAll(ormer)

	Html := ""
	for _, menu := range menuList[0] {
		menuId := convert.Int642str(menu.Id)
		Html += "<tr>"
		Html += "<td align=\"left\"><b>" + menu.Name + "</b></td>"
		Html += "<td align=\"center\">" + convert.Int642str(menu.Order) + "</td>"
		Html += "<td align=\"center\">" + menu.Url + "</td>"

		if menu.Display {
			Html += "<td align=\"center\">菜单显示</td>"
		} else {
			Html += "<td align=\"center\">菜单隐藏</td>"
		}

		Html += "<td align=\"center\">"
		Html += "<a href=\"/menu/add/" + menuId + "/\">添加子菜单</a> |"
		Html += "<a href=\"/menu/edit/" + menuId + "/\">修改</a> |"
		Html += "<a onclick=\"delete_menu(" + menuId + ")\" href=\"javascript:;\">删除</a>"

		Html += "</td>"
		Html += "</tr>"

		for _, menuSec := range menuList[menu.Id] {
			menuSecId := strconv.FormatInt(menuSec.Id, 10)
			Html += "<tr>"
			Html += "<td align=\"left\">&nbsp;&nbsp;&#12288;&#8866;&nbsp;&nbsp;" + menuSec.Name + "</td>"
			Html += "<td align=\"center\">" + convert.Int642str(menuSec.Order) + "</td>"
			Html += "<td align=\"center\">" + menuSec.Url + "</td>"
			if menuSec.Display {
				Html += "<td align=\"center\">菜单显示</td>"
			} else {
				Html += "<td align=\"center\">菜单隐藏</td>"
			}

			Html += "<td align=\"center\">"
			Html += "<a href=\"/menu/add/" + menuSecId + "/\">添加子菜单</a> |"
			Html += "<a href=\"/menu/edit/" + menuSecId + "/\">修改</a> |"
			Html += "<a onclick=\"delete_menu(" + menuSecId + ")\" href=\"javascript:;\">删除</a>"
			Html += "</td>"
			Html += "</tr>"

			for _, menuLast := range menuList[menuSec.Id] {
				menuLastId := convert.Int642str(menuLast.Id)
				Html += "<tr>"
				Html += "<td align=\"left\">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&#12288;&#8866;&nbsp;&nbsp;" + menuLast.Name + "</td>"
				Html += "<td align=\"center\">" + convert.Int642str(menuLast.Order) + "</td>"
				Html += "<td align=\"center\">" + menuLast.Url + "</td>"

				if menuLast.Display {
					Html += "<td align=\"center\">菜单显示</td>"
				} else {
					Html += "<td align=\"center\">菜单隐藏</td>"
				}

				Html += "<td align=\"center\">"
				Html += "<a href=\"/menu/edit/" + menuLastId + "/\">修改</a> |"
				Html += "<a onclick=\"delete_menu(" + menuLastId + ")\" href=\"javascript:;\">删除</a>"
				Html += "</td>"
				Html += "</tr>"
			}
		}
	}

	return template.HTML(Html)
}

//当前位置
func GetMenuPos(id int64, ormer orm.Ormer) string {
	menu := GetMenuById(id, ormer)
	println(fmt.Sprintf("menu: %+v", menu))
	str := ""
	if menu.Pid > 0 {
		str += GetMenuPos(menu.Pid, ormer)
	}

	return str + menu.Name + " > "
}
