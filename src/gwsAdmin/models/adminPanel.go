// Package: models
// File: admin_panel.go
// Created by mint
// Useage: 快捷菜单
// DATE: 14-6-26 10:58
package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type AdminPanel struct {
	Id         int64
	Mid        int64  `orm:"size(11);column(mid)"`
	Menu       *Menu  `orm:"-"`
	Aid        int64  `orm:"size(11);column(aid)"`
	Name       string `orm:"size(40);column(name)"`
	Url        string `orm:"size(100);column(url)"`
	Created    time.Time `orm:"auto_now_add"`
	Updated    time.Time `orm:"auto_now"`
}

func (m *AdminPanel) TableName() string {
	return "adminPanel"
}

//根据Id获取信息
func GetPanelById(id int64, ormer orm.Ormer) *AdminPanel {
	adminPanel := new(AdminPanel)
	ormer.QueryTable(adminPanel).Filter("id", id).One(adminPanel)
	adminPanel.Menu = GetMenuById(adminPanel.Mid, ormer)
	return adminPanel
}

//获取快捷方式列表
func GetPanelList(aid int64, ormer orm.Ormer) []*AdminPanel {
	//初始化菜单
	adminPanels := make([]*AdminPanel, 0)

	ormer.QueryTable(new(AdminPanel)).Filter("aid", aid).All(&adminPanels)
	for i, v := range adminPanels {
		adminPanels[i].Menu = GetMenuById(v.Mid, ormer)
	}

	return adminPanels
}

//根据mid获取快捷方式
func GetPanelByMid(mid int64, aid int64, ormer orm.Ormer) *AdminPanel {
	adminPanel := new(AdminPanel)
	ormer.QueryTable(adminPanel).Filter("mid", mid).Filter("aid", aid).One(adminPanel)
	adminPanel.Menu = GetMenuById(mid, ormer)
	return adminPanel
}

//是否已添加快捷方式
func isPanelExist(mid int64, aid int64, ormer orm.Ormer) bool {
	adminPanel := new(AdminPanel)
	cnt, _ := ormer.QueryTable(adminPanel).Filter("mid", mid).Filter("aid", aid).Count()
	return cnt != 0
}

//删除快捷方式
func DelPanel(mid, aid int64, ormer orm.Ormer) bool {

	if _, err := ormer.QueryTable(new(AdminPanel)).Filter("mid", mid).Filter("aid", aid).Delete(); err != nil {
		return false
	}

	return true
}

//添加快捷方式
func AddPanel(mid, aid int64, ormer orm.Ormer) *AdminPanel {
	menu := GetMenuById(mid, ormer)

	adminPanel := &AdminPanel{
		Mid: mid,
		Aid: aid,
		Name: menu.Name,
		Url: menu.Url,
	}

	id, _ := ormer.Insert(adminPanel)
	adminPanel.Id = id

	return adminPanel
}

