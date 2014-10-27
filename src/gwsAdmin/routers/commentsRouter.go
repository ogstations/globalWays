package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["gwsAdmin/controllers:PublicController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:PublicController"],
		beego.ControllerComments{
			"Message",
			`/message`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"EditInfo",
			`/admin/editInfo`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"EditInfo",
			`/admin/editInfo`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"EditPwd",
			`/admin/editPwd`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"EditPwd",
			`/admin/editPwd`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"ListAdmin",
			`/admin`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"AddAdmin",
			`/admin`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"UpdateAdmin",
			`/admin/:adminId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"DelAdmin",
			`/admin/:adminId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:RoleController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:RoleController"],
		beego.ControllerComments{
			"Index",
			`/role`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:RoleController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:RoleController"],
		beego.ControllerComments{
			"Member",
			`/role/member/:rid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:RoleController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:RoleController"],
		beego.ControllerComments{
			"AddRole",
			`/role/add`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:RoleController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:RoleController"],
		beego.ControllerComments{
			"EditRole",
			`/role/edit`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AjaxController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AjaxController"],
		beego.ControllerComments{
			"Pos",
			`/ajax/pos`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:ChannelTypeController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:ChannelTypeController"],
		beego.ControllerComments{
			"ListChannelType",
			`/channelType`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:ChannelTypeController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:ChannelTypeController"],
		beego.ControllerComments{
			"GetChannelTypeById",
			`/channelType/:channelId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:ChannelTypeController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:ChannelTypeController"],
		beego.ControllerComments{
			"UpdateChannelTypeById",
			`/channelType/:channelId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:ChannelTypeController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:ChannelTypeController"],
		beego.ControllerComments{
			"NewChannelType",
			`/channelType`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:IndexController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:IndexController"],
		beego.ControllerComments{
			"Index",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:IndexController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:IndexController"],
		beego.ControllerComments{
			"Main",
			`/main`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:IndexController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:IndexController"],
		beego.ControllerComments{
			"Left",
			`/left`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:IndexController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:IndexController"],
		beego.ControllerComments{
			"Left",
			`/left`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:MenuController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:MenuController"],
		beego.ControllerComments{
			"Index",
			`/menu`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:MenuController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:MenuController"],
		beego.ControllerComments{
			"AddMenuByPid",
			`/menu/add/:pid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:MenuController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:MenuController"],
		beego.ControllerComments{
			"AddMenuByPid",
			`/menu/add/:pid`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:MenuController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:MenuController"],
		beego.ControllerComments{
			"AddMenuTop",
			`/menu/add`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:MenuController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:MenuController"],
		beego.ControllerComments{
			"AddMenuTop",
			`/menu/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:MenuController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:MenuController"],
		beego.ControllerComments{
			"DelMenu",
			`/menu/delete/:mid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:MenuController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:MenuController"],
		beego.ControllerComments{
			"EditMenu",
			`/menu/edit/:mid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:MenuController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:MenuController"],
		beego.ControllerComments{
			"EditMenu",
			`/menu/edit/:mid`,
			[]string{"post"},
			nil})

}
