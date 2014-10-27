// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"gwsAdmin/models"
	"gwsAdmin/utils"
	"utils/security"
)

var (
	_   beego.Controller
	cpt *captcha.Captcha
)

func init() {
	// use beego cache system store the captcha data
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
}

//管理员controller
type AdminController struct {
	BaseController
}

// @router /login [get]
// @router /login [post]
func (c *AdminController) Login() {
	if c.Controller.Ctx.Request.Method == "GET" {
		c.TplNames = "admin/login.html"
	} else if c.Controller.Ctx.Request.Method == "POST" {
		flash := beego.NewFlash()

		username := c.GetString("username")
		password := c.GetString("password")
		captchaId := c.GetString("captcha_id")
		verify := c.GetString("verify")

		if lang := c.GetSession(utils.Session_Lang); lang != nil {
			//设置语言
			c.Locale.Lang = lang.(string)
		} else {
			//设置默认语言
			c.Locale.Lang = "zh-CN"
		}

		data := make(map[string]string)
		if len(username) <= 0 {
			data["status"] = "0"
			data["message"] = c.Message("login_user_name")
			c.Data["json"] = data
			c.ServeJson()

			return
		}

		if len(password) <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("login_password")
			c.Data["json"] = data
			c.ServeJson()

			return
		}

		if false && len(verify) <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("login_verification_code")
			c.Data["json"] = data
			c.ServeJson()

			return
		}

		if false && !cpt.Verify(captchaId, verify) {
			data["status"] = "0"
			data["message"] = c.Message("login_verification_error")
			c.Data["json"] = data
			c.ServeJson()

			return
		}

		adminInfo := models.GetAdminByName(username, models.Reader)
		if adminInfo.Id <= 0 {
			data["status"] = "0"
			data["message"] = c.Message("admin_username_error")
		} else if adminInfo.Status == false && adminInfo.Id != utils.Founder {
			data["status"] = "0"
			data["message"] = c.Message("admin_forbid_login")
		} else if adminInfo.Role.Status == false && adminInfo.Id != utils.Founder {
			data["status"] = "0"
			data["message"] = c.Message("admin_forbid_role_login")
			//TODO 密码加密解密
		} else if username == adminInfo.UserName && security.CompareHashAndPassword(adminInfo.Password, password) {
			c.Controller.SetSession(utils.Session_Admin, adminInfo)
			c.Controller.SetSession(utils.Session_Lang, adminInfo.Lang)

			flash.Data["success"] = c.Message("login_success")
			flash.Data["url"] = beego.UrlFor("IndexController.Index")
			flash.Store(&c.Controller)

			//更新登陆时间
			models.UpdateAdminLoginInfo(adminInfo.Id, models.Writter)

			controller, action := c.GetControllerAndAction()
			// 操作日志
			addLogs(adminInfo.Id, controller, action, fmt.Sprintf("登陆系统!|^|用户[%v, %v]", adminInfo.Id, adminInfo.UserName))

			data["status"] = "1"
			data["url"] = beego.UrlFor("PublicController.Message")
			data["message"] = c.Message("login_success")

			defer func() {
				cpt.DelKey(captchaId)
			}()
		} else {
			data["status"] = "0"
			data["message"] = c.Message("login_password_error")
		}

		c.Data["json"] = data
		c.ServeJson()
	}
}

// @router /logout [get]
func (c *AdminController) Logout() {
	c.Controller.EnableRender = false
	flash := beego.NewFlash()

	if adminInfo, ok := getAdminInfoBySession(c.Controller); ok {
		controller, action := c.GetControllerAndAction()
		addLogs(adminInfo.Id, controller, action, fmt.Sprintf("退出系统|^|用户[%v, %v]", adminInfo.Id, adminInfo.UserName))

		// del session
		c.DelSession(utils.Session_Admin)
		c.DelSession(utils.Session_Lock)
	}

	flash.Data["success"] = c.Message("operation_success")
	flash.Data["url"] = beego.UrlFor("AdminController.Login")
	flash.Store(&c.Controller)

	c.Redirect(beego.UrlFor("PublicController.Message"), 302)
	return
}

