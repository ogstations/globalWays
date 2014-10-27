// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"github.com/astaxie/beego"
	"gwsAdmin/models"
	"utils/page"
)

var (
	_ beego.Controller
)

type RoleController struct {
	BaseController
}

// @router /role [get]
func (c *RoleController) Index() {
	pageC, err := c.GetInt("page")
	if err != nil {
		pageC = 1
	}

	size, err := c.GetInt("size")
	if err != nil {
		size = 20
	}

	pager := &page.Page{
		Perpage: size,
		Current_page: pageC,
		SubPage_link: beego.UrlFor("RoleController.Index"),
	}
	roleList, pageHtml := models.GetRoleByPage(pager, models.Reader)
	c.Data["role_list"] = roleList
	c.Data["pages"] = pageHtml

	c.TplNames = "setting/role/index.html"
}

// @router /role/member/:rid [get]
func (c *RoleController) Member() {

	roleId, err := c.GetInt(":rid")
	if err != nil {
		roleId = 0
	}
	pageC, err := c.GetInt("page")
	if err != nil {
		pageC = 1
	}
	size, err := c.GetInt("size")
	if err != nil {
		size = 20
	}

	pager := &page.Page{
		Perpage: size,
		Current_page: pageC,
		SubPage_link: beego.UrlFor("RoleController.ListAdmin"),
	}
	admins, pageHtml := models.ListAdmin(pager, roleId, models.Reader)
	c.Data["admin_list"] = admins
	c.Data["pages"] = pageHtml

	c.TplNames = "setting/role/member.html"
}


// @router /role/add [get]
func (c *RoleController) AddRole() {

	c.TplNames = "setting/role/add.html"
}

// @router /role/edit [get]
func (c *RoleController) EditRole() {

	c.TplNames = "setting/role/edit.html"
}
