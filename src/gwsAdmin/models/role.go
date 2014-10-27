// Package: models
// File: role.go
// Created by mint
// Useage: 角色表
// DATE: 14-6-25 20:03
package models

//角色表
import (
	"html/template"
	"time"
	"utils/page"
	"github.com/astaxie/beego/orm"
)

type Role struct {
	Id         int64
	RoleName   string `orm:"unique;column(role_name)"`
	Desc       string
	Data       string `orm:"type(text)"` //菜单列表
	Status     bool
	Created    time.Time `orm:"auto_now_add"`
	Updated    time.Time `orm:"auto_now"`
}

func (m *Role) TableName() string {
	return "role"
}

//根据ID获取角色信息
func GetRoleById(id int64, ormer orm.Ormer) *Role {
	role := new(Role)
	ormer.QueryTable(role).Filter("id", id).One(role)
	return role
}

//获取角色列表
func GetRoleByPage(pager *page.Page, ormer orm.Ormer) ([]*Role, template.HTML) {
	roleList := make([]*Role, 0)
	querySeter := ormer.QueryTable(new(Role))

	//查询总数
	cnt, _ := querySeter.Count()

	// 分页
	pager.Nums = cnt
	pager.SubPage_type = 2
	pages := pager.Show()

	//查询数据
	if cnt > 0 {
		querySeter.Limit(pager.Perpage, (pager.Current_page - 1)*pager.Perpage).All(&roleList)
	}

	return roleList, pages
}

//获取角色列表
func GetRoleList(ormer orm.Ormer) []*Role {
	roleList := make([]*Role, 0)
	ormer.QueryTable(new(Role)).All(&roleList)
	return roleList
}

//添加角色
func AddRole(role *Role, ormer orm.Ormer) bool {
	if _, err := ormer.Insert(role); err != nil {
		return false
	}

	return true
}

//编辑角色
func UpdateRole(role *Role, ormer orm.Ormer) bool {
	if _, err := ormer.Update(role); err != nil {
		return false
	}

	return true
}

//设置状态
func SetRoleStatus(id int64, status bool, ormer orm.Ormer) bool {
	if _, err := ormer.QueryTable(new(Role)).Filter("id", id).Update(orm.Params{"status":status}); err != nil {
		return false
	}

	return true
}

//删除角色
func DelRoleById(id int64, ormer orm.Ormer) bool {
	if _, err := ormer.QueryTable(new(Role)).Filter("id", id).Delete(); err != nil {
		return false
	}

	return true
}
