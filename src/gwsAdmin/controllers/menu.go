// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"gwsAdmin/models"
	"gwsAdmin/utils"
)

type MenuController struct {
	BaseController
}

// @router /menu [get]
func (c *MenuController) Index() {
	if adminInfo, ok := getAdminInfoBySession(c.Controller); ok {
		c.Data["menus"] = models.GetMenuHtml(models.Reader)
		c.Data["controller"], _ = c.GetControllerAndAction()
		// 管理员信息
		c.Data["admin"] = adminInfo
	} else {
		c.Redirect(beego.UrlFor("AdminController.Login"), 302)
		return
	}

	c.TplNames = "setting/menu/index.html"
}

// @router /menu/add/:pid [get]
// @router /menu/add/:pid [post]
func (c *MenuController) AddMenuByPid() {
	pid, err := c.GetInt(":pid")
	if err != nil {
		pid = 0
	}

	addMenu(pid, c)
}

// @router /menu/add [get]
// @router /menu/add [post]
func (c *MenuController) AddMenuTop() {
	addMenu(0, c)
}

// @router /menu/delete/:mid [delete]
func (c *MenuController) DelMenu() {
	c.Controller.EnableRender = false

	data := make(map[string]string)
	adminInfo, _ := getAdminInfoBySession(c.Controller)
	mid, e := c.GetInt(":mid")

	if adminInfo.Id != utils.Founder {
		data["status"] = "0"
		data["message"] = c.Message("menu_del_no_role")
	}else if e != nil {
		data["status"] = "0"
		data["message"] = c.Message("menu_del_param_error")
	} else {
		if models.DelMenuById(mid, models.Reader) {

			controller, action := c.GetControllerAndAction()

			// 添加日志
			addLogs(adminInfo.Id, controller, action, fmt.Sprintf("删除菜单:%v|^|用户[%v, %v]", mid, adminInfo.Id, adminInfo.UserName))

			data["status"] = "1"
			data["message"] = c.Message("menu_del_success")
		} else {
			data["status"] = "0"
			data["message"] = c.Message("menu_del_error")
		}
	}

	c.Data["json"] = data
	c.ServeJson()
}

// @router /menu/edit/:mid [get]
// @router /menu/edit/:mid [post]
func (c *MenuController) EditMenu() {
	mid, e := c.GetInt(":mid")
	if e != nil {
		mid = 0
	}

	if c.Controller.Ctx.Request.Method == "GET" {

		menuInfo := models.GetMenuById(mid, models.Reader)
		menus := models.GetMenuOptionHtml(menuInfo.Pid, models.Reader)

		c.Data["menu_info"] = menuInfo
		c.Data["menus"] = menus

		c.TplNames = "setting/menu/edit.html"
	} else if c.Controller.Ctx.Request.Method == "POST" {
		c.Controller.EnableRender = false
		menu := new(models.Menu)
		flash := beego.NewFlash()

		menu.Id = mid
		//参数检查
		pid, _ := c.GetInt("pid")
		if pid < 0 {
			flash.Error(c.Message("menu_upd_pid_error"))
			flash.Data["url"] = c.Controller.Ctx.Request.URL.Path
			flash.Store(&c.Controller)

			c.Redirect(beego.UrlFor("PublicController.Message"), 302)
			return
		} else {
			menu.Pid = pid
		}

		name := c.GetString("name")
		if len(name) < 0 {
			flash.Error(c.Message("menu_upd_name_error"))
			flash.Data["url"] = c.Controller.Ctx.Request.URL.Path
			flash.Store(&c.Controller)

			c.Redirect(beego.UrlFor("PublicController.Message"), 302)
			return
		} else {
			menu.Name = name
		}

		url := c.GetString("url")
		if len(url) < 0 {
			flash.Error(c.Message("menu_upd_url_error"))
			flash.Data["url"] = c.Controller.Ctx.Request.URL.Path
			flash.Store(&c.Controller)

			c.Redirect(beego.UrlFor("PublicController.Message"), 302)
			return
		} else {
			menu.Url = url
		}

		if order, e := c.GetInt("order"); e != nil {
			flash.Error(c.Message("menu_upd_order_error", e))
			flash.Data["url"] = c.Controller.Ctx.Request.URL.Path
			flash.Store(&c.Controller)

			c.Redirect(beego.UrlFor("PublicController.Message"), 302)
			return
		} else {
			menu.Order = order
		}

		menu.Data = c.GetString("data")

		if display, e := c.GetBool("display"); e != nil {
			flash.Error(c.Message("menu_upd_display_error", e))
			flash.Data["url"] = c.Controller.Ctx.Request.URL.Path
			flash.Store(&c.Controller)

			c.Redirect(beego.UrlFor("PublicController.Message"), 302)
			return
		} else {
			menu.Display = display
		}

		// 菜单添加
		if models.UpdateMenu(menu, models.Writter) {
			adminInfo, _ := getAdminInfoBySession(c.Controller)
			controller, action := c.GetControllerAndAction()
			// 添加日志
			addLogs(adminInfo.Id, controller, action, fmt.Sprintf("更新菜单:%v|^|用户[%v, %v]", menu.Name, adminInfo.Id, adminInfo.UserName))

			flash.Data["success"] = c.Message("menu_upd_success")
			flash.Data["url"] = beego.UrlFor("MenuController.Index")
			flash.Store(&c.Controller)

			c.Redirect(beego.UrlFor("PublicController.Message"), 302)
			return
		} else {
			flash.Error(c.Message("menu_upd_error"))
			flash.Data["url"] = c.Controller.Ctx.Request.URL.Path
			flash.Store(&c.Controller)

			c.Redirect(beego.UrlFor("PublicController.Message"), 302)
			return
		}
	}
}