// @router /admin/editInfo [get]
// @router /admin/editInfo [post]
func (c *AdminController) EditInfo() {
	if c.Controller.Ctx.Request.Method == "GET" {
		adminInfo, _ := getAdminInfoBySession(c.Controller)
		c.Data["admin"] = adminInfo

		c.TplNames = "admin/editInfo.html"
	} else if c.Controller.Ctx.Request.Method == "POST" {
		c.Controller.EnableRender = false
		flash := beego.NewFlash()

		if adminInfo, ok := getAdminInfoBySession(c.Controller); ok {

			realName := c.GetString("realname")
			if len(realName) > 0 {
				adminInfo.RealName = realName
			} else {
				flash.Error(c.Message("admin_realName_error"))
				flash.Data["url"] = beego.UrlFor("AdminController.EditInfo")
				flash.Store(&c.Controller)

				c.Redirect(beego.UrlFor("PublicController.Message"), 302)
				return
			}

			email := c.GetString("email")
			if len(email) > 0 {
				adminInfo.Email = email
			} else {

			}

			if models.UpdateAdminInfo(adminInfo, models.Writter) {
				controller, action := c.GetControllerAndAction()
				addLogs(adminInfo.Id, controller, action, "个人设置|^|个人信息")

				flash.Data["success"] = c.Message("operation_success")
				flash.Data["url"] = beego.UrlFor("AdminController.EditInfo")
				flash.Store(&c.Controller)

				c.Redirect(beego.UrlFor("PublicController.Message"), 302)
				return
			} else {
				flash.Error(c.Message("operation_success"))
				flash.Data["url"] = beego.UrlFor("AdminController.EditInfo")
				flash.Store(&c.Controller)

				c.Redirect(beego.UrlFor("PublicController.Message"), 302)
				return
			}
		} else {
			c.Redirect(beego.UrlFor("AdminController.Login"), 302)
			return
		}
	}
}

// @router /admin/editPwd [get]
// @router /admin/editPwd [post]
func (c *AdminController) EditPwd() {
	if c.Controller.Ctx.Request.Method == "GET" {
		adminInfo, _ := getAdminInfoBySession(c.Controller)
		c.Data["admin"] = adminInfo

		c.TplNames = "admin/editPwd.html"
	} else if c.Controller.Ctx.Request.Method == "POST" {
		c.Controller.EnableRender = false
		flash := beego.NewFlash()

		if adminInfo, ok := getAdminInfoBySession(c.Controller); ok {
			if models.UpdateAdminPwd(adminInfo, models.Writter) {
				flash.Data["success"] = c.Message("operation_success")
				flash.Data["url"] = beego.UrlFor("AdminController.EditPwd")
				flash.Store(&c.Controller)

				c.Redirect(beego.UrlFor("PublicController.Message"), 302)
				return
			} else {
				flash.Error(c.Message("operation_success"))
				flash.Data["url"] = beego.UrlFor("AdminController.EditPwd")
				flash.Store(&c.Controller)

				c.Redirect(beego.UrlFor("PublicController.Message"), 302)
				return
			}
		} else {
			c.Redirect(beego.UrlFor("AdminController.Login"), 302)
			return
		}
	}
}

// @router /admin [get]
func (c *AdminController) ListAdmin() {

}

// @Title addAdmin
// @Description 新增管理员
// @Param body body xxx true "新增管理员的描述信息"
// @Success 200 {object} xxx
// @router /admin [post]
func (c *AdminController) AddAdmin() {

}

// @Title updateAdmin
// @Description 更新管理员
// @Param adminId path int true "待更新的管理员ID"
// @Param body body xxx true "更新后的管理员信息"
// @Success 200 {object} xxx
// @router /admin/:adminId [put]
func (c *AdminController) UpdateAdmin() {

}

// @Title delAdmin
// @Description 删除管理员
// @Param adminId path int true "待删除的管理员ID"
// @router /admin/:adminId [delete]
func (c *AdminController) DelAdmin() {

}
