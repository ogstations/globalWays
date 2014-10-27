package routers

import (
	"gwsAdmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 后台首页路由
	beego.Include(&controllers.IndexController{})
	// 管理员路由
	beego.Include(&controllers.AdminController{})
	// AJAX路由
	beego.Include(&controllers.AjaxController{})
	// PUBLIC路由
	beego.Include(&controllers.PublicController{})
	// 菜单路由
	beego.Include(&controllers.MenuController{})
	// 角色路由
	beego.Include(&controllers.RoleController{})

	// 分发渠道路由
	beego.Include(&controllers.ChannelTypeController{})
}
