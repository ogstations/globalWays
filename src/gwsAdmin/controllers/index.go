package controllers

import (
	"github.com/astaxie/beego"
	"gwsAdmin/models"
	"gwsAdmin/utils"
	"html/template"
)

var (
	_ beego.Controller
)

// 后台首页controller
type IndexController struct {
	BaseController
}

// @router / [get]
func (c *IndexController) Index() {
	if adminInfo, ok := getAdminInfoBySession(c.Controller); ok {

		controller, action := c.GetControllerAndAction()
		// 控制器
		c.Data["controller"] = controller
		// 方法
		c.Data["action"] = action
		// 导航菜单
		c.Data["menus"] = models.GetMenuByPid(0, adminInfo.Id, true, models.Reader)
		// 管理员信息
		c.Data["admin"] = adminInfo

		//是否锁屏
		if lock := c.GetSession(utils.Session_Lock); lock == nil {
			c.Data["lock"] = "0"
		} else {
			c.Data["lock"] = "1"
		}
	} else {
		c.Redirect(beego.UrlFor("AdminController.Login"), 302)
		return
	}

	c.TplNames = "index/index.html"
}

// @router /main [get]
func (c *IndexController) Main() {

	if adminInfo, ok := getAdminInfoBySession(c.Controller); ok {

		// 管理员信息
		c.Data["admin"] = adminInfo

		//快捷面板
		c.Data["panelList"] = models.GetPanelList(adminInfo.Id, models.Reader)
	} else {
		c.Redirect(beego.UrlFor("AdminController.Login"), 302)
		return
	}

	c.TplNames = "index/main.html"
}


// @router /left [get]
// @router /left [post]
func (c *IndexController) Left() {
	pid, err := c.GetInt("pid")

	if adminInfo, ok := getAdminInfoBySession(c.Controller); ok {
		var leftHtml template.HTML
		if err != nil {
			leftHtml = models.GetLeftMenuHtml(1, adminInfo.Id, models.Reader)
		} else {
			leftHtml = models.GetLeftMenuHtml(pid, adminInfo.Id, models.Reader)
		}
		c.Data["left_menu"] = leftHtml
	} else {
		c.Redirect(beego.UrlFor("AdminController.Login"), 302)
		return
	}

	c.TplNames = "index/left.html"
}