// 添加菜单
func addMenu(pid int64, c *MenuController) {
	if c.Controller.Ctx.Request.Method == "GET" {
		c.Data["menus"] = models.GetMenuOptionHtml(pid, models.Reader)
		c.TplNames = "setting/menu/add.html"
	} else if c.Controller.Ctx.Request.Method == "POST" {
		c.Controller.EnableRender = false
		menu := new(models.Menu)
		flash := beego.NewFlash()

		//参数检查
		if pid < 0 {
			flash.Error(c.Message("menu_add_pid_error"))
			flash.Data["url"] = c.Controller.Ctx.Request.URL.Path
			flash.Store(&c.Controller)

			c.Redirect(beego.UrlFor("PublicController.Message"), 302)
			return
		} else {
			menu.Pid = pid
		}

		name := c.GetString("name")
		if len(name) < 0 {
			flash.Error(c.Message("menu_add_name_error"))
			flash.Data["url"] = c.Controller.Ctx.Request.URL.Path
			flash.Store(&c.Controller)

			c.Redirect(beego.UrlFor("PublicController.Message"), 302)
			return
		} else {
			menu.Name = name
		}

		url := c.GetString("url")
		if len(url) < 0 {
			flash.Error(c.Message("menu_add_url_error"))
			flash.Data["url"] = c.Controller.Ctx.Request.URL.Path
			flash.Store(&c.Controller)

			c.Redirect(beego.UrlFor("PublicController.Message"), 302)
			return
		} else {
			menu.Url = url
		}

		if order, e := c.GetInt("order"); e != nil {
			flash.Error(c.Message("menu_add_order_error", e))
			flash.Data["url"] = c.Controller.Ctx.Request.URL.Path
			flash.Store(&c.Controller)

			c.Redirect(beego.UrlFor("PublicController.Message"), 302)
			return
		} else {
			menu.Order = order
		}

		menu.Data = c.GetString("data")

		if display, e := c.GetBool("display"); e != nil {
			flash.Error(c.Message("menu_add_display_error", e))
			flash.Data["url"] = c.Controller.Ctx.Request.URL.Path
			flash.Store(&c.Controller)

			c.Redirect(beego.UrlFor("PublicController.Message"), 302)
			return
		} else {
			menu.Display = display
		}

		// 菜单添加
		if models.AddMenu(menu, models.Writter).Id > 0 {
			adminInfo, _ := getAdminInfoBySession(c.Controller)
			controller, action := c.GetControllerAndAction()
			// 添加日志
			addLogs(adminInfo.Id, controller, action, fmt.Sprintf("添加菜单:%v|^|用户[%v, %v]", menu.Name, adminInfo.Id, adminInfo.UserName))

			flash.Data["success"] = c.Message("menu_add_success")
			flash.Data["url"] = beego.UrlFor("MenuController.Index")
			flash.Store(&c.Controller)

			c.Redirect(beego.UrlFor("PublicController.Message"), 302)
			return
		} else {
			flash.Error(c.Message("menu_add_error"))
			flash.Data["url"] = c.Controller.Ctx.Request.URL.Path
			flash.Store(&c.Controller)

			c.Redirect(beego.UrlFor("PublicController.Message"), 302)
			return
		}
	}
}
